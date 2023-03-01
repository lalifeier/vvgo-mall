package main

import (
	"flag"

	"github.com/lalifeier/vvgo-mall/gen/api/go/common/conf"
	"github.com/lalifeier/vvgo-mall/pkg/bootstrap/config"
	FLAG "github.com/lalifeier/vvgo-mall/pkg/bootstrap/flag"
	"github.com/lalifeier/vvgo-mall/pkg/bootstrap/logger"
	"github.com/lalifeier/vvgo-mall/pkg/bootstrap/tracer"
	"github.com/lalifeier/vvgo-mall/pkg/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"

	_ "go.uber.org/automaxprocs"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	Service = service.NewServiceInfo(
		service.ShopJobService,
		"1.0.0",
		"",
	)

	Flags = FLAG.NewCommandFlags()
)

func init() {
	Flags.Init()
}

func newApp(logger log.Logger, gs *grpc.Server, rr registry.Registrar) *kratos.App {
	return kratos.New(
		kratos.ID(Service.GetInstanceId()),
		kratos.Name(Service.Name),
		kratos.Version(Service.Version),
		kratos.Metadata(Service.Metadata),
		kratos.Logger(logger),
		kratos.Server(
			gs,
		),
		kratos.Registrar(rr),
	)
}

func loadConfig() (*conf.Bootstrap, *conf.Registry) {
	c := config.NewConfigProvider(Flags.ConfigType, Flags.ConfigHost, Flags.Conf, Service.Name)

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	var rc conf.Registry
	if err := c.Scan(&rc); err != nil {
		panic(err)
	}

	return &bc, &rc
}

func main() {
	flag.Parse()

	bc, rc := loadConfig()
	if bc == nil || rc == nil {
		panic("load config failed")
	}

	logger := logger.NewLoggerProvider(logger.LoggerTypeStd, bc.Logger, &Service)

	err := tracer.NewTracerProvider(bc.Trace, &Service)
	if err != nil {
		panic(err)
	}

	app, cleanup, err := wireApp(bc.Server, rc, bc.Data, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}
