package ums

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

type AccountUser struct {
	Id            int64
	Username      string
	Password      string `json:"-"`
	Phone         string
	Email         string
	CreateAt      time.Time
	CreateIPAt    string
	LastLoginAt   time.Time
	LastLoginIPAt string
	LoginTimes    int32
	Status        int8
}

type AccountUserListReq struct {
	PageNum  int64
	PageSize int64
}

type AccountUserListResp struct {
	TotalPages  int64
	CurrentPage int64
	PageSize    int64
	Data        []*AccountUser
}

type AccountUserRepo interface {
	CreateAccountUser(context.Context, *AccountUser) (int64, error)
	UpdateAccountUser(context.Context, *AccountUser) error
	DeleteAccountUser(ctx context.Context, id int64) error
	GetAccountUser(ctx context.Context, id int64) (*AccountUser, error)
	ListAccountUser(context.Context, *AccountUserListReq) (*AccountUserListResp, error)
}

type AccountUserUsecase struct {
	repo AccountUserRepo
	log  *log.Helper
}

func NewAccountUserUsecase(repo AccountUserRepo, logger log.Logger) *AccountUserUsecase {
	return &AccountUserUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *AccountUserUsecase) Register(ctx context.Context, u *AccountUser) (*AccountUser, error) {
	return nil, nil
}

func (uc *AccountUserUsecase) Login(ctx context.Context, u *AccountUser) (string, error) {
	return "", nil
}

func (uc *AccountUserUsecase) Logout(ctx context.Context, u *AccountUser) error {
	return nil
}

func (uc *AccountUserUsecase) UserInfo(ctx context.Context, id int64) (string, error) {
	return "", nil
}

func (uc *AccountUserUsecase) ListUser(ctx context.Context, u *AccountUser) (*[]AccountUser, error) {
	return nil, nil
}

func (uc *AccountUserUsecase) UpdateUser(ctx context.Context, u *AccountUser) (*AccountUser, error) {
	return nil, nil
}

func (uc *AccountUserUsecase) CreateAccountUser(ctx context.Context, au *AccountUser) (int64, error) {
	return uc.repo.CreateAccountUser(ctx, au)
}

func (uc *AccountUserUsecase) UpdateAccountUser(ctx context.Context, au *AccountUser) error {
	return uc.repo.UpdateAccountUser(ctx, au)
}

func (uc *AccountUserUsecase) DeleteAccountUser(ctx context.Context, id int64) error {
	return uc.repo.DeleteAccountUser(ctx, id)
}

func (uc *AccountUserUsecase) GetAccountUser(ctx context.Context, id int64) (*AccountUser, error) {
	return uc.repo.GetAccountUser(ctx, id)
}

func (uc *AccountUserUsecase) ListAccountUser(ctx context.Context, req *AccountUserListReq) (*AccountUserListResp, error) {
	return uc.repo.ListAccountUser(ctx, &AccountUserListReq{
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
	})
}
