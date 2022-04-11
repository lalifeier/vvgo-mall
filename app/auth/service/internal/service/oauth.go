package service

import (
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-oauth2/oauth2/v4"
)

//  http://localhost:8002/auth/authorize?client_id=test_client_1&response_type=code&scope=all&state=xyz&redirect_uri=http://localhost:8002/cb
func (s *AuthService) Authorize(ctx http.Context) error {
	return s.server.HandleAuthorizeRequest(ctx.Response(), ctx.Request())
}

// http://localhost:8002/auth/token?grant_type=client_credentials&client_id=1&client_secret=test&scope=read
func (s *AuthService) Token(ctx http.Context) error {
	return s.server.HandleTokenRequest(ctx.Response(), ctx.Request())
}

// http://localhost:8002/auth/verify
func (s *AuthService) Verify(ctx http.Context) error {
	data, err := s.server.ValidationBearerToken(ctx.Request())
	if err != nil {
		return err
	}
	return ctx.JSON(200, data)
}

// http://localhost:8002/auth/refresh_token?refresh_token=NJCZOTU0MDYTMZHMOC0ZOTAWLTHLNDYTY2JHMJFHY2RLMTGZ
func (s *AuthService) RefreshToken(ctx http.Context) error {
	refreshToken := ctx.Query().Get("refresh_token")
	ti, err := s.server.Manager.LoadRefreshToken(ctx, refreshToken)
	if err != nil {
		return err
	}
	tokenInfo, err := s.server.GetAccessToken(ctx, "refresh_token", &oauth2.TokenGenerateRequest{
		ClientID:            ti.GetClientID(),
		UserID:              ti.GetUserID(),
		RedirectURI:         ti.GetRedirectURI(),
		Scope:               ti.GetScope(),
		Code:                ti.GetCode(),
		CodeChallenge:       ti.GetCodeChallenge(),
		CodeChallengeMethod: ti.GetCodeChallengeMethod(),
		Refresh:             ti.GetRefresh(),
		AccessTokenExp:      ti.GetAccessExpiresIn(),
		Request:             ctx.Request(),
	})
	if err != nil {
		return err
	}
	return ctx.JSON(200, s.server.GetTokenData(tokenInfo))
}

// http://localhost:8002/auth/logout
func (s *AuthService) Logout(ctx http.Context) error {
	accessToken := ctx.Header().Get("Authorization")
	return s.server.Manager.RemoveAccessToken(ctx, accessToken)
}
