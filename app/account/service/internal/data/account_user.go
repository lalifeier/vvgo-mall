package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/lalifeier/vgo/app/account/service/internal/biz"
	"github.com/lalifeier/vgo/app/account/service/internal/data/ent"
	"github.com/lalifeier/vgo/app/account/service/internal/data/ent/predicate"
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

func (rp *accountUserRepo) Create(ctx context.Context, au *ent.AccountUser) (*ent.AccountUser, error) {
	return rp.data.db.AccountUser.Create().SetUsername(au.Username).SetPassword(au.Password).Save(ctx)
}

func (rp *accountUserRepo) CreateBulk(ctx context.Context, builders *ent.AccountUserCreate) ([]*ent.AccountUser, error) {
	return rp.data.db.AccountUser.CreateBulk(builders).Save(ctx)
}

func (rp *accountUserRepo) Update(ctx context.Context, ps predicate.AccountUser) error {
	return rp.data.db.AccountUser.Update().Where(ps).Exec(ctx)
}

func (rp *accountUserRepo) UpdateOne(ctx context.Context, au *ent.AccountUser) error {
	return rp.data.db.AccountUser.UpdateOne(au).Exec(ctx)
}

func (rp *accountUserRepo) UpdateByID(ctx context.Context, id int64) error {
	return rp.data.db.AccountUser.UpdateOneID(id).Exec(ctx)
}

func (rp *accountUserRepo) DeleteMany(ctx context.Context, ps predicate.AccountUser) error {
	_, err := rp.data.db.AccountUser.Delete().Where(ps).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (rp *accountUserRepo) DeleteOne(ctx context.Context, au *ent.AccountUser) error {
	return rp.data.db.AccountUser.DeleteOne(au).Exec(ctx)
}

func (rp *accountUserRepo) DeleteByID(ctx context.Context, id int64) error {
	return rp.data.db.AccountUser.DeleteOneID(id).Exec(ctx)
}

func (rp *accountUserRepo) GetById(ctx context.Context, id int64) (*ent.AccountUser, error) {
	return rp.data.db.AccountUser.Get(ctx, id)
}

func (rp *accountUserRepo) Query(ctx context.Context, ps predicate.AccountUser) ([]*ent.AccountUser, error) {
	return rp.data.db.AccountUser.Query().Where(ps).All(ctx)
}

func (rp *accountUserRepo) QueryAll(ctx context.Context, ps predicate.AccountUser) ([]*ent.AccountUser, error) {
	return rp.data.db.AccountUser.Query().All(ctx)
}

func (rp *accountUserRepo) List(ctx context.Context, ps predicate.AccountUser, page int, limit int) ([]*ent.AccountUser, error) {
	offset := (page - 1) * limit
	return rp.data.db.AccountUser.Query().Offset(offset).Limit(limit).Where(ps).All(ctx)
}
