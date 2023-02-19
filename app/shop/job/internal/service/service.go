package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/lalifeier/vvgo-mall/app/shop/job/internal/biz"
	v1 "github.com/lalifeier/vvgo-mall/gen/api/go/shop/job/v1"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewShopJobService)

type ShopJobService struct {
	v1.UnimplementedShopJobServer

	log *log.Helper
	foo *biz.FooUseCase
}

func NewShopJobService(logger log.Logger, foo *biz.FooUseCase) *ShopJobService {
	return &ShopJobService{
		log: log.NewHelper(log.With(logger, "module", "service/shop-job")),
		foo: foo,
	}
}
