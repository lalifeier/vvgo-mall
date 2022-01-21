package sys

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"

	sysV1 "github.com/lalifeier/vvgo-mall/api/sys/service/v1"
	"github.com/lalifeier/vvgo-mall/app/shop/admin/internal/biz/sys"
	"github.com/lalifeier/vvgo-mall/app/shop/admin/internal/data"
)

var _ sys.DictRepo = (*dictRepo)(nil)

type dictRepo struct {
	data *data.Data
	log  *log.Helper
}

func NewDictRepo(data *data.Data, logger log.Logger) sys.DictRepo {
	return &dictRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "shop-admin/data/dict")),
	}
}

func (rp *dictRepo) CreateDict(ctx context.Context, b *sys.Dict) (int64, error) {
	resp, err := rp.data.SysClient.CreateDict(ctx, &sysV1.CreateDictReq{
		Type:      b.Type,
		Label:     b.Label,
		Value:     b.Value,
		Status:    int32(b.Status),
		Remark:    b.Remark,
		Sort:      int32(b.Sort),
		IsDefault: int32(b.IsDefault),
	})
	if err != nil {
		return 0, err
	}
	return resp.Id, nil
}

func (rp *dictRepo) UpdateDict(ctx context.Context, d *sys.Dict) error {
	_, err := rp.data.SysClient.UpdateDict(ctx, &sysV1.UpdateDictReq{
		Id:   d.Id,
		Type: d.Type,
	})
	return err
}

func (rp *dictRepo) DeleteDict(ctx context.Context, id int64) error {
	_, err := rp.data.SysClient.DeleteDict(ctx, &sysV1.DeleteDictReq{Id: id})
	return err
}

func (rp *dictRepo) GetDict(ctx context.Context, id int64) (*sys.Dict, error) {
	resp, err := rp.data.SysClient.GetDict(ctx, &sysV1.GetDictReq{Id: id})
	if err != nil {
		return nil, err
	}
	dict := sys.Dict{}
	err = copier.Copy(&dict, resp)
	if err != nil {
		return nil, err
	}
	return &dict, nil
}

func (rp *dictRepo) ListDict(ctx context.Context, req *sys.DictListReq) ([]*sys.Dict, error) {
	resp, err := rp.data.SysClient.ListDict(ctx, &sysV1.ListDictReq{
		Type: req.Type,
	})
	if err != nil {
		return nil, err
	}

	dicts := make([]*sys.Dict, 0)
	err = copier.Copy(&dicts, &resp.Data)
	if err != nil {
		return nil, err
	}
	return dicts, nil
}

func (rp *dictRepo) PageListDict(ctx context.Context, req *sys.DictPageListReq) (*sys.DictPageListResp, error) {
	resp, err := rp.data.SysClient.PageListDict(ctx, &sysV1.PageListDictReq{
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
	})

	if err != nil {
		return nil, err
	}

	accountUsers := make([]*sys.Dict, 0)
	err = copier.Copy(&accountUsers, &resp.Data)
	if err != nil {
		return nil, err
	}

	return &sys.DictPageListResp{
		TotalPages:  resp.TotalPages,
		CurrentPage: resp.CurrentPage,
		PageSize:    resp.PageSize,
		Data:        accountUsers,
	}, err
}
