package server

import (
	"github.com/gorilla/handlers"
	v1 "github.com/lalifeier/vvgo-mall/api/shop/admin/v1"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/swagger-api/openapiv2"
	"github.com/lalifeier/vvgo-mall/app/shop/admin/internal/conf"
	"github.com/lalifeier/vvgo-mall/app/shop/admin/internal/service"
)

// func NewWhiteListMatcher() selector.MatchFunc {
// 	whiteList := make(map[string]struct{})
// 	// whiteList["/shop.interface.v1.ShopInterface/Login"] = struct{}{}
// 	// whiteList["/shop.interface.v1.ShopInterface/Register"] = struct{}{}
// 	return func(operation string) bool {
// 		if _, ok := whiteList[operation]; ok {
// 			return false
// 		}
// 		return true
// 	}
// }

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, logger log.Logger, shopService *service.ShopService, accountService *service.AccountService, authService *service.AuthService) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			// selector.Server(
			// 	jwt.Server(func(token *jwt2.Token) (interface{}, error) {
			// 		return []byte(ac.ApiKey), nil
			// 	}, jwt.WithSigningMethod(jwt2.SigningMethodHS256)),
			// ).
			// 	Match(NewWhiteListMatcher()).
			// 	Build(),
		),
		http.Filter(handlers.CORS(
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

	openAPIhandler := openapiv2.NewHandler()
	srv.HandlePrefix("/q/", openAPIhandler)

	v1.RegisterShopHTTPServer(srv, shopService)
	v1.RegisterAccountHTTPServer(srv, accountService)
	v1.RegisterAuthHTTPServer(srv, authService)
	return srv
}
