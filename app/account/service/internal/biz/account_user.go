package biz

import (
	"context"
	"errors"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/lalifeier/vgo/pkg/crypto/bcrypt"
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

type AccountUserRepo interface {
	Register(ctx context.Context, au *AccountUser) (*AccountUser, error)
	GetAccountUserByUsernameOrMobileOrEmail(ctx context.Context, account string) ([]*AccountUser, error)

	Create(ctx context.Context, au *AccountUser) (*AccountUser, error)
	Update(ctx context.Context, id int64, au *AccountUser) (*AccountUser, error)
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context, pageNum, pageSize int64) ([]*AccountUser, error)
	Get(ctx context.Context, id int64) (*AccountUser, error)
}

type AccountUserUsecase struct {
	repo AccountUserRepo
	log  *log.Helper
}

func NewAccountUserUsecase(repo AccountUserRepo, logger log.Logger) *AccountUserUsecase {
	return &AccountUserUsecase{repo: repo, log: log.NewHelper(logger)}
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

func (uc *AccountUserUsecase) Login(ctx context.Context, account, password string) (*AccountUser, error) {
	accounts, err := uc.repo.GetAccountUserByUsernameOrMobileOrEmail(ctx, account)
	if err != nil {
		return nil, err
	}
	for _, account := range accounts {
		if account.CheckPassword(password) {
			_, _ = uc.repo.Update(ctx, account.Id, &AccountUser{})
			return account, nil
		}
	}
	return nil, errors.New("账号或密码错误")
}

func (uc *AccountUserUsecase) Register(ctx context.Context, au *AccountUser) (*AccountUser, error) {
	if au.Username == "" {
		return nil, errors.New("用户名不能为空")
	}
	if au.Password == "" {
		return nil, errors.New("密码不能为空")
	}

	au.SetPassword(au.Password)
	return uc.Register(ctx, au)
}

func (uc *AccountUserUsecase) UpdateAccountUser(ctx context.Context, id int64, au *AccountUser) (*AccountUser, error) {
	if au.Password != "" {
		au.SetPassword(au.Password)
	}
	return uc.Update(ctx, id, au)
}

func (uc *AccountUserUsecase) Create(ctx context.Context, au *AccountUser) (*AccountUser, error) {
	return uc.repo.Create(ctx, au)
}

func (uc *AccountUserUsecase) Update(ctx context.Context, id int64, au *AccountUser) (*AccountUser, error) {
	return uc.repo.Update(ctx, id, au)
}

func (uc *AccountUserUsecase) Delete(ctx context.Context, id int64) error {
	return uc.repo.Delete(ctx, id)
}

func (uc *AccountUserUsecase) Get(ctx context.Context, id int64) (*AccountUser, error) {
	return uc.repo.Get(ctx, id)
}

func (uc *AccountUserUsecase) List(ctx context.Context, pageNum, pageSize int64) ([]*AccountUser, error) {
	return uc.repo.List(ctx, pageNum, pageSize)
}

// func (uc *AccountUserUsecase) LoginByUsername(ctx context.Context, Username, password string) (*AccountUser, error) {

// }

// func (uc *AccountUserUsecase) LoginByPhone(ctx context.Context, phone, code string) (*AccountUser, error) {
// }

// func (uc *AccountUserUsecase) LoginByUEmail(ctx context.Context, email, code string) (*AccountUser, error) {
// }
