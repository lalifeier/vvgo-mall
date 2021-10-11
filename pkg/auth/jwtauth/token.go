package jwtauth

// var _ Token = (*token)(nil)

// type Token interface {
// 	GenerateToken(claims *UserClaims) (string, error)

// 	ParseToken(tokenString string) (*UserClaims, error)

// 	RefreshToken(tokenString string, expire time.Duration) (string, error)
// }

type tokenInfo struct {
	AccessToken string `json:"access_token"`

	RefreshToken string `json:"refresh_token"`
}

func (t *tokenInfo) GetAccessToken() string {
	return t.AccessToken
}

func (t *tokenInfo) GetRefreshToken() string {
	return t.RefreshToken
}
