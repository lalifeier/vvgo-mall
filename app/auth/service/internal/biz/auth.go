package biz

import (
	"github.com/go-kratos/kratos/v2/log"
)

type AuthUseCase struct {
	log *log.Helper
}

func NewAuthUseCase(logger log.Logger) *AuthUseCase {
	return &AuthUseCase{log: log.NewHelper(logger)}
}
