package server

import (
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	kratosHttp "github.com/go-kratos/kratos/v2/transport/http"

	"github.com/gorilla/handlers"

	"github.com/lalifeier/vvgo-mall/gen/api/go/common/conf"
)

// 创建Http服务端
func CreateHttpServer(cfg *conf.Bootstrap, m ...middleware.Middleware) *kratosHttp.Server {
	var opts = []kratosHttp.ServerOption{
		kratosHttp.Filter(handlers.CORS(
			handlers.AllowedHeaders(cfg.Server.Http.Headers),
			handlers.AllowedMethods(cfg.Server.Http.Methods),
			handlers.AllowedOrigins(cfg.Server.Http.Origins),
		)),
	}

	var ms []middleware.Middleware
	ms = append(ms, recovery.Recovery())
	ms = append(ms, tracing.Server())
	ms = append(ms, m...)
	opts = append(opts, kratosHttp.Middleware(ms...))

	if cfg.Server.Http.Network != "" {
		opts = append(opts, kratosHttp.Network(cfg.Server.Http.Network))
	}
	if cfg.Server.Http.Addr != "" {
		opts = append(opts, kratosHttp.Address(cfg.Server.Http.Addr))
	}
	if cfg.Server.Http.Timeout != nil {
		opts = append(opts, kratosHttp.Timeout(cfg.Server.Http.Timeout.AsDuration()))
	}

	return kratosHttp.NewServer(opts...)
}
