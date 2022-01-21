package data

import (
	"context"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/go-redis/redis"
	"github.com/lalifeier/vvgo-mall/app/shop/admin/internal/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"

	sysV1 "github.com/lalifeier/vvgo-mall/api/sys/service/v1"

	consul "github.com/go-kratos/consul/registry"
	consulAPI "github.com/hashicorp/consul/api"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewDiscovery, NewSysServiceClient)

// Data .
type Data struct {
	log       *log.Helper
	SysClient sysV1.SysClient
	db        *gorm.DB
	rdb       *redis.Client
	// casbinEnforcer *casbin.Enforcer
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, sysClient sysV1.SysClient) (*Data, func(), error) {
	log := log.NewHelper(log.With(logger, "module", "shop-admin/data"))

	dbClient := NewDbClient(c, logger)
	redisClient := NewRedisClient(c, logger)
	// casbinEnforcer, err := NewCasbinEnforcer(c.CasbinModelPath, db)

	cleanup := func() {
		log.Info("closing the data resources")
	}
	return &Data{
		log:       log,
		SysClient: sysClient,
		db:        dbClient,
		rdb:       redisClient,
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

func NewSysServiceClient(r registry.Discovery) sysV1.SysClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		// grpc.WithEndpoint("discovery:///vvgo-mall.sys.service"),
		grpc.WithEndpoint("127.0.0.1:9002"),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	c := sysV1.NewSysClient(conn)
	return c
}

func NewDbClient(conf *conf.Data, logger log.Logger) *gorm.DB {
	log := log.NewHelper(log.With(logger, "module", "shop-admin/data/gorm"))

	db, err := gorm.Open(mysql.Open(conf.Database.Source), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	db.Debug()

	// if err := db.AutoMigrate(&Order{}); err != nil {
	// 	log.Fatal(err)
	// }
	return db
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

func NewCasbinEnforcer(rbacModelPath string, db *gorm.DB) (*casbin.Enforcer, error) {
	a, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		return nil, err
	}
	e, err := casbin.NewEnforcer(rbacModelPath, a)
	if err != nil {
		return nil, err
	}
	err = e.LoadPolicy()
	if err != nil {
		return nil, err
	}
	return e, nil
}
