package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	pb "github.com/lalifeier/vvgo-mall/api/account/service/v1"
	"github.com/lalifeier/vvgo-mall/app/account/service/internal/biz"
	"github.com/lalifeier/vvgo-mall/app/account/service/internal/data/ent"

	"github.com/lalifeier/vvgo-mall/pkg/utils/pagination"
)

var _ biz.AccountUserRepo = (*accountUserRepo)(nil)

type entAccountUser ent.AccountUser

func (e entAccountUser) BizStruct() *biz.AccountUser {
	return &biz.AccountUser{
		Id: e.ID,
	}
}

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

func (rp *accountUserRepo) CreateAccountUser(ctx context.Context, b *biz.AccountUser) (*biz.AccountUser, error) {
	p, err := rp.data.db.AccountUser.
		Create().
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return entAccountUser(*p).BizStruct(), nil
}

func (rp *accountUserRepo) UpdateAccountUser(ctx context.Context, b *biz.AccountUser) (*biz.AccountUser, error) {
	_, err := rp.data.db.AccountUser.Get(ctx, b.Id)
	if err != nil {
		return nil, err
	}
	db := rp.data.db.AccountUser.UpdateOneID(b.Id)

	p, err := db.Save(ctx)
	if err != nil {
		return nil, err
	}
	return entAccountUser(*p).BizStruct(), nil
}

func (rp *accountUserRepo) DeleteAccountUser(ctx context.Context, id int64) error {
	return rp.data.db.AccountUser.DeleteOneID(id).Exec(ctx)
}

func (rp *accountUserRepo) GetAccountUser(ctx context.Context, id int64) (*biz.AccountUser, error) {
	p, err := rp.data.db.AccountUser.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return entAccountUser(*p).BizStruct(), nil
}

func (rp *accountUserRepo) ListAccountUser(ctx context.Context, req *pb.ListAccountUserReq) ([]*biz.AccountUser, error) {
	query := rp.data.db.AccountUser.Query()

	ps, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	rv := make([]*biz.AccountUser, 0, len(ps))
	for _, p := range ps {
		rv = append(rv, entAccountUser(*p).BizStruct())
	}
	return rv, nil
}

func (rp *accountUserRepo) PageListAccountUser(ctx context.Context, req *pb.PageListAccountUserReq) ([]*biz.AccountUser, int64, error) {
	query := rp.data.db.AccountUser.Query()

	query.Order(ent.Desc("id"))
	total, err := query.Count(ctx)
	if err != nil {
		return nil, 0, err
	}

	ps, err := query.
		Offset(int(pagination.GetPageOffset(req.PageNum, req.PageSize))).
		Limit(int(req.PageSize)).
		All(ctx)
	if err != nil {
		return nil, 0, err
	}

	rv := make([]*biz.AccountUser, 0, len(ps))
	for _, p := range ps {
		rv = append(rv, entAccountUser(*p).BizStruct())
	}
	return rv, int64(total), nil
}
