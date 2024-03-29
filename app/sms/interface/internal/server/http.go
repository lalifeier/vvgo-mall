package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/lalifeier/vvgo-mall/app/sms/interface/internal/conf"
	"github.com/lalifeier/vvgo-mall/app/sms/interface/internal/service"
	v1 "github.com/lalifeier/vvgo-mall/gen/api/go/sms/interface/v1"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, smsServer *service.SmsService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			validate.Validator(),
		),
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
	v1.RegisterSmsHTTPServer(srv, smsServer)
	return srv
}
