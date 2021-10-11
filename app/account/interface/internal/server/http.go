package server

import (
	"github.com/go-kratos/kratos/v2/log"
	// "github.com/go-kratos/kratos/v2/middleware/jwt"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/swagger-api/openapiv2"
	"github.com/gorilla/handlers"
	v1 "github.com/lalifeier/vgo/api/account/interface/v1"
	"github.com/lalifeier/vgo/app/account/interface/internal/conf"
	"github.com/lalifeier/vgo/app/account/interface/internal/service"
	"github.com/lalifeier/vgo/app/account/interface/middleware/metrics"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/otel/trace"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, logger log.Logger, tp trace.TracerProvider, user *service.UserService) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			// Recover
			recovery.Recovery(),

			// Access logger
			logging.Server(logger),

			// Trace ID
			tracing.Server(
				tracing.WithTracerProvider(tp)),

			// Validate
			validate.Validator(),

			// Metrics
			metrics.NewServer(),

			// Jwt
			// jwt.Server(),
		),

		http.Filter(
			// CORS
			handlers.CORS(
				handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
				handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
				handlers.AllowedOrigins([]string{"*"}),
			)),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)

	// Swagger
	openAPIhandler := openapiv2.NewHandler()
	srv.HandlePrefix("/q/", openAPIhandler)

	srv.Handle("/metrics", promhttp.Handler())

	v1.RegisterUserInterfaceHTTPServer(srv, user)
	return srv
}
