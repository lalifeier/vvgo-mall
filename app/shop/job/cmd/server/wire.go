//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/lalifeier/vvgo-mall/app/shop/job/internal/biz"
	"github.com/lalifeier/vvgo-mall/app/shop/job/internal/conf"
	"github.com/lalifeier/vvgo-mall/app/shop/job/internal/data"
	"github.com/lalifeier/vvgo-mall/app/shop/job/internal/server"
	"github.com/lalifeier/vvgo-mall/app/shop/job/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Registry, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
