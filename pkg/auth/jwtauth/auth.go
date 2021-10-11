package jwtauth

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type TokenInfo interface {
	GetAccessToken() string

	GetRefreshToken() string

	// GetTokenType() string

	// GetExpiresAt() int64
}

type Auth interface {
	GenerateAccessToken(UserID int64, UserName string) (string, error)

	GenerateRefreshToken(tokenString string, expire time.Duration) (string, error)

	ParseToken(tokenString string) error

	DestroyToken(tokenString string) error
}

func New(opts ...Option) *JwtAuth {
	o := defaultOptions
	for _, opt := range opts {
		opt(&o)
	}

	return &JwtAuth{
		opts: &o,
	}
}

type JwtAuth struct {
	opts *options
}

func (a *JwtAuth) GenerateUserClaims(userId int64, userName string) *UserClaims {
	now := time.Now()
	expiresAt := now.Add(time.Duration(a.opts.expired)).Unix()

	return &UserClaims{
		userId,
		userName,
		jwt.StandardClaims{
			NotBefore: now.Unix(),
			IssuedAt:  now.Unix(),
			ExpiresAt: expiresAt,
		},
	}
}

func (a *JwtAuth) GenerateToken(userId int64, userName string) (string, error) {
	userClaims := a.GenerateUserClaims(userId, userName)
	return jwt.NewWithClaims(a.opts.signingMethod, userClaims).SignedString(a.opts.signingKey)
}

func (a *JwtAuth) ParseToken(tokenString string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, a.opts.keyfunc)
	if err != nil || !token.Valid {
		return nil, err
	}
	return token.Claims.(*UserClaims), nil
}

// func (t *token) RefreshToken(tokenString string, expire time.Duration) (string, error) {
// 	claims, err := t.ParseToken(tokenString)
// 	if err != nil {
// 		return "", err
// 	}
// 	jwt.TimeFunc = func() time.Time {
// 		return time.Unix(0, 0)
// 	}
// 	jwt.TimeFunc = time.Now
// 	claims.StandardClaims.ExpiresAt = time.Now().Add(expire).Unix()
// 	return t.GenerateToken(claims)
// }
