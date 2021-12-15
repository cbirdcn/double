package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"strconv"
	"sync"
	"time"
)

func getUserKey(id int64) string{
	return "user:" + strconv.FormatInt(id, 10)
}

func getUserNameSetKey() string{
	return "uq_user_name"
}

func getClubNameSetKey() string{
	return "uq_club_name"
}

func getAutoIncrementId() int64{
	// todo
	return 1
}

func getIdByNameHashKey() string{
	return "name_to_id"
}

/////////////////////////////////////redis setnx uuid lock///////////////////////////////////////////////////////////

/**
redis lock
1.无死锁，但不可重入：依靠锁ttl判断是否占用，重入表示获取一个新锁（可以手动添加uuid解决重入），死锁依靠过期判断和redis自动清理过期id
2.非分布式：即使分布式部署服务器，也要请求同一个redis服务器，否则无法顺序获取id。可以加雪花id做成分布式非连续id。
3.非阻塞锁：未拿到锁时客户端轮询sleep时间请求服务端
4.高可用：大规模请求过来后，只要可以获取redis连接，最差结果就是等待到超时时间后没拿到锁
5.性能：获取锁+释放锁获取自增id，1000ms内并发40-50请求。
使用：
1.提供一个独一无二的lockKey（会作为redis存储的key）
2.获取锁成功，给自增idINCR，最小为1（int64）。失败返回id=0
*/

var WaitTimeOut = 1000 // 最长抢锁等待时间，ms
var KeepTimeOut = 1000 // 最长保持锁定时间，ms

/**
XXX:注意，抢锁和持有锁时间要尽可能短。因为配置文件中规定了grpc timeout和write_timeout、read_timeout。也就是说抢锁时间+上下文运行时间（比如前后有其他redis操作）<= grpc timeout
举例，假设grpc_timeout = 1.5s, redis_rw_timeout = 0.2s, 那么锁定时间1s就是比较合理的。
如果grpc_timeout = 1s。就会导致高并发时，部分无法抢到锁的协程在最后时刻，被context判断超时，报出context deadline exceeded的错误
未来可以考虑在const中写好特定业务的锁定时间。每个业务可以有不同的锁定时间，不允许传入超时参数，只允许修改整个业务模块的超时常量或新增拆分出来的常量
*/


// 获取自增id(setnx+uuid)
func getAutoIncId(ctx context.Context, rdb *redis.Client, lockKey string) (int64, error){
	var wg sync.WaitGroup
	ch := make(chan int64, 1)
	go makeUuidLock(ctx, rdb, &wg, lockKey, &ch)

	if id := <-ch; id != 0 {
		return id, nil
	}
	return 0, errors.New(500, "concurrency", "can not get auto increment id")
}

func incrId(ctx context.Context, p *redis.Client, idKey string) (id int64){
	return p.Incr(ctx, idKey).Val()
}

func currentMilliSecond() int64{
	return int64(time.Now().UnixNano() / 1e6)
}


//////////setnx+uuid///////////
//逻辑：set k v nx px ttl获取锁，拿不到就循环重新获取，直到拿到锁或等待时间到期。只需要释放已拿到的锁，当锁值=uuid时释放。即使不释放redis也可以根据锁到期时间删除锁。
//todo:分布式，需要支持多客户端能索引到同一台服务器的同一个key


// setnx+uuid
func makeUuidLock(ctx context.Context, rdb *redis.Client, wg *sync.WaitGroup, lockKey string, ch *chan int64){
	// 提供参数：连接、锁名、是否自增、并发id(uuid4)
	u4 := uuid.New().String()

	// 实际使用时，需要设定临时lockKey，比如根据方法名
	locked, id := makeSetnxLock(ctx, rdb, lockKey, u4,true)

	if locked {
		//fmt.Println("locking " + strconv.Itoa(i) + " " + u4)
		// 获取到锁才需要释放，即使不释放也会被redis根据ttl到期自动回收
		unlock(ctx, rdb, lockKey, u4)
	}else{
		//fmt.Println("can't get lock " + strconv.Itoa(i) + " " + u4)
	}
	*ch <- id
}


// 获取锁，获取到返回true，否则返回false
// 参数：incr表示是否需要自增，true表示自增（从1开始）,u4表示并发id
// 返回：locked=true表示获取锁成功，id是获取的自增id（incr=false时或locked=false时id=0，正常id>0)，KeepTimeOut是锁剩余时间,lockKey是redis存储自增id的键名（返回INCR后的值，可以避免返回0，todo：和db incr打通）
func makeSetnxLock(ctx context.Context, p *redis.Client, lockKey string, u4 string, incr bool) (locked bool, id int64){
	locked = false
	start := currentMilliSecond()
	tmpLockKey := "tmp" + lockKey

	// 如果等待时间还没到超时时间，循环
	for ;int(currentMilliSecond() - start) < WaitTimeOut; {
		// 官方：SET resource_name my_random_value NX PX 30000
		// The command will set the key only if it does not already exist (NX option), with an expire of 30000 milliseconds (PX option). The key is set to a value “my_random_value”.
		// https://redis.io/topics/distlock
		//res, err := redis.String(p.Do("SET", lockKey, u4, "NX", "PX", keepTimeout))
		// XXX:tmpLockKey只是抢锁才会使用的临时key，所以不要使用tmp开头的字符串作为key
		res := p.SetNX(ctx, tmpLockKey, u4, time.Duration(KeepTimeOut) * time.Millisecond).Val()

		// 假设有ABC三个协程同时获取锁
		// A协程获取到锁
		if res {
			// 如果要返回自增id
			id := int64(0)
			if incr {
				// INCR过程最好放到acquireLock()里，并且必须用INCR返回值
				id = incrId(ctx, p, lockKey)
				// 只能认为id基本可靠
			}

			// 获取锁成功
			return true, id
		}

		// BC协程没获取到锁，加延迟后继续循环，不控制循环次数和续期时长
		time.Sleep(time.Duration(1)*time.Millisecond)
	}

	// 等待超时，获取锁失败
	return false, 0
}

// 释放锁
// 获取锁成功才需要调用
func unlock(ctx context.Context, p *redis.Client, lockKey string, u4 string){
	tmpLockKey := "tmp" + lockKey
	newValue := p.Get(ctx, tmpLockKey).Val()

	// TODO：如果支持lua，把判等和del放到一个lua脚本内执行更好。pika不支持lua所以保持原样
	if u4 == newValue {
		p.Del(ctx, tmpLockKey)
		//fmt.Println("del " + strconv.Itoa(i) + " " + u4)
	}else{
		//fmt.Println("not del " + strconv.Itoa(i) + " " + u4)
	}

	// go-redis自带连接池，不需要释放
	//_ = p.Close()
}

////////////////////////////////////redis setnx uuid lock end////////////////////////////////////////////////////////