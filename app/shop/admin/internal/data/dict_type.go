package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/lalifeier/vvgo-mall/app/shop/admin/internal/biz"
	"github.com/lalifeier/vvgo-mall/app/shop/admin/internal/data/ent"
	"github.com/lalifeier/vvgo-mall/app/shop/admin/internal/data/ent/dicttype"

	"github.com/lalifeier/vvgo-mall/pkg/util/pagination"
)

var _ biz.DictTypeRepo = (*dictTypeRepo)(nil)

type dictTypeRepo struct {
	data *Data
	log  *log.Helper
}

func NewDictTypeRepo(data *Data, logger log.Logger) biz.DictTypeRepo {
	return &dictTypeRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "shop-admin/data")),
	}
}

type entDictType ent.DictType

func (e entDictType) BizDictType() *biz.DictType {
	return &biz.DictType{
		Id:     e.ID,
		Name:   e.Name,
		Type:   e.Type,
		Status: e.Status,
		Remark: e.Remark,
	}
}

func (rp *dictTypeRepo) CreateDictType(ctx context.Context, d *biz.DictType) (int64, error) {
	po, err := rp.data.db.DictType.
		Create().
		SetName(d.Name).
		SetType(d.Type).
		SetStatus(d.Status).
		SetRemark(d.Remark).
		Save(ctx)
	if err != nil {
		return 0, err
	}
	return po.ID, err
}

func (rp *dictTypeRepo) UpdateDictType(ctx context.Context, d *biz.DictType) error {
	po, err := rp.data.db.DictType.Get(ctx, d.Id)
	if err != nil {
		return err
	}
	db := rp.data.db.DictType.UpdateOneID(d.Id)

	if d.Name != "" && d.Name != po.Name {
		db.SetName(d.Name)
	}
	if d.Type != "" && d.Type != po.Type {
		db.SetType(d.Type)
	}

	_, err = db.Save(ctx)
	return err
}

func (rp *dictTypeRepo) DeleteDictType(ctx context.Context, id int64) error {
	return rp.data.db.DictType.DeleteOneID(id).Exec(ctx)
}

func (rp *dictTypeRepo) GetDictType(ctx context.Context, id int64) (*biz.DictType, error) {
	po, err := rp.data.db.DictType.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return entDictType(*po).BizDictType(), nil
}

func (rp *dictTypeRepo) ListDictType(ctx context.Context, req *biz.DictTypeListReq) ([]*biz.DictType, error) {
	query := rp.data.db.DictType.Query()

	if req.Name != "" {
		query.Where(dicttype.NameEQ(req.Name))
	}
	if req.Type != "" {
		query.Where(dicttype.TypeEQ(req.Type))
	}
	if req.Status > 0 {
		query.Where(dicttype.StatusEQ(req.Status))
	}

	pos, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	rv := make([]*biz.DictType, 0)
	for _, po := range pos {
		rv = append(rv, entDictType(*po).BizDictType())
	}
	return rv, nil
}

func (rp *dictTypeRepo) PageListDictType(ctx context.Context, req *biz.DictTypePageListReq) (*biz.DictTypePageListResp, error) {
	query := rp.data.db.DictType.Query()

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

	rv := make([]*biz.DictType, 0)
	for _, po := range pos {
		rv = append(rv, entDictType(*po).BizDictType())
	}
	return &biz.DictTypePageListResp{
		CurrentPage: req.PageNum,
		PageSize:    req.PageSize,
		TotalPages:  pagination.GetTotalPages(int64(total), req.PageSize),
		Data:        rv,
	}, nil
}
