package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/lalifeier/vvgo-mall/app/account/service/internal/biz"
	v1 "github.com/lalifeier/vvgo-mall/gen/api/go/account/service/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

type AccountUserService struct {
	v1.UnimplementedAccountUserServiceServer
	log *log.Helper

	accountUserUseCase *biz.AccountUserUseCase
}

func NewAccountUserService(logger log.Logger, accountUserUseCase *biz.AccountUserUseCase) *AccountUserService {
	return &AccountUserService{
		log:                log.NewHelper(log.With(logger, "module", "account-service/service")),
		accountUserUseCase: accountUserUseCase,
	}
}

func (s *AccountUserService) CreateAccountUser(ctx context.Context, req *v1.CreateAccountUserReq) (*v1.AccountUser, error) {
	return s.accountUserUseCase.CreateAccountUser(ctx, req)
}
func (s *AccountUserService) UpdateAccountUser(ctx context.Context, req *v1.UpdateAccountUserReq) (*v1.AccountUser, error) {
	return s.accountUserUseCase.UpdateAccountUser(ctx, req)
}
func (s *AccountUserService) DeleteAccountUser(ctx context.Context, req *v1.DeleteAccountUserReq) (*emptypb.Empty, error) {
	err := s.accountUserUseCase.DeleteAccountUser(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
func (s *AccountUserService) GetAccountUser(ctx context.Context, req *v1.GetAccountUserReq) (*v1.AccountUser, error) {
	return s.accountUserUseCase.GetAccountUser(ctx, req.Id)
}

func (s *AccountUserService) ListAccountUser(ctx context.Context, req *v1.ListAccountUserReq) (*v1.ListAccountUserResp, error) {
	return s.accountUserUseCase.ListAccountUser(ctx, req)
}
func (s *AccountUserService) PageListAccountUser(ctx context.Context, req *v1.PageListAccountUserReq) (*v1.PageListAccountUserResp, error) {
	return s.accountUserUseCase.PageListAccountUser(ctx, req)
}
