package jwtauth

import (
	"errors"

	"github.com/golang-jwt/jwt"
)

const defaultKey = "vgo"

var defaultOptions = options{
	expired:       2 * 60 * 60,
	signingMethod: jwt.SigningMethodHS512,
	signingKey:    []byte(defaultKey),
	keyfunc: func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token")
		}
		return []byte(defaultKey), nil
	},
}

type options struct {
	signingMethod jwt.SigningMethod
	signingKey    interface{}
	keyfunc       jwt.Keyfunc
	expired       int
}

type Option func(*options)

func WithSigningMethod(method jwt.SigningMethod) Option {
	return func(o *options) {
		o.signingMethod = method
	}
}

func WithSigningKey(key interface{}) Option {
	return func(o *options) {
		o.signingKey = key
	}
}

func WithKeyfunc(keyFunc jwt.Keyfunc) Option {
	return func(o *options) {
		o.keyfunc = keyFunc
	}
}

func WithExpired(expired int) Option {
	return func(o *options) {
		o.expired = expired
	}
}
