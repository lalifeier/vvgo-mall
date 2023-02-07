package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/lalifeier/vvgo-mall/app/account/service/internal/biz"
	pb "github.com/lalifeier/vvgo-mall/gen/api/go/account/service/v1"
)

type AccountService struct {
	pb.UnimplementedAccountServer
	log *log.Helper

	accountUserUseCase *biz.AccountUserUseCase

	authUseCase *biz.AuthUseCase
}

func NewAccountService(logger log.Logger, accountUserUseCase *biz.AccountUserUseCase, authUseCase *biz.AuthUseCase) *AccountService {
	return &AccountService{
		log:                log.NewHelper(log.With(logger, "module", "account-service/service")),
		accountUserUseCase: accountUserUseCase,
		authUseCase:        authUseCase,
	}
}
