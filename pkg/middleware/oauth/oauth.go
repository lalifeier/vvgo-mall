package casbin

import (
	"context"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	transporthttp "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-oauth2/oauth2/v4/server"
)

const (
	reason string = "UNAUTHORIZED"
)

var (
	ErrNoPermission = errors.Forbidden(reason, "no permission")
)

func Server(srv *server.Server) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if tr, ok := transport.FromServerContext(ctx); ok {
				if tr.Kind() == transport.KindHTTP {
					h := tr.(*transporthttp.Transport)

					r := h.Request()
					_, err := srv.ValidationBearerToken(r)
					if err != nil {
						return nil, ErrNoPermission
					}
				}
			}
			return handler(ctx, req)
		}
	}
}
