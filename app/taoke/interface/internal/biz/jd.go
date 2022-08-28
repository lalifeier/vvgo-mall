package biz

import (
	"github.com/lalifeier/vvgo-mall/app/taoke/interface/internal/conf"
)

type JDUseCase struct {
	key string
}

func NewJDUseCase(conf *conf.Auth) *JDUseCase {
	return &JDUseCase{
		key: conf.ApiKey,
	}
}
