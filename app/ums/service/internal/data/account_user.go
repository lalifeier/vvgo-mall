package data

import (
	"context"
	"errors"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/lalifeier/vvgo/app/ums/service/internal/biz"
	"github.com/lalifeier/vvgo/app/ums/service/internal/data/ent"
	"github.com/lalifeier/vvgo/app/ums/service/internal/data/ent/accountuser"
	"github.com/lalifeier/vvgo/pkg/utils/pagination"
	"github.com/spf13/cast"
)

var _ biz.AccountUserRepo = (*accountUserRepo)(nil)

type accountUserRepo struct {
	data *Data
	log  *log.Helper
}

func NewAccountUserRepo(data *Data, logger log.Logger) biz.AccountUserRepo {
	return &accountUserRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "ums-service/data")),
	}
}

type entAccountUser ent.AccountUser

func (e entAccountUser) BizAccountUser() *biz.AccountUser {
	return &biz.AccountUser{
		Id:            e.ID,
		Username:      e.Username,
		Password:      e.Password,
		Phone:         e.Phone,
		Email:         e.Email,
		CreateAt:      cast.ToTime(e.CreateAt),
		CreateIPAt:    e.CreateIPAt,
		LastLoginAt:   cast.ToTime(e.LastLoginIPAt),
		LastLoginIPAt: e.LastLoginIPAt,
		LoginTimes:    e.LoginTimes,
		Status:        e.Status,
	}
}

func (rp *accountUserRepo) Register(ctx context.Context, au *biz.AccountUser) (int64, error) {
	po, err := rp.data.db.AccountUser.Query().Where(accountuser.Username(au.Username)).First(ctx)
	if err != nil {
		return 0, err
	}
	if po != nil {
		return 0, errors.New("用户已存在")
	}
	return rp.Create(ctx, au)
}

func (rp *accountUserRepo) GetAccountUserByUsernameOrMobileOrEmail(ctx context.Context, account string) ([]*biz.AccountUser, error) {
	ps, err := rp.data.db.AccountUser.Query().Where(accountuser.Or(accountuser.Username(account), accountuser.Phone(account), accountuser.Email(account))).All(ctx)
	if err != nil {
		return nil, err
	}

	rv := make([]*biz.AccountUser, 0)
	for _, p := range ps {
		rv = append(rv, entAccountUser(*p).BizAccountUser())
	}
	return rv, nil
}

func (rp *accountUserRepo) GetAccountUserByUsername(ctx context.Context, username string) (*biz.AccountUser, error) {
	p, err := rp.data.db.AccountUser.Query().Where(accountuser.Username(username)).First(ctx)
	if err != nil {
		return nil, err
	}
	if p == nil {
		return nil, errors.New("用户不存在")
	}
	return entAccountUser(*p).BizAccountUser(), err
}

func (rp *accountUserRepo) GetAccountUserByPhone(ctx context.Context, phone string) (*biz.AccountUser, error) {
	p, err := rp.data.db.AccountUser.Query().Where(accountuser.Phone(phone)).First(ctx)
	if err != nil {
		return nil, err
	}
	if p == nil {
		return nil, errors.New("用户不存在")
	}
	return entAccountUser(*p).BizAccountUser(), err
}

func (rp *accountUserRepo) GetAccountUserByEmail(ctx context.Context, email string) (*biz.AccountUser, error) {
	p, err := rp.data.db.AccountUser.Query().Where(accountuser.Email(email)).First(ctx)
	if err != nil {
		return nil, err
	}
	if p == nil {
		return nil, errors.New("用户不存在")
	}
	return entAccountUser(*p).BizAccountUser(), err
}

func (rp *accountUserRepo) Create(ctx context.Context, au *biz.AccountUser) (int64, error) {

	create := rp.data.db.AccountUser.Create()
	if au.Username != "" {
		create.SetUsername(au.Username)
	}
	if au.Password != "" {
		create.SetPassword(au.Password)
	}
	if au.Phone != "" {
		create.SetPhone(au.Phone)
	}
	if au.Email != "" {
		create.SetEmail(au.Email)
	}
	po, err := create.Save(ctx)
	if err != nil {
		return 0, err
	}
	return po.ID, err
}

func (rp *accountUserRepo) Update(ctx context.Context, au *biz.AccountUser) error {
	po, err := rp.data.db.AccountUser.Get(ctx, au.Id)
	if err != nil {
		return err
	}

	update := rp.data.db.AccountUser.Update()
	if au.Username != "" && au.Username != po.Username {
		update.SetUsername(au.Username)
	}
	if au.Password != "" && au.Password != po.Password {
		update.SetPassword(au.Password)
	}
	if au.Phone != "" && au.Phone != po.Phone {
		update.SetPhone(au.Phone)
	}
	if au.Email != "" && au.Email != po.Email {
		update.SetEmail(au.Email)
	}

	_, err = update.Save(ctx)
	return err
}

func (rp *accountUserRepo) Delete(ctx context.Context, id int64) error {
	return rp.data.db.AccountUser.DeleteOneID(id).Exec(ctx)
}

func (rp *accountUserRepo) Get(ctx context.Context, id int64) (*biz.AccountUser, error) {
	po, err := rp.data.db.AccountUser.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return entAccountUser(*po).BizAccountUser(), nil
}

func (rp *accountUserRepo) List(ctx context.Context, req *biz.AccountUserListReq) (*biz.AccountUserListResp, error) {
	// count, err := rp.Count(ctx, req)
	// if err != nil {
	// 	return nil, err
	// }
	pos, err := rp.data.db.AccountUser.Query().
		Offset(int(pagination.GetPageOffset(req.PageNum, req.PageSize))).
		Limit(int(req.PageSize)).
		All(ctx)
	if err != nil {
		return nil, err
	}

	rv := make([]*biz.AccountUser, 0)
	for _, po := range pos {
		rv = append(rv, entAccountUser(*po).BizAccountUser())
	}
	return &biz.AccountUserListResp{
		TotalPages:  0,
		CurrentPage: req.PageNum,
		PageSize:    req.PageSize,
		Data:        rv,
	}, nil
}

func (rp *accountUserRepo) Count(ctx context.Context, req *biz.AccountUserListReq) (int64, error) {
	count, err := rp.data.db.AccountUser.Query().Count(ctx)
	if err != nil {
		return 0, err
	}
	return int64(count), nil
}
