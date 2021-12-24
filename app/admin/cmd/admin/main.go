package main

import (
	"double/app/admin/internal/conf"
	"flag"
	consulConfig "github.com/go-kratos/kratos/contrib/config/consul/v2"
	"github.com/go-kratos/kratos/v2/config/file"
	consulAPI "github.com/hashicorp/consul/api"
	"gopkg.in/yaml.v2"

	//consulReg "github.com/go-kratos/consul/registry"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	//"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"os"
	"github.com/go-kratos/kratos/contrib/log/fluent/v2"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name = "admin"
	// Version is the version of the compiled software.
	Version = "1.0.0"
	// flagconf is the config flag.
	flagconf string

	id, _ = os.Hostname()
)

func init() {
	id = id + "_" + Name // xxx:同一容器内不同端口启动不同服务时，注册时需要保证id、name不一致
	// flag包解析命令行参数
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
}

func newApp(logger log.Logger, hs *http.Server, gs *grpc.Server, rr registry.Registrar) *kratos.App {
	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			hs,
			gs,
		),
		kratos.Registrar(rr), // 注册
	)
}

func main() {
	// flag解析命令行参数
	flag.Parse()
	// 日志记录方式1：默认stdout
	//logger := log.With(log.NewStdLogger(os.Stdout),
	//	"ts", log.DefaultTimestamp,
	//	"caller", log.DefaultCaller,
	//	"service.id", id,
	//	"service.name", Name,
	//	"service.version", Version,
	//	"trace_id", tracing.TraceID(),
	//	"span_id", tracing.SpanID(),
	//)
	// 日志记录方式2：fluentd服务器
	loggerSource, err := fluent.NewLogger("tcp://host.docker.internal:24224")
	if err != nil {
		panic(err)
	}
	//flog := log.NewHelper(logger)
	//flog.Debug("%s log test", "fluent")
	logger := log.With(loggerSource,
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", id,
		"service.name", Name,
		"service.version", Version,
		"trace_id", tracing.TraceID(),
		"span_id", tracing.SpanID(),
	)

	////////xxx:读配置方法1：本地文件配置////////

	//// 通过file source 创建 config实例，默认为configs/config.yaml
	//c := config.New(
	//	config.WithSource(
	//		file.NewSource(flagconf),
	//	),
	//)
	//defer c.Close()
	//
	//// 加载整个配置文件到cache(涉及多文件合并、覆盖等)
	//if err := c.Load(); err != nil {
	//	panic(err)
	//}
	//
	//// 解析配置文件到conf.pb.go中生成的配置类型，Bootstrap包含整个结构
	//var bc conf.Bootstrap
	//if err := c.Scan(&bc); err != nil {
	//	panic(err)
	//}
	//
	//// 解析其他文件，类似Bootstrap
	//var rc conf.Registry
	//if err := c.Scan(&rc); err != nil {
	//	panic(err)
	//}
	//
	//// 使用配置：
	//// jaeger支持服务间路由追踪，提供ui
	//// https://www.cnblogs.com/whuanle/p/14598049.html
	//exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(bc.Trace.Endpoint)))
	//if err != nil {
	//	panic(err)
	//}
	//// trace内容
	//tp := tracesdk.NewTracerProvider(
	//	tracesdk.WithBatcher(exp),
	//	tracesdk.WithResource(resource.NewSchemaless(
	//		semconv.ServiceNameKey.String(Name),
	//	)),
	//)
	//
	//app, cleanup, err := initApp(bc.Server, &rc, bc.Data, logger, tp)
	//if err != nil {
	//	panic(err)
	//}
	//defer cleanup()


	////////xxx:读配置方法2：配置中心////////

	// 从配置文件（configs/registry.yaml）读取注册中心地址
	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)
	defer c.Close()
	// 加载本地配置文件到缓存
	if err := c.Load(); err != nil {
		panic(err)
	}
	// 解析本地的注册配置文件
	var rc conf.Registry
	if err := c.Scan(&rc); err != nil {
		panic(err)
	}

	// 通过本地注册配置文件，连接consul服务器
	consulClient, err := consulAPI.NewClient(&consulAPI.Config{
		Address: rc.Consul.Address,
		Scheme: rc.Consul.Scheme,
	})
	if err != nil {
		panic(err)
	}
	// 构造config source，注意withPath直接填key名，key不建议用const，最好放在yaml中
	cs, err := consulConfig.New(consulClient, consulConfig.WithPath(rc.Consul.Key))
	if err != nil {
		panic(err)
	}
	// 对consul中key的value执行解码、转map过程，并创建config实例
	// 官方推荐更好的decoder方法:https://go-kratos.dev/docs/component/config/#5%E9%85%8D%E7%BD%AE%E8%A7%A3%E6%9E%90decoder
	c = config.New(config.WithSource(cs),
		config.WithDecoder(func(kv *config.KeyValue, v map[string]interface{}) error {
			return yaml.Unmarshal(kv.Value, v)
		}))
	defer c.Close()

	// 与本地文件一样，需要加载config map到cache，合并、覆盖kv等
	if err := c.Load(); err != nil {
		panic(err)
	}
	// 使用配置，与本地过程相同，解析到pb.go中的type里，此处忽略解码失败的处理过程
	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}
	// jaeger支持服务间路由追踪，提供ui
	// https://www.cnblogs.com/whuanle/p/14598049.html
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(bc.Trace.Endpoint)))
	if err != nil {
		panic(err)
	}
	// trace内容
	tp := tracesdk.NewTracerProvider(
		tracesdk.WithBatcher(exp),
		tracesdk.WithResource(resource.NewSchemaless(
			semconv.ServiceNameKey.String(Name),
		)),
	)

	app, cleanup, err := initApp(bc.Server, &rc, bc.Data, logger, tp)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}
