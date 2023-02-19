package biz

import (
	"github.com/go-kratos/kratos/v2/log"
)

type FooRepo interface {
}

type FooUseCase struct {
	log  *log.Helper
	repo FooRepo
}

func NewFooUseCase(logger log.Logger, repo FooRepo) *FooUseCase {
	l := log.NewHelper(log.With(logger, "module", "biz/usecase/logger-service"))
	return &FooUseCase{log: l, repo: repo}
}
