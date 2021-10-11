package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	userv1 "github.com/lalifeier/vgo/api/account/service/v1"
	"github.com/lalifeier/vgo/pkg/auth/jwtauth"
)

type User struct {
	Id       int64
	Username string
	Password string
}

type UserRepo interface {
	Register(ctx context.Context, u *User) (*User, error)
	Login(ctx context.Context, u *User) (string, error)
}

type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
	uc   userv1.UserClient
}

func NewUserUsecase(repo UserRepo, logger log.Logger, uc userv1.UserClient) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(logger), uc: uc}
}

func (uc *UserUsecase) Register(ctx context.Context, u *User) (*User, error) {
	return uc.repo.Register(ctx, u)
}

func (uc *UserUsecase) Login(ctx context.Context, u *User) (string, error) {
	token, err := uc.repo.Login(ctx, u)
	if err != nil {
		return token, err
	}
	return jwtauth.New(jwtauth.WithExpired(2*60*60)).GenerateToken(u.Id, u.Username)
}

func (uc *UserUsecase) Logout(ctx context.Context, u *User) error {
	return nil
}
