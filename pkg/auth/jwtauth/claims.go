package jwtauth

import "github.com/golang-jwt/jwt"

type UserClaims struct {
	UserID   int64
	UserName string
	jwt.StandardClaims
}
