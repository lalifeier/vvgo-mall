package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-redis/redis"

	consul "github.com/go-kratos/consul/registry"
	"github.com/go-kratos/kratos/v2/log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	consulAPI "github.com/hashicorp/consul/api"
	"github.com/lalifeier/vvgo-mall/app/auth/service/internal/conf"
	"github.com/lalifeier/vvgo-mall/app/auth/service/internal/data/ent"

	"github.com/lalifeier/vvgo-mall/app/auth/service/internal/data/ent/migrate"
	_ "github.com/lalifeier/vvgo-mall/app/auth/service/internal/data/ent/runtime"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewEntClient, NewRedisClient, NewDiscovery, NewRegistrar)

// Data .
type Data struct {
	log *log.Helper

	db  *ent.Client
	rdb *redis.Client
}

// NewData .
func NewData(logger log.Logger, entClient *ent.Client, redisClient *redis.Client) (*Data, func(), error) {
	log := log.NewHelper(log.With(logger, "module", "auth-service/data"))

	d := &Data{
		log: log,
		db:  entClient,
		rdb: redisClient,
	}

	cleanup := func() {
		if err := d.db.Close(); err != nil {
			log.Error(err)
		}
	}
	return d, cleanup, nil
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

func NewRegistrar(conf *conf.Registry) registry.Registrar {
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

func NewEntClient(conf *conf.Data, logger log.Logger) *ent.Client {
	log := log.NewHelper(log.With(logger, "module", "auth-service/data/ent"))

	client, err := ent.Open(
		conf.Database.Driver,
		conf.Database.Source,
		ent.Debug(),
	)
	if err != nil {
		log.Fatalf("failed opening connection to db: %v", err)
	}
	// Run the auto migration tool.
	if err := client.Debug().Schema.Create(
		context.Background(),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return client
}

func NewRedisClient(c *conf.Data, logger log.Logger) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:         c.Redis.Network,
		Password:     c.Redis.Password,
		DB:           int(c.Redis.Db),
		DialTimeout:  c.Redis.DialTimeout.AsDuration(),
		ReadTimeout:  c.Redis.ReadTimeout.AsDuration(),
		WriteTimeout: c.Redis.WriteTimeout.AsDuration(),
	})
	// err := client.Ping().Err()
	// if err != nil {
	// 	log.Fatalf("redis connect error: %v", err)
	// }
	return client
}
