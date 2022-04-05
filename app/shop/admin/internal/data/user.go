package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/lalifeier/vvgo-mall/app/shop/admin/internal/biz"
	"github.com/lalifeier/vvgo-mall/app/shop/admin/internal/data/ent"

	"github.com/lalifeier/vvgo-mall/pkg/util/pagination"
)

var _ biz.UserRepo = (*userRepo)(nil)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "shop-admin/data")),
	}
}

type entUser ent.User

func (e entUser) BizUser() *biz.User {
	return &biz.User{
		Id: e.ID,
	}
}

func (rp *userRepo) CreateUser(ctx context.Context, item *biz.User) (int64, error) {
	po, err := rp.data.db.User.
		Create().
		Save(ctx)
	if err != nil {
		return 0, err
	}
	return po.ID, err
}

func (rp *userRepo) UpdateUser(ctx context.Context, d *biz.User) error {
	_, err := rp.data.db.User.Get(ctx, d.Id)
	if err != nil {
		return err
	}
	db := rp.data.db.User.UpdateOneID(d.Id)

	_, err = db.Save(ctx)
	return err
}

func (rp *userRepo) DeleteUser(ctx context.Context, id int64) error {
	return rp.data.db.User.DeleteOneID(id).Exec(ctx)
}

func (rp *userRepo) GetUser(ctx context.Context, id int64) (*biz.User, error) {
	po, err := rp.data.db.User.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return entUser(*po).BizUser(), nil
}

func (rp *userRepo) ListUser(ctx context.Context, req *biz.UserListReq) ([]*biz.User, error) {
	query := rp.data.db.User.Query()

	pos, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	rv := make([]*biz.User, 0)
	for _, po := range pos {
		rv = append(rv, entUser(*po).BizUser())
	}
	return rv, nil
}

func (rp *userRepo) PageListUser(ctx context.Context, req *biz.UserPageListReq) (*biz.UserPageListResp, error) {
	query := rp.data.db.User.Query()

	query.Order(ent.Desc("id"))
	total, err := query.Count(ctx)
	if err != nil {
		return nil, err
	}

	pos, err := query.
		Offset(int(pagination.GetPageOffset(req.PageNum, req.PageSize))).
		Limit(int(req.PageSize)).
		All(ctx)
	if err != nil {
		return nil, err
	}

	rv := make([]*biz.User, 0)
	for _, po := range pos {
		rv = append(rv, entUser(*po).BizUser())
	}
	return &biz.UserPageListResp{
		CurrentPage: req.PageNum,
		PageSize:    req.PageSize,
		TotalPages:  pagination.GetTotalPages(int64(total), req.PageSize),
		Data:        rv,
	}, nil
}
