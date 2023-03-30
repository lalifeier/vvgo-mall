package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/lalifeier/vvgo-mall/app/account/service/internal/biz"

	v1 "github.com/lalifeier/vvgo-mall/gen/api/go/account/service/v1"
	"github.com/lalifeier/vvgo-mall/pkg/util/entgo"

	"github.com/lalifeier/vvgo-mall/app/account/service/internal/data/ent"
	"github.com/lalifeier/vvgo-mall/app/account/service/internal/data/ent/accountuser"

	paging "github.com/lalifeier/vvgo-mall/pkg/util/pagination"
)

var _ biz.AccountUserRepo = (*accountUserRepo)(nil)

func (rp *accountUserRepo) convertEntToProto(e *ent.AccountUser) *v1.AccountUser {
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

func (rp *accountUserRepo) Create(ctx context.Context, req *v1.CreateAccountUserReq) (*v1.AccountUser, error) {
	po, err := rp.data.db.AccountUser.
		Create().
		SetNillableUsername(req.AccountUser.Username).
		SetNillableEmail(req.AccountUser.Email).
		SetNillablePhone(req.AccountUser.Phone).
		SetNillablePassword(req.AccountUser.Password).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return rp.convertEntToProto(po), nil
}

func (rp *accountUserRepo) Update(ctx context.Context, req *v1.UpdateAccountUserReq) (*v1.AccountUser, error) {
	builder := rp.data.db.AccountUser.UpdateOneID(req.AccountUser.Id).
		SetNillableUsername(req.AccountUser.Username)

	po, err := builder.Save(ctx)
	if err != nil {
		return nil, err
	}

	return rp.convertEntToProto(po), nil
}

func (rp *accountUserRepo) Delete(ctx context.Context, id uint32) error {
	return rp.data.db.AccountUser.DeleteOneID(id).Exec(ctx)
}

func (rp *accountUserRepo) Get(ctx context.Context, id uint32) (*v1.AccountUser, error) {
	po, err := rp.data.db.AccountUser.Get(ctx, id)
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	}

	return rp.convertEntToProto(po), nil
}

func (rp *accountUserRepo) List(ctx context.Context, req *v1.ListAccountUserReq) (*v1.ListAccountUserResp, error) {
	whereCond, orderCond := entgo.QueryCommandToSelector(req.GetQuery(), req.GetOrderBy())

	builder1 := rp.data.db.AccountUser.Query()
	if len(whereCond) != 0 {
		for _, v := range whereCond {
			builder1.Where(v)
		}
	}
	if len(orderCond) != 0 {
		for _, v := range orderCond {
			builder1.Order(v)
		}
	}

	accountUsers, err := builder1.All(ctx)
	if err != nil {
		return nil, err
	}

	items := make([]*v1.AccountUser, 0, len(accountUsers))
	for _, po := range accountUsers {
		item := rp.convertEntToProto(po)
		items = append(items, item)
	}

	return &v1.ListAccountUserResp{
		List: items,
	}, nil
}

func (rp *accountUserRepo) PageList(ctx context.Context, req *v1.PageListAccountUserReq) (*v1.PageListAccountUserResp, error) {
	whereCond, orderCond := entgo.QueryCommandToSelector(req.GetQuery(), req.GetOrderBy())

	builder1 := rp.data.db.AccountUser.Query()
	if len(whereCond) != 0 {
		for _, v := range whereCond {
			builder1.Where(v)
		}
	}
	if len(orderCond) != 0 {
		for _, v := range orderCond {
			builder1.Order(v)
		}
	}

	builder1.
		Offset(paging.GetPageOffset(req.GetPageSize(), req.GetPageSize())).
		Limit(int(req.GetPageSize()))

	accountUsers, err := builder1.All(ctx)
	if err != nil {
		return nil, err
	}

	builder2 := rp.data.db.AccountUser.Query()
	if len(whereCond) != 0 {
		for _, v := range whereCond {
			builder2.Where(v)
		}
	}

	count, err := builder2.
		Select(accountuser.FieldID).
		Count(ctx)
	if err != nil {
		return nil, err
	}

	items := make([]*v1.AccountUser, 0, len(accountUsers))
	for _, po := range accountUsers {
		item := rp.convertEntToProto(po)
		items = append(items, item)
	}

	return &v1.PageListAccountUserResp{
		PageNum:  req.GetPageNum(),
		PageSize: req.GetPageSize(),
		Total:    int32(count),
		List:     items,
	}, err
}

func (rp *accountUserRepo) FindByUsername(ctx context.Context, username string) (*v1.AccountUser, error) {
	po, err := rp.data.db.AccountUser.
		Query().
		Where(accountuser.UsernameEQ(username)).
		Only(ctx)
	if err != nil {
		return nil, biz.ErrUserNotFound
	}

	return rp.convertEntToProto(po), nil
}

func (rp *accountUserRepo) FindByEmail(ctx context.Context, email string) (*v1.AccountUser, error) {
	po, err := rp.data.db.AccountUser.
		Query().
		Where(accountuser.EmailEQ(email)).
		Only(ctx)
	if err != nil {
		return nil, biz.ErrUserNotFound
	}

	return rp.convertEntToProto(po), nil
}

func (rp *accountUserRepo) FindByPhone(ctx context.Context, phone string) (*v1.AccountUser, error) {
	po, err := rp.data.db.AccountUser.
		Query().
		Where(accountuser.PhoneEQ(phone)).
		Only(ctx)
	if err != nil {
		return nil, biz.ErrUserNotFound
	}

	return rp.convertEntToProto(po), nil
}
