package casbin

import (
	"context"

	"github.com/casbin/casbin/v2"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	transporthttp "github.com/go-kratos/kratos/v2/transport/http"
)

const (
	reason string = "UNAUTHORIZED"
)

var (
	ErrNoPermission = errors.Forbidden(reason, "no permission")
)

func Server(enforcer *casbin.Enforcer) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if tr, ok := transport.FromServerContext(ctx); ok {
				if tr.Kind() == transport.KindHTTP {
					h := tr.(*transporthttp.Transport)

					// enforcer.AddRoleForUser("test1", "test")
					// enforcer.AddRoleForUser("test2", "admin")
					enforcer.AddPermissionForUser("admin", "/api/v1/register", "POST")
					enforcer.AddPermissionForUser("admin", "/api/v1/login", "POST")
					obj := h.Request().URL.Path
					act := h.Request().Method
					sub := "admin"
					println(sub, obj, act)
					if ok, _ := enforcer.Enforce(sub, obj, act); !ok {
						return nil, ErrNoPermission
					}
				}
			}
			return handler(ctx, req)
		}
	}
}
