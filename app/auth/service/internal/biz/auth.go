package biz

import (
	"github.com/go-kratos/kratos/v2/log"
)

type AuthUsecase struct {
	log *log.Helper
}

func NewAuthUsecase(logger log.Logger) *AuthUsecase {
	return &AuthUsecase{log: log.NewHelper(logger)}
}
