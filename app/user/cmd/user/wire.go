// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"double/app/user/internal/biz"
	"double/app/user/internal/conf"
	"double/app/user/internal/data"
	"double/app/user/internal/server"
	"double/app/user/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// initApp init kratos application.
// 添加*conf.Registry配置参数，用于在data中注入NewRegistrar
func initApp(*conf.Server, *conf.Registry, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
