package biz

import (
	"context"
	"errors"

	"github.com/go-kratos/kratos/v2/log"
	v1 "github.com/lalifeier/vvgo-mall/gen/api/go/account/service/v1"
)

var (
	ErrPasswordInvalid = errors.New("password invalid")
	ErrUsernameInvalid = errors.New("username invalid")
	ErrUserNotFound    = errors.New("user not found")
)

type AccountUserRepo interface {
	Create(ctx context.Context, req *v1.CreateAccountUserReq) (*v1.AccountUser, error)
	Update(ctx context.Context, req *v1.UpdateAccountUserReq) (*v1.AccountUser, error)
	Delete(ctx context.Context, id uint32) error
	Get(ctx context.Context, id uint32) (*v1.AccountUser, error)
	List(ctx context.Context, req *v1.ListAccountUserReq) (*v1.ListAccountUserResp, error)
	PageList(ctx context.Context, req *v1.PageListAccountUserReq) (*v1.PageListAccountUserResp, error)

	FindByUsername(ctx context.Context, username string) (*v1.AccountUser, error)
	FindByEmail(ctx context.Context, email string) (*v1.AccountUser, error)
	FindByPhone(ctx context.Context, phone string) (*v1.AccountUser, error)
}

type AccountUserUseCase struct {
	repo AccountUserRepo
	log  *log.Helper
}

func NewAccountUserUseCase(repo AccountUserRepo, logger log.Logger) *AccountUserUseCase {
	return &AccountUserUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *AccountUserUseCase) CreateAccountUser(ctx context.Context, req *v1.CreateAccountUserReq) (*v1.AccountUser, error) {
	return uc.repo.Create(ctx, req)
}

func (uc *AccountUserUseCase) UpdateAccountUser(ctx context.Context, req *v1.UpdateAccountUserReq) (*v1.AccountUser, error) {
	return uc.repo.Update(ctx, req)
}

func (uc *AccountUserUseCase) DeleteAccountUser(ctx context.Context, id uint32) error {
	return uc.repo.Delete(ctx, id)
}

func (uc *AccountUserUseCase) GetAccountUser(ctx context.Context, id uint32) (*v1.AccountUser, error) {
	return uc.repo.Get(ctx, id)
}

func (uc *AccountUserUseCase) ListAccountUser(ctx context.Context, req *v1.ListAccountUserReq) (*v1.ListAccountUserResp, error) {
	return uc.repo.List(ctx, req)
}

func (uc *AccountUserUseCase) PageListAccountUser(ctx context.Context, req *v1.PageListAccountUserReq) (*v1.PageListAccountUserResp, error) {
	return uc.repo.PageList(ctx, req)
}
