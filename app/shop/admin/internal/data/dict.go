package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/lalifeier/vvgo-mall/app/shop/admin/internal/biz"
	"github.com/lalifeier/vvgo-mall/app/shop/admin/internal/data/ent"

	"github.com/lalifeier/vvgo-mall/pkg/utils/pagination"
)

var _ biz.DictRepo = (*dictRepo)(nil)

type dictRepo struct {
	data *Data
	log  *log.Helper
}

func NewDictRepo(data *Data, logger log.Logger) biz.DictRepo {
	return &dictRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "sys-service/data")),
	}
}

type entDict ent.Dict

func (e entDict) BizDict() *biz.Dict {
	return &biz.Dict{
		Id:        e.ID,
		Label:     e.Label,
		Value:     e.Value,
		Status:    int(e.Status),
		Remark:    e.Remark,
		Sort:      int(e.Sort),
		IsDefault: int(e.IsDefault),
	}
}

func (rp *dictRepo) CreateDict(ctx context.Context, d *biz.Dict) (int64, error) {
	po, err := rp.data.db.Dict.
		Create().
		SetLabel(d.Label).
		SetValue(d.Value).
		SetStatus(int8(d.Status)).
		SetRemark(d.Remark).
		SetSort(int8(d.Sort)).
		SetIsDefault(int8(d.IsDefault)).
		Save(ctx)
	if err != nil {
		return 0, err
	}
	return po.ID, err
}

func (rp *dictRepo) UpdateDict(ctx context.Context, d *biz.Dict) error {
	po, err := rp.data.db.Dict.Get(ctx, d.Id)
	if err != nil {
		return err
	}
	db := rp.data.db.Dict.UpdateOneID(d.Id)

	if d.Label != "" && d.Label != po.Label {
		db.SetLabel(d.Label)
	}

	_, err = db.Save(ctx)
	return err
}

func (rp *dictRepo) DeleteDict(ctx context.Context, id int64) error {
	return rp.data.db.Dict.DeleteOneID(id).Exec(ctx)
}

func (rp *dictRepo) GetDict(ctx context.Context, id int64) (*biz.Dict, error) {
	po, err := rp.data.db.Dict.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return entDict(*po).BizDict(), nil
}

func (rp *dictRepo) ListDict(ctx context.Context, req *biz.DictListReq) ([]*biz.Dict, error) {
	query := rp.data.db.Dict.Query()

	// if req.Type != "" {
	// 	query.Where(dict.TypeEQ(req.Type))
	// }

	pos, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	rv := make([]*biz.Dict, 0)
	for _, po := range pos {
		rv = append(rv, entDict(*po).BizDict())
	}
	return rv, nil
}

func (rp *dictRepo) PageListDict(ctx context.Context, req *biz.DictPageListReq) (*biz.DictPageListResp, error) {
	query := rp.data.db.Dict.Query()

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

	rv := make([]*biz.Dict, 0)
	for _, po := range pos {
		rv = append(rv, entDict(*po).BizDict())
	}
	return &biz.DictPageListResp{
		CurrentPage: req.PageNum,
		PageSize:    req.PageSize,
		TotalPages:  pagination.GetTotalPages(int64(total), req.PageSize),
		Data:        rv,
	}, nil
}
