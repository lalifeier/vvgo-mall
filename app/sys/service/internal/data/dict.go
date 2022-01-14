package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/lalifeier/vvgo/app/sys/service/internal/biz"
	"github.com/lalifeier/vvgo/app/sys/service/internal/data/ent"
	"github.com/lalifeier/vvgo/pkg/util/pagination"
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
		Id: e.ID,
	}
}

func (rp *dictRepo) CreateDict(ctx context.Context, dict *biz.Dict) (int64, error) {

	create := rp.data.db.Dict.Create()
	po, err := create.Save(ctx)
	if err != nil {
		return 0, err
	}
	return po.ID, err
}

func (rp *dictRepo) UpdateDict(ctx context.Context, dict *biz.Dict) error {
	_, err := rp.data.db.Dict.Get(ctx, dict.Id)
	if err != nil {
		return err
	}

	update := rp.data.db.Dict.Update()

	_, err = update.Save(ctx)
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
	pos, err := rp.data.db.Dict.Query().All(ctx)
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
	// count, err := rp.PageCountDict(ctx, req)
	// if err != nil {
	// 	return nil, err
	// }
	pos, err := rp.data.db.Dict.Query().
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
		TotalPages:  0,
		CurrentPage: req.PageNum,
		PageSize:    req.PageSize,
		Data:        rv,
	}, nil
}

func (rp *dictRepo) PageCountDict(ctx context.Context, req *biz.DictListReq) (int64, error) {
	count, err := rp.data.db.Dict.Query().Count(ctx)
	if err != nil {
		return 0, err
	}
	return int64(count), nil
}
