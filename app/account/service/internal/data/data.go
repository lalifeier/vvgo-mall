package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
	"github.com/lalifeier/vvgo-mall/app/account/service/internal/data/ent"
	"github.com/lalifeier/vvgo-mall/gen/api/go/common/conf"
	"github.com/lalifeier/vvgo-mall/pkg/bootstrap"

	_ "github.com/lib/pq"

	"github.com/lalifeier/vvgo-mall/app/account/service/internal/data/ent/migrate"
	_ "github.com/lalifeier/vvgo-mall/app/account/service/internal/data/ent/runtime"

	"github.com/go-redis/redis/extra/redisotel/v8"
	"github.com/go-redis/redis/v8"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,

	NewDiscovery,

	NewEntClient,
	NewRedisClient,

	NewAccountUserRepo,
)

// Data .
type Data struct {
	log *log.Helper
	db  *ent.Client
	rdb *redis.Client
}

// NewData .
func NewData(logger log.Logger, entClient *ent.Client, redisClient *redis.Client) (*Data, func(), error) {
	log := log.NewHelper(log.With(logger, "module", "account-service/data"))

	d := &Data{
		db:  entClient,
		rdb: redisClient,
		log: log,
	}

	cleanup := func() {
		log.Info("closing the data resources")

		if err := d.db.Close(); err != nil {
			log.Error(err)
		}
		if err := d.rdb.Close(); err != nil {
			log.Error(err)
		}
	}
	return d, cleanup, nil
}

func NewEntClient(conf *conf.Bootstrap, logger log.Logger) *ent.Client {
	log := log.NewHelper(log.With(logger, "module", "account-service/data/ent"))

	client, err := ent.Open(
		conf.Data.Database.Driver,
		conf.Data.Database.Source,
		// ent.Debug(),
	)
	if err != nil {
		log.Fatalf("failed opening connection to db: %v", err)
	}
	// Run the auto migration tool.
	if conf.Data.Database.Migrate {
		if err := client.Schema.Create(context.Background(), migrate.WithForeignKeys(false)); err != nil {
			log.Fatalf("failed creating schema resources: %v", err)
		}
	}
	return client
}

func NewRedisClient(conf *conf.Bootstrap, logger log.Logger) *redis.Client {
	log := log.NewHelper(log.With(logger, "module", "account-service/data/redis"))

	client := redis.NewClient(&redis.Options{
		Addr:         conf.Data.Redis.Network,
		Password:     conf.Data.Redis.Password,
		DB:           int(conf.Data.Redis.Db),
		DialTimeout:  conf.Data.Redis.DialTimeout.AsDuration(),
		ReadTimeout:  conf.Data.Redis.ReadTimeout.AsDuration(),
		WriteTimeout: conf.Data.Redis.WriteTimeout.AsDuration(),
	})

	if client == nil {
		log.Fatalf("failed opening connection to redis")
	}

	client.AddHook(redisotel.NewTracingHook())

	return client
}

// NewDiscovery 创建服务发现客户端
func NewDiscovery(cfg *conf.Bootstrap) registry.Discovery {
	return bootstrap.NewConsulRegistry(cfg.Registry)
}
