package biz

import (
	"context"

	"github.com/golang-jwt/jwt"
	"github.com/lalifeier/vvgo-mall/app/shop/admin/internal/conf"
)

type AccountUser struct {
	Id       int64
	Username string
	Password string
	Email    string
	Phone    string
	Token    string
}

type AccountRepo interface {
	Register(ctx context.Context, u *AccountUser) (int64, error)
	Login(ctx context.Context, u *AccountUser) (*AccountUser, error)
}

type AuthUseCase struct {
	key         string
	userRepo    UserRepo
	accountRepo AccountRepo
}

func NewAuthUseCase(conf *conf.Auth, userRepo UserRepo, accountRepo AccountRepo) *AuthUseCase {
	return &AuthUseCase{
		key:         conf.ApiKey,
		userRepo:    userRepo,
		accountRepo: accountRepo,
	}
}

func (uc *AuthUseCase) Register(ctx context.Context, u *AccountUser) (int64, error) {
	id, err := uc.accountRepo.Register(ctx, u)
	if err != nil {
		return 0, err
	}

	return id, nil
}
func (uc *AuthUseCase) Login(ctx context.Context, u *AccountUser) (*AccountUser, error) {
	user, err := uc.accountRepo.Login(ctx, u)
	if err != nil {
		return nil, err
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.Id,
	})
	signedString, err := claims.SignedString([]byte(uc.key))
	if err != nil {
		return nil, err
	}
	user.Token = signedString
	return user, nil
}
