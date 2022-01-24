package service

import (
	"github.com/go-kratos/kratos/v2/log"
	pb "github.com/lalifeier/vvgo-mall/api/account/service/v1"
	"github.com/lalifeier/vvgo-mall/app/account/service/internal/biz"
)

type AccountService struct {
	pb.UnimplementedAccountServer
	log *log.Helper

	accountUserUsecase *biz.AccountUserUsecase
	staffUsecase       *biz.StaffUsecase
}

func NewAccountService(logger log.Logger, accountUserUsecase *biz.AccountUserUsecase, staffUsecase *biz.StaffUsecase) *AccountService {
	return &AccountService{
		log:                log.NewHelper(log.With(logger, "module", "account-service/service")),
		accountUserUsecase: accountUserUsecase,
		staffUsecase:       staffUsecase,
	}
}
