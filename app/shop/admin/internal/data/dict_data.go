package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/lalifeier/vvgo-mall/app/shop/admin/internal/biz"
	"github.com/lalifeier/vvgo-mall/app/shop/admin/internal/data/ent"
	"github.com/lalifeier/vvgo-mall/app/shop/admin/internal/data/ent/dictdata"

	"github.com/lalifeier/vvgo-mall/pkg/util/pagination"
)

var _ biz.DictDataRepo = (*dictdataRepo)(nil)

type dictdataRepo struct {
	data *Data
	log  *log.Helper
}

func NewDictDataRepo(data *Data, logger log.Logger) biz.DictDataRepo {
	return &dictdataRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "shop-admin/data")),
	}
}

type entDictData ent.DictData

func (e entDictData) BizDictData() *biz.DictData {
	return &biz.DictData{
		Id:        e.ID,
		Label:     e.Label,
		Value:     e.Value,
		Status:    e.Status,
		Remark:    e.Remark,
		Sort:      e.Sort,
		IsDefault: e.IsDefault,
	}
}

func (rp *dictdataRepo) CreateDictData(ctx context.Context, d *biz.DictData) (int64, error) {
	po, err := rp.data.db.DictData.
		Create().
		SetDictTypeID(d.DictTypeId).
		SetLabel(d.Label).
		SetValue(d.Value).
		SetStatus(d.Status).
		SetRemark(d.Remark).
		SetSort(d.Sort).
		SetIsDefault(d.IsDefault).
		Save(ctx)
	if err != nil {
		return 0, err
	}
	return po.ID, err
}

func (rp *dictdataRepo) UpdateDictData(ctx context.Context, d *biz.DictData) error {
	po, err := rp.data.db.DictData.Get(ctx, d.Id)
	if err != nil {
		return err
	}
	db := rp.data.db.DictData.UpdateOneID(d.Id)

	if d.Label != "" && d.Label != po.Label {
		db.SetLabel(d.Label)
	}
	if d.Value != "" && d.Value != po.Value {
		db.SetValue(d.Value)
	}

	_, err = db.Save(ctx)
	return err
}

func (rp *dictdataRepo) DeleteDictData(ctx context.Context, id int64) error {
	return rp.data.db.DictData.DeleteOneID(id).Exec(ctx)
}

func (rp *dictdataRepo) GetDictData(ctx context.Context, id int64) (*biz.DictData, error) {
	po, err := rp.data.db.DictData.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return entDictData(*po).BizDictData(), nil
}

func (rp *dictdataRepo) ListDictData(ctx context.Context, req *biz.DictDataListReq) ([]*biz.DictData, error) {
	query := rp.data.db.DictData.Query()

	if req.DictTypeId > 0 {
		query.Where(dictdata.DictTypeIDEQ(req.DictTypeId))
	}

	pos, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	rv := make([]*biz.DictData, 0)
	for _, po := range pos {
		rv = append(rv, entDictData(*po).BizDictData())
	}
	return rv, nil
}

func (rp *dictdataRepo) PageListDictData(ctx context.Context, req *biz.DictDataPageListReq) (*biz.DictDataPageListResp, error) {
	query := rp.data.db.DictData.Query()

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

	rv := make([]*biz.DictData, 0)
	for _, po := range pos {
		rv = append(rv, entDictData(*po).BizDictData())
	}
	return &biz.DictDataPageListResp{
		CurrentPage: req.PageNum,
		PageSize:    req.PageSize,
		TotalPages:  pagination.GetTotalPages(int64(total), req.PageSize),
		Data:        rv,
	}, nil
}
