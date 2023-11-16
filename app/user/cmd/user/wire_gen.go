// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"double/app/user/internal/biz"
	"double/app/user/internal/conf"
	"double/app/user/internal/data"
	"double/app/user/internal/server"
	"double/app/user/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

// Injectors from wire.go:

// initApp init kratos application.
// 添加*conf.Registry配置参数，用于在data中注入NewRegistrar
func initApp(confServer *conf.Server, registry *conf.Registry, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	userRepo := data.NewUserRepo(dataData, logger)
	userUseCase := biz.NewUserUseCase(userRepo, logger)
	userService := service.NewUserService(userUseCase, logger)
	httpServer := server.NewHTTPServer(confServer, userService, logger, userUseCase)
	grpcServer := server.NewGRPCServer(confServer, userService, logger, userUseCase)
	registrar := data.NewRegistrar(registry)
	app := newApp(logger, httpServer, grpcServer, registrar)
	return app, func() {
		cleanup()
	}, nil
}
