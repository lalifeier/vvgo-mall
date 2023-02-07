package data

import (

	// "github.com/casbin/casbin/v2"
	// gormadapter "github.com/casbin/gorm-adapter/v3"

	"context"
	"database/sql"
	"time"

	"entgo.io/ent/dialect"
	"github.com/Shopify/sarama"
	"github.com/go-redis/redis"
	"github.com/lalifeier/vvgo-mall/app/shop/admin/internal/conf"

	// "go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/mongo/readpref"

	// "gorm.io/driver/mysql"
	// "gorm.io/gorm"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"

	"github.com/google/wire"

	"github.com/lalifeier/vvgo-mall/app/shop/admin/internal/data/ent"

	_ "github.com/go-sql-driver/mysql"

	"github.com/lalifeier/vvgo-mall/app/shop/admin/internal/data/ent/migrate"
	_ "github.com/lalifeier/vvgo-mall/app/shop/admin/internal/data/ent/runtime"

	entsql "entgo.io/ent/dialect/sql"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/persist"
	entadapter "github.com/casbin/ent-adapter"
	entadapterent "github.com/casbin/ent-adapter/ent"
	consul "github.com/go-kratos/consul/registry"
	consulAPI "github.com/hashicorp/consul/api"

	accountv1 "github.com/lalifeier/vvgo-mall/gen/api/go/account/service/v1"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData, NewEntClient, NewRedisClient,
	NewKafkaProducer, NewKafkaConsumer,
	NewDiscovery, NewRegistrar,
	NewEnforcer,
	NewUserRepo, NewRoleRepo, NewDictDataRepo,
	NewAccountServiceClient, NewAccountRepo,
)

// Data .
type Data struct {
	log *log.Helper
	db  *ent.Client
	// db  *gorm.DB
	rdb *redis.Client
	// mdb *mongo.Database
	kp sarama.AsyncProducer
	kc sarama.Consumer

	enforcer *casbin.Enforcer

	ac accountv1.AccountClient
}

// NewData .
func NewData(logger log.Logger,
	entClient *ent.Client, redisClient *redis.Client,
	enforcer *casbin.Enforcer,
	producer sarama.AsyncProducer, consumer sarama.Consumer,
	ac accountv1.AccountClient,
) (*Data, func(), error) {
	log := log.NewHelper(log.With(logger, "module", "shop-admin/data"))

	d := &Data{
		log: log,
		db:  entClient,
		rdb: redisClient,
		kp:  producer,
		kc:  consumer,
		// mdb: database,
		enforcer: enforcer,

		ac: ac,
	}

	cleanup := func() {
		if err := d.db.Close(); err != nil {
			log.Error(err)
		}
		d.rdb.Close()
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

func NewDB() *sql.DB {
	db, err := sql.Open("", "")
	if err != nil {
		log.Fatalf("connect mysql failed err : %v", err)
	}
	return db
}

// func NewEntClient(db *sql.DB) *ent.Client {
// 	drv := entsql.OpenDB(dialect.Postgres, db)
// 	client := ent.NewClient(ent.Driver(drv))

// 	if err := client.Debug().Schema.Create(
// 		context.Background(),
// 		schema.WithDropColumn(true), // 迁移可以删除字段
// 		schema.WithDropIndex(true),  // 迁移可以删除索引
// 	); err != nil {
// 		log.Fatalf("failed creating schema resources: %v", err)
// 	}

// 	return client
// }

func NewEntAdapter(db *sql.DB) *entadapter.Adapter {
	drv := entsql.OpenDB(dialect.MySQL, db)
	client := entadapterent.NewClient(entadapterent.Driver(drv))

	adapter, err := entadapter.NewAdapterWithClient(client)
	if err != nil {
		panic(err)
	}
	return adapter
}

func NewEnforcer(c *conf.Data) *casbin.Enforcer {
	a, err := entadapter.NewAdapter(c.Database.Driver, c.Database.Source)
	if err != nil {
		panic(err)
	}
	e, err := casbin.NewEnforcer("rbac_model.conf", a)
	if err != nil {
		panic(err)
	}

	err = e.LoadPolicy()
	if err != nil {
		panic(err)
	}

	e.EnableAutoSave(true)
	return e
}

func NewCasbinEnforcer(c *conf.Casbin, adapter persist.Adapter) (*casbin.SyncedEnforcer, func(), error) {
	if c.ModelPath == "" {
		return new(casbin.SyncedEnforcer), nil, nil
	}

	e, err := casbin.NewSyncedEnforcer(c.ModelPath)
	if err != nil {
		return nil, nil, err
	}
	e.EnableLog(c.Debug)

	err = e.InitWithModelAndAdapter(e.GetModel(), adapter)
	if err != nil {
		return nil, nil, err
	}
	e.EnableEnforce(c.Enable)

	cleanup := func() {}
	if c.AutoLoad {
		cleanup = func() {
			e.StopAutoLoadPolicy()
		}
		e.StartAutoLoadPolicy(time.Duration(c.AutoLoadInternal) * time.Second)
	}

	return e, cleanup, nil
}

// func NewMongo(conf *conf.Data) *mongo.Database {
// 	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
// 	client, err := mongo.Connect(ctx, options.Client().ApplyURI(conf.Mongodb.Uri))
// 	if err != nil {
// 		panic(err)
// 	}
// 	err = client.Ping(ctx, readpref.Primary())
// 	if err != nil {
// 		panic(err)
// 	}
// 	return client.Database(conf.Mongodb.Database)
// }

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

func NewAccountServiceClient(r registry.Discovery, tp *tracesdk.TracerProvider) accountv1.AccountClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///vvgo-mall.account.service"),
		// grpc.WithEndpoint("127.0.0.1:9002"),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			tracing.Client(tracing.WithTracerProvider(tp)),
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	c := accountv1.NewAccountClient(conn)
	return c
}
