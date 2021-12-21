# Double（梅开二度）

## 介绍

足球俱乐部养成游戏，原型为[hattrick](https://m.hattrick.org/)
- rpc
- http


## By Bilibili Kratos

前期准备见[kratos](https://github.com/go-kratos/kratos)

```
go get -u github.com/go-kratos/kratos/cmd/kratos/v2@latest
```

## Docker

部署golang容器，开放8000做http端口，9000做rpc端口

## 注册中心

consul:docker官方容器，默认开放8500端口。不需要注册中心需要修改main.go与data/data.go，去掉registry部分

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

userService：rpc与http ok。可以注册到consul

clubService：初始化完毕。rpc doing

