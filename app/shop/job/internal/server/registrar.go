package server

import (
	"github.com/go-kratos/kratos/v2/registry"

	"github.com/lalifeier/vvgo-mall/gen/api/go/common/conf"
	xregistry "github.com/lalifeier/vvgo-mall/pkg/bootstrap/registry"
)

func NewRegistrar(conf *conf.Registry) registry.Registrar {
	return xregistry.NewConsulRegistry(conf.Consul.Address, conf.Consul.Scheme, conf.Consul.HealthCheck)
}
