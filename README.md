# Double（梅开二度）

## 介绍

足球俱乐部养成游戏，原型为[hattrick](https://m.hattrick.org/)
- rpc
- http
- 服务注册与发现
- consul配置中心
- 自增唯一id
- redis快速存储
- docker
- jaeger service trace
- fluent日志收集


## By Bilibili Kratos

前期准备见[kratos](https://github.com/go-kratos/kratos)

```
go get -u github.com/go-kratos/kratos/cmd/kratos/v2@latest
```

## Docker

### golang
```shell
docker run -itd -p 8880:8000 -p 9990:9000 --name golang -v /Users/gaea/docker/golang/workspace:/workspace --workdir /workspace golang
```
golang2容器用于负载均衡测试
```shell
docker run -itd -p 8880:8000 -p 9990:9000 --name golang2 -v /Users/gaea/docker/golang/workspace:/workspace --workdir /workspace golang
```

### 其他
```shell
# redis容器开放6379，pika就开放7272

# jaeger容器默认端口
docker run -d -p 5775:5775/udp -p 16686:16686 -p 14250:14250 -p 14268:14268 jaegertracing/all-in-one:latest

# fluentd
docker run -d -p 24224:24224 -p 24224:24224/udp fluent/fluentd:v1.3-debian-1 
```
容器部署后，涉及到多容器间数据交互时，如果对性能测试有要求，不能配置host.docker.internal，需要让容器在同一网段（比如connect same network），否则影响性能。

经测试涉及服务间交互的ListUser过程（grpc，1989字节）
```shell
curl 127.0.0.1:8002/admin/v1/user
```
在jaeger UI显示平均耗时在10-15ms之间

## 注册中心

### consul
docker官方容器，默认开放8500端口。不需要注册中心需要修改main.go与data/data.go，去掉registry部分

## RUN
项目根路径下
`kratos run`

支持选择app启动

## 错误处理

- 如果 go generate ./... 失败，提示wire包不完整，执行：

`go get -d github.com/google/wire/cmd/wire@v0.5.0`

- 包导入错误，重整mod

`go mod tidy`

- 配置错误，重新生成配置。注意Makefile中的配置路径查找命令

`make config`

- 连接DB、Redis错误。请确认是否能正常通信，比如localhost与host.docker.internal。

- wire_gen.go生成失败。需要手动修改或刷新一下wire_gen.go，再重新generate

- pgv validate proto file error,提示
  
  `// no validation rules for xxx`
  这是已知的PGV bug：[github issue](https://github.com/envoyproxy/protoc-gen-validate/issues/240)

---

## 当前进度

repo：可以用go-redis@v8 连接redis，支持并发唯一自增id。rpc服务超时时间1.2s。自增id每秒分配50。

userService：账号管理。支持rpc与http。注册到consul server

clubService：常规rpc service。初始化完毕。doing。为了完成目标，可能会简化业务

adminService:后台管理服务，服务发现demo。支持consul发现、调用其他service。配置不再读本地，改读consul配置中心。通过jaeger查看服务间trace过程

## 计划
目前最重要的部分都已完工或有demo体现。

官方文档支持的功能中，只有部分middleware没有涉及。包括middleware传递metadata，监控Prometheus，Auth认证鉴权，熔断与限流。

短时间内不打算支持ent/Gorm，因为不考虑mysql作为存储对象。有计划将redis替换成pika。

暂时不考虑efk，目前只支持fluent收集日志。可以把efk实现BI系统作为最后目标。

swagger需要看官方更新，目前存在一些问题。