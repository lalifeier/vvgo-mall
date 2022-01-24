package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	pb "github.com/lalifeier/vvgo-mall/api/account/service/v1"
)

type AccountUser struct {
	Id int64
}

type AccountUserRepo interface {
	CreateAccountUser(ctx context.Context, accountuser *AccountUser) (*AccountUser, error)
	UpdateAccountUser(ctx context.Context, accountuser *AccountUser) (*AccountUser, error)
	DeleteAccountUser(ctx context.Context, id int64) error
	GetAccountUser(ctx context.Context, id int64) (*AccountUser, error)
	ListAccountUser(ctx context.Context, req *pb.ListAccountUserReq) ([]*AccountUser, error)
	PageListAccountUser(ctx context.Context, req *pb.PageListAccountUserReq) ([]*AccountUser, int64, error)
}

type AccountUserUsecase struct {
	repo AccountUserRepo
	log  *log.Helper
}

func NewAccountUserUsecase(repo AccountUserRepo, logger log.Logger) *AccountUserUsecase {
	return &AccountUserUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *AccountUserUsecase) CreateAccountUser(ctx context.Context, accountuser *AccountUser) (*AccountUser, error) {
	return uc.repo.CreateAccountUser(ctx, accountuser)
}

func (uc *AccountUserUsecase) UpdateAccountUser(ctx context.Context, accountuser *AccountUser) (*AccountUser, error) {
	return uc.repo.UpdateAccountUser(ctx, accountuser)
}

func (uc *AccountUserUsecase) DeleteAccountUser(ctx context.Context, id int64) error {
	return uc.repo.DeleteAccountUser(ctx, id)
}

func (uc *AccountUserUsecase) GetAccountUser(ctx context.Context, id int64) (*AccountUser, error) {
	return uc.repo.GetAccountUser(ctx, id)
}

func (uc *AccountUserUsecase) ListAccountUser(ctx context.Context, req *pb.ListAccountUserReq) ([]*AccountUser, error) {
	return uc.repo.ListAccountUser(ctx, req)
}

func (uc *AccountUserUsecase) PageListAccountUser(ctx context.Context, req *pb.PageListAccountUserReq) ([]*AccountUser, int64, error) {
	return uc.repo.PageListAccountUser(ctx, req)
}
