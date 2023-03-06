package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/lalifeier/vvgo-mall/app/account/service/internal/service"
	v1 "github.com/lalifeier/vvgo-mall/gen/api/go/account/service/v1"
	"github.com/lalifeier/vvgo-mall/gen/api/go/common/conf"
	"github.com/lalifeier/vvgo-mall/pkg/bootstrap/server"
)

func newHTTPMiddleware(logger log.Logger) []middleware.Middleware {
	var ms []middleware.Middleware

	ms = append(ms, logging.Server(logger))

	return ms
}

// NewHTTPServer new a HTTP server.
func NewHTTPServer(cfg *conf.Bootstrap, logger log.Logger, accountSvc *service.AccountService) *http.Server {
	srv := server.CreateHttpServer(cfg, newHTTPMiddleware(logger)...)

	v1.RegisterAccountServiceHTTPServer(srv, accountSvc)

	return srv
}
