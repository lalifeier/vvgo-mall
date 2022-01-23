package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/lalifeier/vvgo-mall/app/account/service/internal/conf"
	"github.com/lalifeier/vvgo-mall/app/account/service/internal/data/ent"

	_ "github.com/go-sql-driver/mysql"

	_ "github.com/lalifeier/vvgo-mall/app/account/service/internal/data/ent/runtime"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewEntClient)

// Data .
type Data struct {
	db  *ent.Client
	log *log.Helper
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper(log.With(logger, "module", "account-service/data"))

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
