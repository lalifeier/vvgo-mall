package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis"
	"github.com/google/wire"
	"github.com/lalifeier/vvgo-mall/app/account/service/internal/conf"
	"github.com/lalifeier/vvgo-mall/app/account/service/internal/data/ent"

	_ "github.com/lib/pq"

	_ "github.com/lalifeier/vvgo-mall/app/account/service/internal/data/ent/runtime"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewEntClient, NewRedisClient, NewAccountUserRepo)

// Data .
type Data struct {
	db  *ent.Client
	rdb *redis.Client
	log *log.Helper
}

// NewData .
func NewData(c *conf.Data, redisClient *redis.Client, logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper(log.With(logger, "module", "account-service/data"))

	entClient := NewEntClient(c, logger)
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

func NewEntClient(conf *conf.Data, logger log.Logger) *ent.Client {
	log := log.NewHelper(log.With(logger, "module", "account-service/data/ent"))

	client, err := ent.Open(
		conf.Database.Driver,
		conf.Database.Source,
		ent.Debug(),
	)
	if err != nil {
		log.Fatalf("failed opening connection to db: %v", err)
	}
	// Run the auto migration tool.
	// if err := client.Schema.Create(context.Background(), migrate.WithForeignKeys(false)); err != nil {
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
