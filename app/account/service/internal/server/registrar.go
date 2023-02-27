package server

import (
	"github.com/go-kratos/kratos/v2/registry"

	"github.com/lalifeier/vvgo-mall/app/account/service/internal/conf"
	"github.com/lalifeier/vvgo-mall/pkg/bootstrap"
)

func NewRegistrar(conf *conf.Registry) registry.Registrar {
	return bootstrap.NewConsulRegistry(conf.Consul.Address, conf.Consul.Scheme, conf.Consul.HealthCheck)
}
