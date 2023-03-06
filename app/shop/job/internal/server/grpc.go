package server

import (
	"github.com/lalifeier/vvgo-mall/app/shop/job/internal/service"
	"github.com/lalifeier/vvgo-mall/gen/api/go/common/conf"
	v1 "github.com/lalifeier/vvgo-mall/gen/api/go/shop/job/v1"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/lalifeier/vvgo-mall/pkg/bootstrap/server"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(cfg *conf.Bootstrap, logger log.Logger, shopJobService *service.ShopJobService) *grpc.Server {
	srv := server.CreateGrpcServer(cfg, logging.Server(logger))
	v1.RegisterShopJobServer(srv, shopJobService)
	return srv
}
