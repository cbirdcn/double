package data

import (
	"context"
	userv1 "double/api/user/v1"
	"double/app/admin/internal/conf"

	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/selector/wrr"
	"github.com/go-redis/redis/v8"

	//"fmt"
	consul "github.com/go-kratos/consul/registry"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	_ "github.com/go-kratos/kratos/v2/transport/grpc"
	_ "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-redis/redis/extra/redisotel"
	"github.com/google/wire"
	consulAPI "github.com/hashicorp/consul/api"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"

	//"github.com/go-kratos/kratos/v2/selector/p2c"
	"github.com/go-kratos/kratos/v2/selector/filter"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewDiscovery, NewRegistrar, NewUserServiceClient, NewUserRepo)

// Data .
type Data struct {
	// wrapped database client
	rdb *redis.Client
	uc  userv1.UserClient // 其他服务的client
}

// 数据库连接：NewData .
func NewData(c *conf.Data, uc userv1.UserClient, logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper(logger)
	rdb := redis.NewClient(&redis.Options{
		Addr:         c.Redis.Addr,
		Password:     c.Redis.Password,
		DB:           int(c.Redis.Db),
		DialTimeout:  c.Redis.DialTimeout.AsDuration(),
		WriteTimeout: c.Redis.WriteTimeout.AsDuration(),
		ReadTimeout:  c.Redis.ReadTimeout.AsDuration(),
	})
	rdb.AddHook(redisotel.TracingHook{})
	d := &Data{
		rdb: rdb,
		uc:  uc,
	}

	return d, func() {
		log.Info("message", "closing the data resources")
		if err := d.rdb.Close(); err != nil {
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

// 服务发现
// xxx：服务发现的目的是要调用其他服务，比如admin后台查询其他服务操作记录。这里仅做展示。
// 融合写法
//func NewUserServiceClient(conf *conf.Registry) userv1.UserClient {
//	// consul客户端
//	c := consulAPI.DefaultConfig()
//	c.Address = conf.Consul.Address
//	client, err := consulAPI.NewClient(c)
//	if err != nil {
//		panic(err)
//	}
//
//	dis := consul.New(client)
//
//	endpoint := "discovery:///user"
//	conn, err := grpc.Dial(context.Background(), grpc.WithEndpoint(endpoint), grpc.WithDiscovery(dis))
//	fmt.Println(conn)
//	if err != nil {
//		panic(err)
//	}
//	ret := userv1.NewUserClient(conn)
//
//	fmt.Println(ret)
//	return ret
//}
// 分离写法
func NewUserServiceClient(r registry.Discovery, tp *tracesdk.TracerProvider) userv1.UserClient {
	// 创建路由Filter：筛选版本号为"2.0.0"的实例
	filterCase := filter.Version("1.0.0")
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///user"), // 格式：<schema>://[namespace]/<service-name>
		grpc.WithDiscovery(r),
		// 由于gRPC框架的限制只能使用全局balancer name的方式来注入selector
		// http用p2c
		grpc.WithBalancerName(wrr.Name),
		// 通过grpc.WithFilter注入路由Filter(WithSelectFilter可能换成了WithFilter)
		grpc.WithFilter(filterCase),
		grpc.WithMiddleware(
			tracing.Client(tracing.WithTracerProvider(tp)),
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	return userv1.NewUserClient(conn)
}
