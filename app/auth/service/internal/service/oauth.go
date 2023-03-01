package service

import (
	"context"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	// khttp "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-session/session"
	kgin "github.com/lalifeier/vvgo-mall/pkg/bootstrap/server/gin"
)

type LoginForm struct {
	UserName string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func (s *AuthService) Login(ctx *gin.Context) {
	var form LoginForm
	err := ctx.ShouldBind(&form)
	if err != nil {
		kgin.Error(ctx, errors.BadRequest("params_error", err.Error()))
		return
	}

	store, err := session.Start(context.Background(), ctx.Writer, ctx.Request)
	if err != nil {
		kgin.Error(ctx, errors.InternalServer("session_error", "session error"))
		return
	}

	if !(form.UserName == "test" && form.Password == "test") {
		kgin.Error(ctx, errors.Unauthorized("auth_error", "username or password error"))
	}

	store.Set("LoggedInUserID", "test")
	store.Save()

	ctx.Redirect(http.StatusFound, "/auth")
}

func (s *AuthService) Auth(ctx *gin.Context) {
	store, err := session.Start(context.Background(), ctx.Writer, ctx.Request)
	if err != nil {
		kgin.Error(ctx, errors.InternalServer("session_error", "session error"))
		return
	}

	if _, ok := store.Get("LoggedInUserID"); !ok {
		ctx.Redirect(http.StatusFound, "/login")
		return
	}

	ctx.HTML(http.StatusOK, "auth.html", gin.H{})
}

// // khttp://localhost:8002/oauth/authorize?client_id=test_client_1&response_type=code&scope=all&state=xyz&redirect_uri=khttp://localhost:8002/cb
// func (s *AuthService) Authorize(ctx khttp.Context) error {
// 	return s.server.HandleAuthorizeRequest(ctx.Response(), ctx.Request())
// }

func (s *AuthService) Authorize(ctx *gin.Context) {
	store, err := session.Start(ctx, ctx.Writer, ctx.Request)
	if err != nil {
		kgin.Error(ctx, errors.InternalServer("session_error", "session error"))
		return
	}
	var form url.Values
	if v, ok := store.Get("ReturnUri"); ok {
		form = v.(url.Values)
	}
	ctx.Request.Form = form
	store.Delete("ReturnUri")
	store.Save()

	err = s.server.HandleAuthorizeRequest(ctx.Writer, ctx.Request)
	if err != nil {
		kgin.Error(ctx, errors.Unauthorized("auth_error", err.Error()))
		return
	}
}

// // khttp://localhost:8002/oauth/token?grant_type=authorization_code&code=OWU1NMFLMJMTYTC4ZI0ZNTC5LTG5NDUTODG1NTQXNMRKNWRI&redirect_uri=khttp://localhost:8002/cb
// // khttp://localhost:8002/oauth/token?grant_type=refresh_token&refresh_token=M2JLYMNKYJETNJVIOS01YTG3LTK3N2MTNZGZY2JINZNLYZY5
// // basic auth
// // - username: `client_id`
// // - password: `client_secret`
//
//	func (s *AuthService) Token(ctx khttp.Context) error {
//		return s.server.HandleTokenRequest(ctx.Response(), ctx.Request())
//	}
func (s *AuthService) Token(ctx *gin.Context) {
	err := s.server.HandleTokenRequest(ctx.Writer, ctx.Request)
	if err != nil {
		kgin.Error(ctx, errors.Unauthorized("auth_error", err.Error()))
		return
	}
}

// // khttp://localhost:8002/oauth/verify
// // Bearer Token
// // - Token: `access_token`
// func (s *AuthService) Verify(ctx khttp.Context) error {
// 	data, err := s.server.ValidationBearerToken(ctx.Request())
// 	if err != nil {
// 		return err
// 	}
// 	return ctx.JSON(200, data)
// }

// // khttp://localhost:8002/auth/logout
//
//	func (s *AuthService) Logout(ctx khttp.Context) error {
//		accessToken := ctx.Header().Get("Authorization")
//		return s.server.Manager.RemoveAccessToken(ctx, accessToken)
//	}
func (s *AuthService) Logout(ctx *gin.Context) {

}
