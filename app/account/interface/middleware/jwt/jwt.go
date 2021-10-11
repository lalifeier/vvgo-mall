package jwt

// import (
// 	"context"
// 	"fmt"
// 	"strings"

// 	"github.com/go-kratos/kratos/v2/errors"
// 	"github.com/go-kratos/kratos/v2/middleware"
// 	"github.com/go-kratos/kratos/v2/transport"
// )

// var (
// 	ErrMissingJwtToken        = errors.Unauthorized("UNAUTHORIZED", "JWT token is missing")
// 	ErrMissingKeyFunc         = errors.Unauthorized("UNAUTHORIZED", "keyFunc is missing")
// 	ErrTokenInvalid           = errors.Unauthorized("UNAUTHORIZED", "Token is invalid")
// 	ErrTokenExpired           = errors.Unauthorized("UNAUTHORIZED", "JWT token has expired")
// 	ErrTokenParseFail         = errors.Unauthorized("UNAUTHORIZED", "Fail to parse JWT token ")
// 	ErrUnSupportSigningMethod = errors.Unauthorized("UNAUTHORIZED", "Wrong signing method")
// 	ErrWrongContext           = errors.Unauthorized("UNAUTHORIZED", "Wrong context for middelware")
// 	ErrNeedTokenProvider      = errors.Unauthorized("UNAUTHORIZED", "Token provider is missing")
// 	ErrSignToken              = errors.Unauthorized("UNAUTHORIZED", "Can not sign token.Is the key correct?")
// 	ErrGetKey                 = errors.Unauthorized("UNAUTHORIZED", "Can not get key while signing token")
// )

// func Server() middleware.Middleware {
// 	return func(handler middleware.Handler) middleware.Handler {
// 		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {

// 			if tr, ok := transport.FromServerContext(ctx); ok {
// 				fmt.Println(tr.Operation())
// 				fmt.Println(tr.RequestHeader())

// 				auths := strings.SplitN(tr.RequestHeader().Get(authorizationKey), " ", 2)
// 				if len(auths) != 2 || !strings.EqualFold(auths[0], bearerWord) {
// 					return nil, ErrMissingJwtToken
// 				}
// 				jwtToken := auths[1]
// 				tokenInfo, err := jwt.Parse(jwtToken, keyFunc)
// 				if err != nil {
// 					if ve, ok := err.(*jwt.ValidationError); ok {
// 						if ve.Errors&jwt.ValidationErrorMalformed != 0 {
// 							return nil, ErrTokenInvalid
// 						} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
// 							return nil, ErrTokenExpired
// 						} else {
// 							return nil, ErrTokenParseFail
// 						}
// 					}
// 				} else if !tokenInfo.Valid {
// 					return nil, ErrTokenInvalid
// 				} else if tokenInfo.Method != o.signingMethod {
// 					return nil, ErrUnSupportSigningMethod
// 				}
// 				ctx = NewContext(ctx, tokenInfo.Claims)
// 				return handler(ctx, req)

// 				// var token string
// 				// if request, ok := req.(http.Request); ok {
// 				// 	fmt.Println(request)
// 				// 	token = request.Header.Get("Authorization")
// 				// } else if md, ok := metadata.FromIncomingContext(ctx); ok {
// 				// 	token = md.Get("Authorization")[0]
// 				// } else {
// 				// 	return nil, errors.ErrAuthFail
// 				// }

// 				// claims, err := jwt.New("").ParseToken(token)
// 				// if err != nil {
// 				// 	return nil, errors.ErrAuthFail
// 				// }
// 				// ctx = context.WithValue(ctx, "x-md-global-uid", claims.UserID)
// 				return handler(ctx, req)
// 			}
// 			return nil, ErrWrongContext
// 		}
// 	}
// }
