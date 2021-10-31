package data

import (
	"context"

	"github.com/lalifeier/vgo/app/shop/admin/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"

	umsV1 "github.com/lalifeier/vgo/api/ums/service/v1"

	consul "github.com/go-kratos/consul/registry"
	consulAPI "github.com/hashicorp/consul/api"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewDiscovery, NewUmsServiceClient)

// Data .
type Data struct {
	log       *log.Helper
	UmsClient umsV1.UmsClient
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, umsClient umsV1.UmsClient) (*Data, func(), error) {
	log := log.NewHelper(log.With(logger, "module", "shop-admin/data"))
	cleanup := func() {
		log.Info("closing the data resources")
	}
	return &Data{
		log:       log,
		UmsClient: umsClient,
	}, cleanup, nil
}

func NewDiscovery(conf *conf.Registry) registry.Discovery {
	c := consulAPI.DefaultConfig()
	c.Address = conf.Consul.Address
	c.Scheme = conf.Consul.Scheme
	cli, err := consulAPI.NewClient(c)
	if err != nil {
		panic(err)
	}
	r := consul.New(cli, consul.WithHealthCheck(false))
	return r
}

func NewUmsServiceClient(r registry.Discovery) umsV1.UmsClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		// grpc.WithEndpoint("discovery:///vgo.ums.service"),
		grpc.WithEndpoint("127.0.0.1:9001"),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	c := umsV1.NewUmsClient(conn)
	return c
}
