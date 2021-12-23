// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"double/app/admin/internal/biz"
	"double/app/admin/internal/conf"
	"double/app/admin/internal/data"
	"double/app/admin/internal/server"
	"double/app/admin/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
)

// initApp init kratos application.
// 添加*conf.Registry配置参数，用于在data中注入NewRegistrar
// 添加trace
func initApp(*conf.Server, *conf.Registry, *conf.Data, log.Logger, *tracesdk.TracerProvider) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
