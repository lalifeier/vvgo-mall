package biz

import (
	"context"
	"errors"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/lalifeier/vvgo-mall/pkg/util/crypto/bcrypt"
	"github.com/lalifeier/vvgo-mall/pkg/validate"

	v1 "github.com/lalifeier/vvgo-mall/gen/api/go/account/service/v1"
)

type AuthUseCase struct {
	accountUserRepo AccountUserRepo
	log             *log.Helper
}

func NewAuthUseCase(accountUserRepo AccountUserRepo, logger log.Logger) *AuthUseCase {
	return &AuthUseCase{accountUserRepo: accountUserRepo, log: log.NewHelper(logger)}
}

func (uc *AuthUseCase) GetUserByUsername(ctx context.Context, username string) (*AccountUser, error) {
	user, err := uc.accountUserRepo.FindByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (uc *AuthUseCase) Register(ctx context.Context, accountuser *AccountUser) (int64, error) {
	_, err := uc.accountUserRepo.FindByUsername(ctx, accountuser.Username)
	if !errors.Is(err, ErrUserNotFound) {
		return 0, v1.ErrorRegisterFailed("username already exists")
	}

	hashedPassword, err := bcrypt.GeneratePassword(accountuser.Password)
	if err != nil {
		return 0, err
	}
	accountuser.Password = hashedPassword

	user, err := uc.accountUserRepo.Create(ctx, accountuser)
	if err != nil {
		return 0, err
	}
	return user.Id, err
}

func (uc *AuthUseCase) Login(ctx context.Context, accountuser *AccountUser) (*AccountUser, error) {
	var (
		user *AccountUser
		err  error
	)

	if validate.IsEmail(accountuser.Email) {
		user, err = uc.accountUserRepo.FindByEmail(ctx, accountuser.Email)
	} else if validate.IsPhone(accountuser.Phone) {
		user, err = uc.accountUserRepo.FindByPhone(ctx, accountuser.Phone)
	} else {
		user, err = uc.accountUserRepo.FindByUsername(ctx, accountuser.Username)
	}

	if errors.Is(err, ErrUserNotFound) {
		return nil, v1.ErrorLoginFailed("username or password error")
	}

	ok := bcrypt.ValidatePassword(accountuser.Password, user.Password)

	if !ok {
		return nil, v1.ErrorLoginFailed("username or password error")
	}

	return user, nil
}
