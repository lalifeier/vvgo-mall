package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/lalifeier/vvgo-mall/app/account/service/internal/service"
	v1 "github.com/lalifeier/vvgo-mall/gen/api/go/account/service/v1"
	"github.com/lalifeier/vvgo-mall/gen/api/go/common/conf"
	"github.com/lalifeier/vvgo-mall/pkg/bootstrap/server"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(cfg *conf.Bootstrap, logger log.Logger, accountSvc *service.AccountService, accountUserSvc *service.AccountUserService) *grpc.Server {
	srv := server.CreateGrpcServer(cfg, logging.Server(logger))

	v1.RegisterAccountServiceServer(srv, accountSvc)
	v1.RegisterAccountUserServiceServer(srv, accountUserSvc)

	return srv
}
