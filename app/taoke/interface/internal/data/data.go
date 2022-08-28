package data

import (
	"github.com/go-redis/redis"
	"github.com/lalifeier/vvgo-mall/app/taoke/interface/internal/conf"
	"github.com/lalifeier/vvgo-mall/pkg/dataoke"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"

	"github.com/google/wire"

	"github.com/lalifeier/vvgo-mall/app/taoke/interface/internal/data/ent"

	_ "github.com/go-sql-driver/mysql"

	_ "github.com/lalifeier/vvgo-mall/app/taoke/interface/internal/data/ent/runtime"

	consul "github.com/go-kratos/consul/registry"
	consulAPI "github.com/hashicorp/consul/api"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData, NewEntClient, NewRedisClient,
	NewDiscovery, NewRegistrar,
	NewDaTaoKeClient,
	// NewTaoBaoRepo,
)

// Data .
type Data struct {
	log *log.Helper
	db  *ent.Client
	rdb *redis.Client

	// dtk *dataoke.Client
}

// NewData .
func NewData(logger log.Logger,
	entClient *ent.Client, redisClient *redis.Client,
	// daTaoKeClient *dataoke.Client,
) (*Data, func(), error) {
	log := log.NewHelper(log.With(logger, "module", "taoke-interface/data"))

	d := &Data{
		log: log,
		db:  entClient,
		rdb: redisClient,
		// dtk: daTaoKeClient,
	}

	cleanup := func() {
		if err := d.db.Close(); err != nil {
			log.Error(err)
		}
		d.rdb.Close()
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
	log := log.NewHelper(log.With(logger, "module", "taoke-interface/data/ent"))

	client, err := ent.Open(
		conf.Database.Driver,
		conf.Database.Source,
		ent.Debug(),
	)
	if err != nil {
		log.Fatalf("failed opening connection to db: %v", err)
	}
	// Run the auto migration tool.
	// if err := client.Debug().Schema.Create(
	// 	context.Background(),
	// 	migrate.WithDropIndex(true),
	// 	migrate.WithDropColumn(true),
	// ); err != nil {
	// 	log.Fatalf("failed creating schema resources: %v", err)
	// }
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

func NewDaTaoKeClient(c *conf.Data, logger log.Logger) *dataoke.Client {
	client := dataoke.NewClient(c.DaTaoKe.AppKey, c.DaTaoKe.AppSecret)
	return client
}
