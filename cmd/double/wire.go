// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"double/internal/biz"
	"double/internal/conf"
	"double/internal/data"
	"double/internal/server"
	"double/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
