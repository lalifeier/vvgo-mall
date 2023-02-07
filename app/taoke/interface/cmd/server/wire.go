//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/lalifeier/vvgo-mall/app/taoke/interface/internal/biz"
	"github.com/lalifeier/vvgo-mall/app/taoke/interface/internal/conf"
	"github.com/lalifeier/vvgo-mall/app/taoke/interface/internal/data"
	"github.com/lalifeier/vvgo-mall/app/taoke/interface/internal/server"
	"github.com/lalifeier/vvgo-mall/app/taoke/interface/internal/service"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Data, *conf.Auth, *conf.Registry, log.Logger, *tracesdk.TracerProvider) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
