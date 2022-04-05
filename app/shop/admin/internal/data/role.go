package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/lalifeier/vvgo-mall/app/shop/admin/internal/biz"
	"github.com/lalifeier/vvgo-mall/app/shop/admin/internal/data/ent"

	"github.com/lalifeier/vvgo-mall/pkg/util/pagination"
)

var _ biz.RoleRepo = (*roleRepo)(nil)

type roleRepo struct {
	data *Data
	log  *log.Helper
}

func NewRoleRepo(data *Data, logger log.Logger) biz.RoleRepo {
	return &roleRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "shop-admin/data")),
	}
}

type entRole ent.Role

func (e entRole) BizRole() *biz.Role {
	return &biz.Role{
		Id: e.ID,
	}
}

func (rp *roleRepo) CreateRole(ctx context.Context, d *biz.Role) (int64, error) {
	po, err := rp.data.db.Role.
		Create().
		Save(ctx)
	if err != nil {
		return 0, err
	}
	return po.ID, err
}

func (rp *roleRepo) UpdateRole(ctx context.Context, d *biz.Role) error {
	_, err := rp.data.db.Role.Get(ctx, d.Id)
	if err != nil {
		return err
	}
	db := rp.data.db.Role.UpdateOneID(d.Id)

	_, err = db.Save(ctx)
	return err
}

func (rp *roleRepo) DeleteRole(ctx context.Context, id int64) error {
	return rp.data.db.Role.DeleteOneID(id).Exec(ctx)
}

func (rp *roleRepo) GetRole(ctx context.Context, id int64) (*biz.Role, error) {
	po, err := rp.data.db.Role.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return entRole(*po).BizRole(), nil
}

func (rp *roleRepo) ListRole(ctx context.Context, req *biz.RoleListReq) ([]*biz.Role, error) {
	query := rp.data.db.Role.Query()

	pos, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	rv := make([]*biz.Role, 0)
	for _, po := range pos {
		rv = append(rv, entRole(*po).BizRole())
	}
	return rv, nil
}

func (rp *roleRepo) PageListRole(ctx context.Context, req *biz.RolePageListReq) (*biz.RolePageListResp, error) {
	query := rp.data.db.Role.Query()

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

	rv := make([]*biz.Role, 0)
	for _, po := range pos {
		rv = append(rv, entRole(*po).BizRole())
	}
	return &biz.RolePageListResp{
		CurrentPage: req.PageNum,
		PageSize:    req.PageSize,
		TotalPages:  pagination.GetTotalPages(int64(total), req.PageSize),
		Data:        rv,
	}, nil
}
