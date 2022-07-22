package data

import (
	"context"
	"database/sql"
	"double/app/user/internal/conf"
	consul "github.com/go-kratos/consul/registry"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	_ "github.com/go-kratos/kratos/v2/transport/grpc"
	_ "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-redis/redis/extra/redisotel"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	consulAPI "github.com/hashicorp/consul/api"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewUserRepo, NewRegistrar)

// Data .
type Data struct {
	// wrapped database client
	rdb *redis.Client
	mydb *sql.DB
	mdb *mongo.Client
	mdbCtx context.Context
}

// 数据库连接：NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper(logger)
	var err error

	// 创建redis连接
	newRedisClient := redis.NewClient(&redis.Options{
		Addr:         c.Redis.Addr,
		Password:     c.Redis.Password,
		DB:           int(c.Redis.Db),
		DialTimeout:  c.Redis.DialTimeout.AsDuration(),
		WriteTimeout: c.Redis.WriteTimeout.AsDuration(),
		ReadTimeout:  c.Redis.ReadTimeout.AsDuration(),
	})
	newRedisClient.AddHook(redisotel.TracingHook{})

	// 创建mysql连接
	newMysqlClient, err := gorm.Open(mysql.Open(c.Mysql.Source), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	sqlDb, err := newMysqlClient.DB()
	if err != nil {
		log.Fatalf("failed opening connection pool to mysql: %v", err)
	}
	// TODO:暂时只提供三个参数
	sqlDb.SetMaxOpenConns(int(c.Mysql.MaxOpenConns))
	sqlDb.SetMaxIdleConns(int(c.Mysql.MaxIdleConns))
	sqlDb.SetConnMaxLifetime(c.Mysql.MaxLifeTime.AsDuration())
	if err = sqlDb.Ping(); err != nil {
		log.Fatalf("failed ping connection pool to mysql: %v", err)
	}

	// 创建mongodb连接
	//mongoCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//defer cancel()
	mdbCtx := context.TODO() // TODO:暂时不考虑ctx时间限制
	clientOptions := options.Client().ApplyURI(c.Mongo.Addr)
	// TODO:暂不提供writeConcern相关参数，w给默认值majority，j与wtimeout不管
	clientOptions.SetWriteConcern(writeconcern.New(writeconcern.WMajority()))
	clientOptions.SetMaxPoolSize(c.Mongo.MaxPoolSize)
	clientOptions.SetMaxConnIdleTime(c.Mongo.MaxIdleTimeMs.AsDuration())
	// 创建连接实例，没有真正连接数据库
	NewMongoClient, err := mongo.Connect(mdbCtx, clientOptions)
	if err != nil {
		panic(err)
	}
	// 连接数据库检查
	if err = NewMongoClient.Ping(mdbCtx, nil); err != nil {
		panic(err)
	}

	d := &Data{
		rdb: newRedisClient,
		mydb: sqlDb,
		mdb: NewMongoClient,
		mdbCtx: mdbCtx,
	}

	return d, func() {
		log.Info("message", "closing the data resources")
		if err = d.mydb.Close(); err != nil {
			log.Error(err)
		}
		if err := d.rdb.Close(); err != nil {
			log.Error(err)
		}
		if err = d.mdb.Disconnect(mdbCtx); err != nil {
			log.Error(err)
		}
	}, nil
}

// 注册服务（连接注册中心）
func NewRegistrar(conf *conf.Registry) registry.Registrar {
	c := consulAPI.DefaultConfig()
	c.Address = conf.Consul.Address
	c.Scheme = conf.Consul.Scheme
	client, err := consulAPI.NewClient(c)
	if err != nil {
		panic(err)
	}

	r := consul.New(client, consul.WithHealthCheck(false))
	return r
}

func NewDiscovery(conf *conf.Registry) registry.Discovery {
	c := consulAPI.DefaultConfig()
	c.Address = conf.Consul.Address
	c.Scheme = conf.Consul.Scheme
	client, err := consulAPI.NewClient(c)
	if err != nil {
		panic(err)
	}

	r := consul.New(client, consul.WithHealthCheck(false))
	return r
}

// 服务发现：todo:需要一个其他service
//func NewGameServiceClient(conf *conf.Registry) registry.Discovery {
//	// consul客户端
//	c := api.DefaultConfig()
//	c.Address = conf.Consul.Address
//	client, err := api.NewClient(c)
//	if err != nil {
//		panic(err)
//	}
//
//	dis := consul.New(client)
//
//	endpoint := "discovery://default/user"
//	conn, err := grpc.Dial(context.Background(), grpc.WithEndpoint(endpoint), grpc.WithDiscovery(dis))
//	if err != nil {
//		panic(err)
//	}
//	return conn
//}

