package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/lalifeier/vvgo-mall/app/shop/admin/internal/biz"
	acV1 "github.com/lalifeier/vvgo-mall/gen/api/go/account/service/v1"
)

var _ biz.AccountRepo = (*accountRepo)(nil)

type accountRepo struct {
	data *Data
	log  *log.Helper
}

func NewAccountRepo(data *Data, logger log.Logger) biz.AccountRepo {
	return &accountRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "shop-admin/data")),
	}
}

func (rp *accountRepo) Register(ctx context.Context, u *biz.AccountUser) (int64, error) {
	resp, err := rp.data.ac.Register(ctx, &acV1.RegisterReq{
		Username: u.Username,
		Password: u.Password,
		Email:    u.Email,
		Phone:    u.Phone,
	})
	if err != nil {
		return 0, err
	}
	return resp.Id, nil
}

func (rp *accountRepo) Login(ctx context.Context, u *biz.AccountUser) (*biz.AccountUser, error) {
	resp, err := rp.data.ac.Login(ctx, &acV1.LoginReq{
		Username: u.Username,
		Password: u.Password,
		Email:    u.Email,
		Phone:    u.Phone,
	})
	if err != nil {
		return nil, err
	}
	return &biz.AccountUser{
		Id:       resp.Id,
		Username: resp.Username,
		Email:    resp.Email,
		Phone:    resp.Phone,
	}, nil
}
