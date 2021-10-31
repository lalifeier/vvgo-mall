// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/lalifeier/vgo/app/shop/admin/internal/biz"
	"github.com/lalifeier/vgo/app/shop/admin/internal/conf"
	"github.com/lalifeier/vgo/app/shop/admin/internal/data"
	"github.com/lalifeier/vgo/app/shop/admin/internal/data/ums"
	"github.com/lalifeier/vgo/app/shop/admin/internal/server"
	"github.com/lalifeier/vgo/app/shop/admin/internal/service"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Data, *conf.Registry, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, ums.ProviderSet, newApp))
}
