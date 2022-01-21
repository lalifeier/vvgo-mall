package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/lalifeier/vvgo-mall/app/ums/service/internal/conf"
	"github.com/lalifeier/vvgo-mall/app/ums/service/internal/data/ent"

	_ "github.com/go-sql-driver/mysql"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewEntClient, NewAccountUserRepo)

// Data .
type Data struct {
	db  *ent.Client
	log *log.Helper
}

// func NewDB(conf *conf.Data, logger log.Logger) *gorm.DB {
// 	log := log.NewHelper(log.With(logger, "module", "ums-service/data/gorm"))

// 	db, err := gorm.Open(mysql.Open(conf.Database.Source), &gorm.Config{})
// 	if err != nil {
// 		log.Fatalf("failed opening connection to mysql: %v", err)
// 	}

// 	// if err := db.AutoMigrate(&Order{}); err != nil {
// 	// 	log.Fatal(err)
// 	// }
// 	return db
// }

// func NewMongoClient(conf *conf.Data, logger log.Logger) *mongo.Client {
// 	log := log.NewHelper(log.With(logger, "module", "user-service/data/mongo"))

// 	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
// 	client, err := mongo.Connect(context.TODO(), clientOptions)
// 	if err != nil {
// 		log.Fatalf("failed opening connection to mongo: %v", err)
// 	}

// 	err = client.Ping(context.TODO(), nil)
// 	if err != nil {
// 		log.Fatalf("failed opening connection to mongo: %v", err)
// 	}

// 	return client
// }

func NewEntClient(conf *conf.Data, logger log.Logger) *ent.Client {
	log := log.NewHelper(log.With(logger, "module", "ums-service/data/ent"))

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

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper(log.With(logger, "module", "ums-service/data"))

	entClient := NewEntClient(c, logger)
	d := &Data{
		db:  entClient,
		log: log,
	}

	cleanup := func() {
		log.Info("closing the data resources")

		if err := d.db.Close(); err != nil {
			log.Error(err)
		}
	}
	return d, cleanup, nil
}
