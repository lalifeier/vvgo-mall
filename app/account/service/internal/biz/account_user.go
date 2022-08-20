package biz

import (
	"context"
	"errors"

	"github.com/go-kratos/kratos/v2/log"
	pb "github.com/lalifeier/vvgo-mall/api/account/service/v1"
)

var (
	ErrPasswordInvalid = errors.New("password invalid")
	ErrUsernameInvalid = errors.New("username invalid")
	ErrUserNotFound    = errors.New("user not found")
)

type AccountUser struct {
	Id       int64
	Username string
	Password string
	Email    string
	Phone    string
}

type AccountUserRepo interface {
	Create(ctx context.Context, accountuser *AccountUser) (*AccountUser, error)
	Update(ctx context.Context, accountuser *AccountUser) (*AccountUser, error)
	Delete(ctx context.Context, id int64) error
	Get(ctx context.Context, id int64) (*AccountUser, error)
	List(ctx context.Context, req *pb.ListAccountUserReq) ([]*AccountUser, error)
	PageList(ctx context.Context, req *pb.PageListAccountUserReq) ([]*AccountUser, int64, error)

	FindByUsername(ctx context.Context, username string) (*AccountUser, error)
	FindByEmail(ctx context.Context, email string) (*AccountUser, error)
	FindByPhone(ctx context.Context, phone string) (*AccountUser, error)
}

type AccountUserUseCase struct {
	repo AccountUserRepo
	log  *log.Helper
}

func NewAccountUserUseCase(repo AccountUserRepo, logger log.Logger) *AccountUserUseCase {
	return &AccountUserUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *AccountUserUseCase) CreateAccountUser(ctx context.Context, accountuser *AccountUser) (*AccountUser, error) {
	return uc.repo.Create(ctx, accountuser)
}

func (uc *AccountUserUseCase) UpdateAccountUser(ctx context.Context, accountuser *AccountUser) (*AccountUser, error) {
	return uc.repo.Update(ctx, accountuser)
}

func (uc *AccountUserUseCase) DeleteAccountUser(ctx context.Context, id int64) error {
	return uc.repo.Delete(ctx, id)
}

func (uc *AccountUserUseCase) GetAccountUser(ctx context.Context, id int64) (*AccountUser, error) {
	return uc.repo.Get(ctx, id)
}

func (uc *AccountUserUseCase) ListAccountUser(ctx context.Context, req *pb.ListAccountUserReq) ([]*AccountUser, error) {
	return uc.repo.List(ctx, req)
}

func (uc *AccountUserUseCase) PageListAccountUser(ctx context.Context, req *pb.PageListAccountUserReq) ([]*AccountUser, int64, error) {
	return uc.repo.PageList(ctx, req)
}
