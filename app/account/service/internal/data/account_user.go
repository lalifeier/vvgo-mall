package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/lalifeier/vvgo-mall/app/account/service/internal/biz"
	"github.com/lalifeier/vvgo-mall/app/account/service/internal/data/ent"
	"github.com/lalifeier/vvgo-mall/app/account/service/internal/data/ent/accountuser"
	v1 "github.com/lalifeier/vvgo-mall/gen/api/go/account/service/v1"

	"github.com/lalifeier/vvgo-mall/pkg/util/pagination"
)

var _ biz.AccountUserRepo = (*accountUserRepo)(nil)

type entAccountUser ent.AccountUser

func (e entAccountUser) convertEntToProto() *v1.AccountUser {
	return &v1.AccountUser{
		Id:       e.ID,
		Username: e.Username,
		Email:    e.Email,
		Phone:    e.Phone,
		Password: e.Password,
	}
}

type accountUserRepo struct {
	data *Data
	log  *log.Helper
}

func NewAccountUserRepo(data *Data, logger log.Logger) biz.AccountUserRepo {
	return &accountUserRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "account-service/repo/account-user")),
	}
}

func (rp *accountUserRepo) Create(ctx context.Context, b *biz.AccountUser) (*biz.AccountUser, error) {
	po, err := rp.data.db.AccountUser.
		Create().
		SetNillableUsername(&b.Username).
		SetNillableEmail(&b.Email).
		SetNillablePhone(&b.Phone).
		SetNillablePassword(&b.Password).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return entAccountUser(*po).BizStruct(), nil
}

func (rp *accountUserRepo) Update(ctx context.Context, b *biz.AccountUser) (*biz.AccountUser, error) {
	builder := rp.data.db.AccountUser.UpdateOneID(b.Id).
		SetNillableUsername(&b.Username)

	p, err := builder.Save(ctx)
	if err != nil {
		return nil, err
	}

	return entAccountUser(*p).BizStruct(), nil
}

func (rp *accountUserRepo) Delete(ctx context.Context, id int64) error {
	return rp.data.db.AccountUser.DeleteOneID(id).Exec(ctx)
}

func (rp *accountUserRepo) Get(ctx context.Context, id int64) (*biz.AccountUser, error) {
	po, err := rp.data.db.AccountUser.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return entAccountUser(*po).BizStruct(), nil
}

func (rp *accountUserRepo) List(ctx context.Context, req *v1.ListAccountUserReq) ([]*biz.AccountUser, error) {
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

func (rp *accountUserRepo) PageList(ctx context.Context, req *v1.PageListAccountUserReq) ([]*biz.AccountUser, int64, error) {
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

func (rp *accountUserRepo) FindByUsername(ctx context.Context, username string) (*biz.AccountUser, error) {
	p, err := rp.data.db.AccountUser.
		Query().
		Where(accountuser.UsernameEQ(username)).
		Only(ctx)
	if err != nil {
		return nil, biz.ErrUserNotFound
	}
	return entAccountUser(*p).BizStruct(), nil
}

func (rp *accountUserRepo) FindByEmail(ctx context.Context, email string) (*biz.AccountUser, error) {
	p, err := rp.data.db.AccountUser.
		Query().
		Where(accountuser.EmailEQ(email)).
		Only(ctx)
	if err != nil {
		return nil, biz.ErrUserNotFound
	}
	return entAccountUser(*p).BizStruct(), nil
}

func (rp *accountUserRepo) FindByPhone(ctx context.Context, phone string) (*biz.AccountUser, error) {
	p, err := rp.data.db.AccountUser.
		Query().
		Where(accountuser.PhoneEQ(phone)).
		Only(ctx)
	if err != nil {
		return nil, biz.ErrUserNotFound
	}
	return entAccountUser(*p).BizStruct(), nil
}
