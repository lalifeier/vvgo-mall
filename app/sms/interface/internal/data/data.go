package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis"
	"github.com/google/wire"
	"github.com/lalifeier/vvgo-mall/app/sms/interface/internal/conf"
	sms "github.com/lalifeier/vvgo-mall/pkg/sms/aliyun"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewSmsRepo, NewSmsClient, NewRedisClient)

// Data .
type Data struct {
	log       *log.Helper
	smsClient *sms.Client
	rdb       *redis.Client
}

// NewData .
func NewData(logger log.Logger, smsClient *sms.Client, redisClient *redis.Client) (*Data, func(), error) {
	log := log.NewHelper(log.With(logger, "module", "sms-internal/data"))
	cleanup := func() {
		log.Info("closing the data resources")
	}
	return &Data{log: log, smsClient: smsClient, rdb: redisClient}, cleanup, nil
}

func NewSmsClient(c *conf.Sms) *sms.Client {
	client, err := sms.NewClient(c.RegionId, c.AccessKey, c.AccessSecret, c.SignName, c.TemplateCode)
	if err != nil {
		panic(err)
	}
	return client
}

func NewRedisClient(conf *conf.Data, logger log.Logger) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     conf.Redis.Network,
		Password: "",
		DB:       0,
	})
	if err := client.Ping().Err(); err != nil {
		panic(err)
	}
	return client
}
