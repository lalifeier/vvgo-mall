package service

import (
	"github.com/go-kratos/kratos/v2/transport/http"
)

// http://localhost:8002/oauth/authorize?client_id=test_client_1&response_type=code&scope=all&state=xyz&redirect_uri=http://localhost:8002/cb
func (s *AuthService) Authorize(ctx http.Context) error {
	return s.server.HandleAuthorizeRequest(ctx.Response(), ctx.Request())
}

// http://localhost:8002/oauth/token?grant_type=authorization_code&code=OWU1NMFLMJMTYTC4ZI0ZNTC5LTG5NDUTODG1NTQXNMRKNWRI&redirect_uri=http://localhost:8002/cb
// http://localhost:8002/oauth/token?grant_type=refresh_token&refresh_token=M2JLYMNKYJETNJVIOS01YTG3LTK3N2MTNZGZY2JINZNLYZY5
// basic auth
// - username: `client_id`
// - password: `client_secret`
func (s *AuthService) Token(ctx http.Context) error {
	return s.server.HandleTokenRequest(ctx.Response(), ctx.Request())
}

// http://localhost:8002/oauth/verify
// Bearer Token
// - Token: `access_token`
func (s *AuthService) Verify(ctx http.Context) error {
	data, err := s.server.ValidationBearerToken(ctx.Request())
	if err != nil {
		return err
	}
	return ctx.JSON(200, data)
}

// http://localhost:8002/auth/logout
func (s *AuthService) Logout(ctx http.Context) error {
	accessToken := ctx.Header().Get("Authorization")
	return s.server.Manager.RemoveAccessToken(ctx, accessToken)
}
