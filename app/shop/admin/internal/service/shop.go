package service

import (
	"github.com/go-kratos/kratos/v2/log"
	pb "github.com/lalifeier/vvgo-mall/api/shop/admin/v1"
	"github.com/lalifeier/vvgo-mall/app/shop/admin/internal/biz"
)

type ShopService struct {
	pb.UnimplementedShopServer

	log *log.Helper

	userUsecase *biz.UserUsecase
	roleUsecase *biz.RoleUsecase

	authUsecase     *biz.AuthUsecase
	dictDataUsecase *biz.DictDataUsecase
}

func NewShopService(logger log.Logger, userUsecase *biz.UserUsecase, roleUsecase *biz.RoleUsecase, authUsecase *biz.AuthUsecase, dictDataUsecase *biz.DictDataUsecase) *ShopService {
	return &ShopService{
		log: log.NewHelper(log.With(logger, "module", "shop-admin/service")),

		userUsecase: userUsecase,
		roleUsecase: roleUsecase,

		authUsecase:     authUsecase,
		dictDataUsecase: dictDataUsecase,
	}
}
