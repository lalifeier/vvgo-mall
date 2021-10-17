package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/lalifeier/vgo/app/account/service/internal/data/ent"
)

type AccountUser struct {
	Id       int64
	Email    string
	Phone    string
	Username string
	Password string
}

type AccountUserRepo interface {
	Create(ctx context.Context, au *ent.AccountUser) (*ent.AccountUser, error)
	// Update(ctx context.Context, u *AccountUser) error
	// Delete(ctx context.Context, id int64) error
	// Get(ctx context.Context, id int64) (*AccountUser, error)
	// List(ctx context.Context, u *AccountUser) (*[]AccountUser, error)
}

type AccountUserUsecase struct {
	repo AccountUserRepo
	log  *log.Helper
}

func NewAccountUserUsecase(repo AccountUserRepo, logger log.Logger) *AccountUserUsecase {
	return &AccountUserUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *AccountUserUsecase) Create(ctx context.Context, u *AccountUser) (*AccountUser, error) {
	// out, err := uc.repo.Create(ctx, u)
	// if err != nil {
	// 	return nil, err
	// }
	// return out, nil
	return nil, nil
}

func (uc *AccountUserUsecase) Update(ctx context.Context, u *AccountUser) error {
	// return uc.repo.Update(ctx, u)
	return nil
}

func (uc *AccountUserUsecase) Get(ctx context.Context, id int64) (*AccountUser, error) {
	// return uc.repo.Get(ctx, id)
	return nil, nil
}
