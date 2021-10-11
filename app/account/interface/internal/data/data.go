package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"

	"github.com/go-kratos/kratos/v2/registry"
	userv1 "github.com/lalifeier/vgo/api/account/service/v1"
	"go.opentelemetry.io/otel/trace"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewUserServiceClient, NewUserRepo)

// Data .
type Data struct {
	log *log.Helper
	uc  userv1.UserClient
}

// NewData .
func NewData(logger log.Logger, uc userv1.UserClient) (*Data, error) {
	log := log.NewHelper(log.With(logger, "module", "user-internal/data"))
	return &Data{log: log, uc: uc}, nil
}

func NewUserServiceClient(logger log.Logger, r registry.Discovery, tp trace.TracerProvider) userv1.UserClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///vgo.user.service"),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			logging.Client(logger),
			recovery.Recovery(),
			tracing.Client(tracing.WithTracerProvider(tp)),
		),
	)
	if err != nil {
		panic(err)
	}
	c := userv1.NewUserClient(conn)
	return c
}
