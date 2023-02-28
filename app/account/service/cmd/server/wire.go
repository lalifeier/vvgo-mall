//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/lalifeier/vvgo-mall/app/account/service/internal/biz"
	"github.com/lalifeier/vvgo-mall/app/account/service/internal/data"
	"github.com/lalifeier/vvgo-mall/app/account/service/internal/server"
	"github.com/lalifeier/vvgo-mall/app/account/service/internal/service"
	"github.com/lalifeier/vvgo-mall/gen/api/go/common/conf"
)

// initApp init kratos application.
func wireApp(*conf.Server, *conf.Registry, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
