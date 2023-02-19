package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/lalifeier/vvgo-mall/app/shop/job/internal/biz"
)

var _ biz.FooRepo = (*fooRepo)(nil)

type fooRepo struct {
	data *Data
	log  *log.Helper
}

func NewFooRepo(data *Data, logger log.Logger) biz.FooRepo {
	return &fooRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "sensor/repo/logger-service")),
	}
}
