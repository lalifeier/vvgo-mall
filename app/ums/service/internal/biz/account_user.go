package biz

import (
	"context"
	"errors"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/lalifeier/vvgo-mall/pkg/util/crypto/bcrypt"
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
	Register(ctx context.Context, au *AccountUser) (int64, error)
	GetAccountUserByUsernameOrMobileOrEmail(ctx context.Context, account string) ([]*AccountUser, error)

	Create(ctx context.Context, au *AccountUser) (int64, error)
	Update(ctx context.Context, au *AccountUser) error
	Delete(ctx context.Context, id int64) error
	Get(ctx context.Context, id int64) (*AccountUser, error)
	List(ctx context.Context, req *AccountUserListReq) (*AccountUserListResp, error)
}

type AccountUserUseCase struct {
	repo AccountUserRepo
	log  *log.Helper
}

func NewAccountUserUseCase(repo AccountUserRepo, logger log.Logger) *AccountUserUseCase {
	return &AccountUserUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (au *AccountUser) CheckPassword(password string) bool {
	return bcrypt.ValidatePassword(password, au.Password)
}

func (au *AccountUser) SetPassword(password string) {
	p, err := bcrypt.GeneratePassword(password)
	if err == nil {
		au.Password = p
	}
}

func (uc *AccountUserUseCase) Login(ctx context.Context, account, password string) (*AccountUser, error) {
	accounts, err := uc.repo.GetAccountUserByUsernameOrMobileOrEmail(ctx, account)
	if err != nil {
		return nil, err
	}
	for _, account := range accounts {
		if account.CheckPassword(password) {
			_ = uc.repo.Update(ctx, &AccountUser{})
			return account, nil
		}
	}
	return nil, errors.New("账号或密码错误")
}

func (uc *AccountUserUseCase) Register(ctx context.Context, au *AccountUser) (*AccountUser, error) {
	if au.Username == "" {
		return nil, errors.New("用户名不能为空")
	}
	if au.Password == "" {
		return nil, errors.New("密码不能为空")
	}

	au.SetPassword(au.Password)
	return uc.Register(ctx, au)
}

func (uc *AccountUserUseCase) UpdateAccountUser(ctx context.Context, id int64, au *AccountUser) (*AccountUser, error) {
	if au.Password != "" {
		au.SetPassword(au.Password)
	}
	return au, uc.Update(ctx, au)
}

func (uc *AccountUserUseCase) Create(ctx context.Context, au *AccountUser) (int64, error) {
	return uc.repo.Create(ctx, au)
}

func (uc *AccountUserUseCase) Update(ctx context.Context, au *AccountUser) error {
	return uc.repo.Update(ctx, au)
}

func (uc *AccountUserUseCase) Delete(ctx context.Context, id int64) error {
	return uc.repo.Delete(ctx, id)
}

func (uc *AccountUserUseCase) Get(ctx context.Context, id int64) (*AccountUser, error) {
	return uc.repo.Get(ctx, id)
}

func (uc *AccountUserUseCase) List(ctx context.Context, req *AccountUserListReq) (*AccountUserListResp, error) {
	return uc.repo.List(ctx, &AccountUserListReq{
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
	})
}
