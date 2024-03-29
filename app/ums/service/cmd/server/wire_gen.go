// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/lalifeier/vvgo-mall/app/ums/service/internal/biz"
	"github.com/lalifeier/vvgo-mall/app/ums/service/internal/conf"
	"github.com/lalifeier/vvgo-mall/app/ums/service/internal/data"
	"github.com/lalifeier/vvgo-mall/app/ums/service/internal/server"
	"github.com/lalifeier/vvgo-mall/app/ums/service/internal/service"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(confServer *conf.Server, confData *conf.Data, registry *conf.Registry, logger log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	accountUserRepo := data.NewAccountUserRepo(dataData, logger)
	accountUserUseCase := biz.NewAccountUserUseCase(accountUserRepo, logger)
	umsService := service.NewUmsService(accountUserUseCase, logger)
	grpcServer := server.NewGRPCServer(confServer, umsService, logger)
	registrar := server.NewRegistrar(registry)
	app := newApp(logger, grpcServer, registrar)
	return app, func() {
		cleanup()
	}, nil
}
