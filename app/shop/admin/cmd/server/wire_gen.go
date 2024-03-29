// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/lalifeier/vvgo-mall/app/shop/admin/internal/biz"
	"github.com/lalifeier/vvgo-mall/app/shop/admin/internal/conf"
	"github.com/lalifeier/vvgo-mall/app/shop/admin/internal/data"
	"github.com/lalifeier/vvgo-mall/app/shop/admin/internal/server"
	"github.com/lalifeier/vvgo-mall/app/shop/admin/internal/service"
	"go.opentelemetry.io/otel/sdk/trace"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(confServer *conf.Server, confData *conf.Data, casbin *conf.Casbin, auth *conf.Auth, registry *conf.Registry, logger log.Logger, tracerProvider *trace.TracerProvider) (*kratos.App, func(), error) {
	client := data.NewEntClient(confData, logger)
	redisClient := data.NewRedisClient(confData, logger)
	enforcer := data.NewEnforcer(confData)
	asyncProducer := data.NewKafkaProducer(confData)
	consumer := data.NewKafkaConsumer(confData)
	discovery := data.NewDiscovery(registry)
	invalid type := data.NewAccountServiceClient(discovery, tracerProvider)
	dataData, cleanup, err := data.NewData(logger, client, redisClient, enforcer, asyncProducer, consumer, invalid type)
	if err != nil {
		return nil, nil, err
	}
	userRepo := data.NewUserRepo(dataData, logger)
	userUseCase := biz.NewUserUseCase(userRepo, enforcer, logger)
	roleRepo := data.NewRoleRepo(dataData, logger)
	roleUseCase := biz.NewRoleUseCase(roleRepo, enforcer, logger)
	accountRepo := data.NewAccountRepo(dataData, logger)
	authUseCase := biz.NewAuthUseCase(auth, userRepo, accountRepo)
	dictDataRepo := data.NewDictDataRepo(dataData, logger)
	dictDataUseCase := biz.NewDictDataUseCase(dictDataRepo, logger)
	shopService := service.NewShopService(logger, userUseCase, roleUseCase, authUseCase, dictDataUseCase)
	httpServer := server.NewHTTPServer(confServer, auth, logger, shopService, enforcer)
	grpcServer := server.NewGRPCServer(confServer, auth, logger, shopService)
	registrar := data.NewRegistrar(registry)
	app := newApp(logger, httpServer, grpcServer, registrar)
	return app, func() {
		cleanup()
	}, nil
}

