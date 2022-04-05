package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/lalifeier/vvgo-mall/app/shop/admin/internal/biz"
	"github.com/lalifeier/vvgo-mall/app/shop/admin/internal/data/ent"
	"github.com/lalifeier/vvgo-mall/pkg/util/pagination"
)

var _ biz.ApiRepo = (*apiRepo)(nil)

type apiRepo struct {
	data *Data
	log  *log.Helper
}

func NewApiRepo(data *Data, logger log.Logger) biz.ApiRepo {
	return &apiRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "shop-admin/data")),
	}
}

type entApi ent.Api

func (e entApi) BizApi() *biz.Api {
	return &biz.Api{
		Id: e.ID,
	}
}

func (rp *apiRepo) CreateApi(ctx context.Context, d *biz.Api) (int64, error) {
	po, err := rp.data.db.Api.
		Create().
		Save(ctx)
	if err != nil {
		return 0, err
	}
	return po.ID, err
}

func (rp *apiRepo) UpdateApi(ctx context.Context, d *biz.Api) error {
	_, err := rp.data.db.Api.Get(ctx, d.Id)
	if err != nil {
		return err
	}
	db := rp.data.db.Api.UpdateOneID(d.Id)

	_, err = db.Save(ctx)
	return err
}

func (rp *apiRepo) DeleteApi(ctx context.Context, id int64) error {
	return rp.data.db.Api.DeleteOneID(id).Exec(ctx)
}

func (rp *apiRepo) GetApi(ctx context.Context, id int64) (*biz.Api, error) {
	po, err := rp.data.db.Api.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return entApi(*po).BizApi(), nil
}

func (rp *apiRepo) ListApi(ctx context.Context, req *biz.ApiListReq) ([]*biz.Api, error) {
	query := rp.data.db.Api.Query()

	pos, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	rv := make([]*biz.Api, 0)
	for _, po := range pos {
		rv = append(rv, entApi(*po).BizApi())
	}
	return rv, nil
}

func (rp *apiRepo) PageListApi(ctx context.Context, req *biz.ApiPageListReq) (*biz.ApiPageListResp, error) {
	query := rp.data.db.Api.Query()

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

	rv := make([]*biz.Api, 0)
	for _, po := range pos {
		rv = append(rv, entApi(*po).BizApi())
	}
	return &biz.ApiPageListResp{
		CurrentPage: req.PageNum,
		PageSize:    req.PageSize,
		TotalPages:  pagination.GetTotalPages(int64(total), req.PageSize),
		Data:        rv,
	}, nil
}
