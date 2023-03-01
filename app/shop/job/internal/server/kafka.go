package server

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/lalifeier/vvgo-mall/app/shop/job/internal/service"
	"github.com/lalifeier/vvgo-mall/gen/api/go/common/conf"

	"github.com/tx7do/kratos-transport/transport/kafka"
)

// NewKafkaServer create a kafka server.
func NewKafkaServer(c *conf.Server, _ log.Logger, svc *service.ShopJobService) *kafka.Server {
	ctx := context.Background()

	srv := kafka.NewServer(
		kafka.WithAddress(c.Kafka.Addrs),
		kafka.WithCodec("json"),
	)

	registerKafkaSubscribers(ctx, srv, svc)

	return srv
}

func registerKafkaSubscribers(ctx context.Context, srv *kafka.Server, svc *service.ShopJobService) {
	// _ = srv.RegisterSubscriber(ctx,
	// 	"logger.sensor.ts",
	// 	"sensor_logger",
	// 	false,
	// 	// registerSensorDataHandler(svc.InsertSensorData),
	// 	// sensorDataCreator,
	// )
}
