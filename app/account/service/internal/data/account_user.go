package data

import (
	"context"
	"errors"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/lalifeier/vgo/app/account/service/internal/biz"
	"github.com/lalifeier/vgo/app/account/service/internal/data/ent"
	"github.com/lalifeier/vgo/app/account/service/internal/data/ent/accountuser"
	"github.com/lalifeier/vgo/pkg/util/pagination"
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
		log:  log.NewHelper(log.With(logger, "module", "account-service/data")),
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

func (rp *accountUserRepo) Register(ctx context.Context, au *biz.AccountUser) (*biz.AccountUser, error) {
	po, err := rp.data.db.AccountUser.Query().Where(accountuser.Username(au.Username)).First(ctx)
	if err != nil {
		return nil, err
	}
	if po != nil {
		return nil, errors.New("用户已存在")
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

func (rp *accountUserRepo) Create(ctx context.Context, au *biz.AccountUser) (*biz.AccountUser, error) {
	auCreate := rp.data.db.AccountUser.Create()
	if au.Password != "" {
		auCreate.SetPassword(au.Password)
	}
	if au.Phone != "" {
		auCreate.SetPhone(au.Phone)
	}
	if au.Email != "" {
		auCreate.SetEmail(au.Email)
	}
	po, err := auCreate.SetUsername(au.Username).Save(ctx)
	if err != nil {
		return nil, err
	}
	return entAccountUser(*po).BizAccountUser(), err
}

func (rp *accountUserRepo) Update(ctx context.Context, id int64, au *biz.AccountUser) (*biz.AccountUser, error) {
	po, err := rp.data.db.AccountUser.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	auUpdate := rp.data.db.AccountUser.Update()
	if au.Username != "" && au.Username != po.Username {
		auUpdate.SetUsername(au.Username)
	}
	if au.Password != "" && au.Password != po.Password {
		auUpdate.SetPassword(au.Password)
	}
	if au.Phone != "" && au.Phone != po.Phone {
		auUpdate.SetPhone(au.Phone)
	}
	if au.Email != "" && au.Email != po.Email {
		auUpdate.SetEmail(au.Email)
	}

	_, err = auUpdate.Save(ctx)
	return entAccountUser(*po).BizAccountUser(), err
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

func (rp *accountUserRepo) List(ctx context.Context, pageNum, pageSize int64) ([]*biz.AccountUser, error) {
	query := rp.data.db.AccountUser.Query()
	pos, err := query.Offset(int(pagination.GetPageOffset(pageNum, pageSize))).Limit(int(pageSize)).Order(ent.Desc("id")).All(ctx)
	if err != nil {
		return nil, err
	}

	rv := make([]*biz.AccountUser, 0)
	for _, po := range pos {
		rv = append(rv, entAccountUser(*po).BizAccountUser())
	}
	return rv, nil
}

// func (rp *accountUserRepo) CreateBulk(ctx context.Context, builders *ent.AccountUserCreate) ([]*ent.AccountUser, error) {
// 	return rp.data.db.AccountUser.CreateBulk(builders).Save(ctx)
// }

// func (rp *accountUserRepo) Update(ctx context.Context, ps predicate.AccountUser) error {
// 	return rp.data.db.AccountUser.Update().Where(ps).Exec(ctx)
// }

// func (rp *accountUserRepo) UpdateOne(ctx context.Context, au *ent.AccountUser) error {
// 	return rp.data.db.AccountUser.UpdateOne(au).Exec(ctx)
// }

// func (rp *accountUserRepo) UpdateByID(ctx context.Context, id int64) error {
// 	return rp.data.db.AccountUser.UpdateOneID(id).Exec(ctx)
// }

// func (rp *accountUserRepo) DeleteMany(ctx context.Context, ps predicate.AccountUser) error {
// 	_, err := rp.data.db.AccountUser.Delete().Where(ps).Exec(ctx)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (rp *accountUserRepo) DeleteOne(ctx context.Context, au *ent.AccountUser) error {
// 	return rp.data.db.AccountUser.DeleteOne(au).Exec(ctx)
// }

// func (rp *accountUserRepo) DeleteByID(ctx context.Context, id int64) error {
// 	return rp.data.db.AccountUser.DeleteOneID(id).Exec(ctx)
// }

// func (rp *accountUserRepo) GetById(ctx context.Context, id int64) (*ent.AccountUser, error) {
// 	return rp.data.db.AccountUser.Get(ctx, id)
// }

// func (rp *accountUserRepo) Query(ctx context.Context, ps predicate.AccountUser) ([]*ent.AccountUser, error) {
// 	return rp.data.db.AccountUser.Query().Where(ps).All(ctx)
// }

// func (rp *accountUserRepo) QueryAll(ctx context.Context, ps predicate.AccountUser) ([]*ent.AccountUser, error) {
// 	return rp.data.db.AccountUser.Query().All(ctx)
// }

// func (rp *accountUserRepo) List(ctx context.Context, ps predicate.AccountUser, page int, limit int) ([]*ent.AccountUser, error) {
// 	offset := (page - 1) * limit
// 	return rp.data.db.AccountUser.Query().Offset(offset).Limit(limit).Where(ps).All(ctx)
// }
