package data

import (

	// "github.com/casbin/casbin/v2"
	// gormadapter "github.com/casbin/gorm-adapter/v3"

	"context"
	"time"

	"github.com/Shopify/sarama"
	"github.com/go-redis/redis"
	"github.com/lalifeier/vvgo-mall/app/shop/admin/internal/conf"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	// "gorm.io/driver/mysql"
	// "gorm.io/gorm"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"

	"github.com/lalifeier/vvgo-mall/app/shop/admin/internal/data/ent"

	_ "github.com/go-sql-driver/mysql"

	_ "github.com/lalifeier/vvgo-mall/app/shop/admin/internal/data/ent/runtime"

	consul "github.com/go-kratos/consul/registry"
	consulAPI "github.com/hashicorp/consul/api"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewEntClient, NewRedisClient, NewKafkaProducer, NewKafkaConsumer, NewDiscovery, NewRegistrar, NewDictRepo)

// Data .
type Data struct {
	log *log.Helper
	db  *ent.Client
	// db  *gorm.DB
	rdb *redis.Client
	// casbinEnforcer *casbin.Enforcer
	// mdb *mongo.Database
	kp sarama.AsyncProducer
	kc sarama.Consumer
}

// NewData .
func NewData(entClient *ent.Client, redisClient *redis.Client, producer sarama.AsyncProducer, consumer sarama.Consumer, logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper(log.With(logger, "module", "shop-admin/data"))
	// casbinEnforcer, err := NewCasbinEnforcer(c.CasbinModelPath, db)

	d := &Data{
		log: log,
		db:  entClient,
		rdb: redisClient,
		kp:  producer,
		kc:  consumer,
		// mdb: database,
	}

	cleanup := func() {
		if err := d.db.Close(); err != nil {
			log.Error(err)
		}
		d.kp.Close()
		d.kc.Close()
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

// func NewSysServiceClient(r registry.Discovery) sysV1.SysClient {
// 	conn, err := grpc.DialInsecure(
// 		context.Background(),
// 		// grpc.WithEndpoint("discovery:///vvgo-mall.sys.service"),
// 		grpc.WithEndpoint("127.0.0.1:9002"),
// 		grpc.WithDiscovery(r),
// 		grpc.WithMiddleware(
// 			recovery.Recovery(),
// 		),
// 	)
// 	if err != nil {
// 		panic(err)
// 	}
// 	c := sysV1.NewSysClient(conn)
// 	return c
// }

// func NewDbClient(conf *conf.Data, logger log.Logger) *gorm.DB {
// 	log := log.NewHelper(log.With(logger, "module", "shop-admin/data/gorm"))

// 	db, err := gorm.Open(mysql.Open(conf.Database.Source), &gorm.Config{})
// 	if err != nil {
// 		log.Fatalf("failed opening connection to mysql: %v", err)
// 	}
// 	db.Debug()

// 	// if err := db.AutoMigrate(&Order{}); err != nil {
// 	// 	log.Fatal(err)
// 	// }
// 	return db
// }

func NewEntClient(conf *conf.Data, logger log.Logger) *ent.Client {
	log := log.NewHelper(log.With(logger, "module", "shop-admin/data/ent"))

	client, err := ent.Open(
		conf.Database.Driver,
		conf.Database.Source,
		ent.Debug(),
	)
	if err != nil {
		log.Fatalf("failed opening connection to db: %v", err)
	}
	// Run the auto migration tool.
	// if err := client.Schema.Create(context.Background(), migrate.WithDropIndex(true), migrate.WithDropColumn(true)); err != nil {
	// 	log.Fatalf("failed creating schema resources: %v", err)
	// }
	return client
}

func NewRedisClient(conf *conf.Data, logger log.Logger) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:         conf.Redis.Network,
		Password:     "",
		DB:           0,
		ReadTimeout:  conf.Redis.ReadTimeout.AsDuration(),
		WriteTimeout: conf.Redis.WriteTimeout.AsDuration(),
		PoolSize:     10,
	})
	err := client.Ping().Err()
	if err != nil {
		log.Fatalf("redis connect error: %v", err)
	}
	return client
}

// func NewCasbinEnforcer(rbacModelPath string, db *gorm.DB) (*casbin.Enforcer, error) {
// 	a, err := gormadapter.NewAdapterByDB(db)
// 	if err != nil {
// 		return nil, err
// 	}
// 	e, err := casbin.NewEnforcer(rbacModelPath, a)
// 	if err != nil {
// 		return nil, err
// 	}
// 	err = e.LoadPolicy()
// 	if err != nil {
// 		return nil, err
// 	}
// 	return e, nil
// }

func NewMongo(conf *conf.Data) *mongo.Database {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(conf.Mongodb.Uri))
	if err != nil {
		panic(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}
	return client.Database(conf.Mongodb.Database)
}

func NewKafkaProducer(conf *conf.Data) sarama.AsyncProducer {
	c := sarama.NewConfig()
	p, err := sarama.NewAsyncProducer(conf.Kafka.Addrs, c)
	if err != nil {
		panic(err)
	}
	return p
}

func NewKafkaConsumer(conf *conf.Data) sarama.Consumer {
	c := sarama.NewConfig()
	p, err := sarama.NewConsumer(conf.Kafka.Addrs, c)
	if err != nil {
		panic(err)
	}
	return p
}
