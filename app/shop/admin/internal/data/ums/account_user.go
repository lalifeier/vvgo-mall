package ums

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"

	umsV1 "github.com/lalifeier/vvgo/api/ums/service/v1"
	"github.com/lalifeier/vvgo/app/shop/admin/internal/biz/ums"
	"github.com/lalifeier/vvgo/app/shop/admin/internal/data"
)

var _ ums.AccountUserRepo = (*accountUserRepo)(nil)

type accountUserRepo struct {
	data *data.Data
	log  *log.Helper
}

func NewAccountUserRepo(data *data.Data, logger log.Logger) ums.AccountUserRepo {
	return &accountUserRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "front-admin/data/account-user")),
	}
}

func (rp *accountUserRepo) CreateAccountUser(ctx context.Context, au *ums.AccountUser) (int64, error) {
	resp, err := rp.data.UmsClient.CreateAccountUser(ctx, &umsV1.CreateAccountUserReq{
		Username: au.Username,
		Password: au.Password,
		Phone:    au.Phone,
		Email:    au.Email,
	})
	if err != nil {
		return 0, err
	}
	return resp.Id, nil
}

func (rp *accountUserRepo) UpdateAccountUser(ctx context.Context, au *ums.AccountUser) error {
	_, err := rp.data.UmsClient.UpdateAccountUser(ctx, &umsV1.UpdateAccountUserReq{
		Id:       au.Id,
		Username: au.Username,
		Password: au.Password,
		Phone:    au.Phone,
		Email:    au.Email,
	})
	return err
}

func (rp *accountUserRepo) DeleteAccountUser(ctx context.Context, id int64) error {
	_, err := rp.data.UmsClient.DeleteAccountUser(ctx, &umsV1.DeleteAccountUserReq{Id: id})
	return err
}

func (rp *accountUserRepo) GetAccountUser(ctx context.Context, id int64) (*ums.AccountUser, error) {
	resp, err := rp.data.UmsClient.GetAccountUser(ctx, &umsV1.GetAccountUserReq{Id: id})
	if err != nil {
		return nil, err
	}

	accountUser := ums.AccountUser{}
	err = copier.Copy(&accountUser, resp)
	if err != nil {
		return nil, err
	}

	return &accountUser, nil
}

func (rp *accountUserRepo) ListAccountUser(ctx context.Context, au *ums.AccountUserListReq) (*ums.AccountUserListResp, error) {
	resp, err := rp.data.UmsClient.ListAccountUser(ctx, &umsV1.ListAccountUserReq{
		PageNum:  au.PageNum,
		PageSize: au.PageSize,
	})

	if err != nil {
		return nil, err
	}

	accountUsers := make([]*ums.AccountUser, 0)
	err = copier.Copy(&accountUsers, &resp.Data)
	if err != nil {
		return nil, err
	}

	return &ums.AccountUserListResp{
		TotalPages:  resp.TotalPages,
		CurrentPage: resp.CurrentPage,
		PageSize:    resp.PageSize,
		Data:        accountUsers,
	}, err
}
