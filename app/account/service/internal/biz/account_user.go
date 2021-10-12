package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type AccountUser struct {
	Id       int64
	Email    string
	Phone    int64
	Username string
	Password string
}

type AccountUserRepo interface {
	CreateAccountUser(ctx context.Context, u *AccountUser) (*AccountUser, error)
	UpdateAccountUser(ctx context.Context, u *AccountUser) (*AccountUser, error)
	DeleteAccountUser(ctx context.Context, id int64) (*AccountUser, error)
	GetAccountUser(ctx context.Context, id int64) (*AccountUser, error)
	ListAccountUser(ctx context.Context, id int64) (*[]AccountUser, error)
}

type AccountUserUsecase struct {
	repo AccountUserRepo
	log  *log.Helper
}

func NewAccountUserUsecase(repo AccountUserRepo, logger log.Logger) *AccountUserUsecase {
	return &AccountUserUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *AccountUserUsecase) Create(ctx context.Context, u *AccountUser) (*AccountUser, error) {
	out, err := uc.repo.CreateAccountUser(ctx, u)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (uc *AccountUserUsecase) Update(ctx context.Context, u *AccountUser) (*AccountUser, error) {
	out, err := uc.repo.UpdateAccountUser(ctx, u)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (uc *AccountUserUsecase) Get(ctx context.Context, id int64) (*AccountUser, error) {
	return uc.repo.GetAccountUser(ctx, id)
}
