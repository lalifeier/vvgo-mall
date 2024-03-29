// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/lalifeier/vvgo-mall/app/paser/service/internal/biz"
	"github.com/lalifeier/vvgo-mall/app/paser/service/internal/conf"
	"github.com/lalifeier/vvgo-mall/app/paser/service/internal/data"
	"github.com/lalifeier/vvgo-mall/app/paser/service/internal/server"
	"github.com/lalifeier/vvgo-mall/app/paser/service/internal/service"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	paserRepo := data.NewPaserRepo(dataData, logger)
	paserUseCase := biz.NewPaserUseCase(paserRepo, logger)
	paserService := service.NewPaserService(paserUseCase, logger)
	httpServer := server.NewHTTPServer(confServer, paserService, logger)
	grpcServer := server.NewGRPCServer(confServer, paserService, logger)
	app := newApp(logger, httpServer, grpcServer)
	return app, func() {
		cleanup()
	}, nil
}
