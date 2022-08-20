package service

import (
	"github.com/go-kratos/kratos/v2/log"
	pb "github.com/lalifeier/vvgo-mall/api/shop/admin/v1"
	"github.com/lalifeier/vvgo-mall/app/shop/admin/internal/biz"
)

type ShopService struct {
	pb.UnimplementedShopAdminServer

	log *log.Helper

	userUseCase *biz.UserUseCase
	roleUseCase *biz.RoleUseCase

	authUseCase     *biz.AuthUseCase
	dictDataUseCase *biz.DictDataUseCase
}

func NewShopService(logger log.Logger, userUseCase *biz.UserUseCase, roleUseCase *biz.RoleUseCase, authUseCase *biz.AuthUseCase, dictDataUseCase *biz.DictDataUseCase) *ShopService {
	return &ShopService{
		log: log.NewHelper(log.With(logger, "module", "shop-admin/service")),

		userUseCase: userUseCase,
		roleUseCase: roleUseCase,

		authUseCase:     authUseCase,
		dictDataUseCase: dictDataUseCase,
	}
}
