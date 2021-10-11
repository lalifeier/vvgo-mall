package data

import (
	"context"
	"errors"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/lalifeier/vgo/app/account/interface/internal/biz"

	userv1 "github.com/lalifeier/vgo/api/account/service/v1"
)

var _ biz.UserRepo = (*userRepo)(nil)

type userRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "user-interface/data")),
	}
}

func (rp *userRepo) Register(ctx context.Context, u *biz.User) (*biz.User, error) {
	reply, err := rp.data.uc.CreateUser(ctx, &userv1.CreateUserRequest{
		Username: u.Username,
		Password: u.Password,
	})
	if err != nil {
		return nil, err
	}

	return &biz.User{
		Id:       reply.Id,
		Username: reply.Username,
	}, err
}

func (rp *userRepo) Login(ctx context.Context, u *biz.User) (string, error) {
	reply, err := rp.data.uc.VerifyPassword(ctx, &userv1.VerifyPasswordRequest{
		Username: u.Username,
		Password: u.Password,
	})
	if err != nil {
		return "", err
	}
	if reply.Ok {
		return "some_token", nil
	}
	return "", errors.New("login failed")
}

func (rp *userRepo) Logout(ctx context.Context, u *biz.User) error {
	return nil
}

func (r *userRepo) GetUser(ctx context.Context, id int64) (*biz.User, error) {
	reply, err := r.data.uc.GetUser(ctx, &userv1.GetUserRequest{
		Id: id,
	})
	if err != nil {
		return nil, err
	}
	return &biz.User{
		Id: reply.Id,
	}, err
}

func (r *userRepo) ListUser(ctx context.Context, pageNum, pageSize int64) ([]*biz.User, error) {
	reply, err := r.data.uc.ListUser(ctx, &userv1.ListUserRequest{})
	if err != nil {
		return nil, err
	}
	rv := make([]*biz.User, 0)
	for _, x := range reply.Results {
		rv = append(rv, &biz.User{
			Id: x.Id,
		})
	}
	return rv, err
}
