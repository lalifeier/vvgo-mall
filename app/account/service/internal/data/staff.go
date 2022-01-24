package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	pb "github.com/lalifeier/vvgo-mall/api/account/service/v1"
	"github.com/lalifeier/vvgo-mall/app/account/service/internal/biz"
	"github.com/lalifeier/vvgo-mall/app/account/service/internal/data/ent"

	"github.com/lalifeier/vvgo-mall/pkg/utils/pagination"
)

var _ biz.StaffRepo = (*staffRepo)(nil)

type entStaff ent.Staff

func (e entStaff) BizStruct() *biz.Staff {
	return &biz.Staff{
		// Id: e.ID,
	}
}

type staffRepo struct {
	data *Data
	log  *log.Helper
}

func NewStaffRepo(data *Data, logger log.Logger) biz.StaffRepo {
	return &staffRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "account-service/data")),
	}
}

func (rp *staffRepo) CreateStaff(ctx context.Context, b *biz.Staff) (*biz.Staff, error) {
	p, err := rp.data.db.Staff.
		Create().
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return entStaff(*p).BizStruct(), nil
}

func (rp *staffRepo) UpdateStaff(ctx context.Context, b *biz.Staff) (*biz.Staff, error) {
	_, err := rp.data.db.Staff.Get(ctx, b.Id)
	if err != nil {
		return nil, err
	}
	db := rp.data.db.Staff.UpdateOneID(b.Id)

	p, err := db.Save(ctx)
	if err != nil {
		return nil, err
	}
	return entStaff(*p).BizStruct(), nil
}

func (rp *staffRepo) DeleteStaff(ctx context.Context, id int64) error {
	return rp.data.db.Staff.DeleteOneID(id).Exec(ctx)
}

func (rp *staffRepo) GetStaff(ctx context.Context, id int64) (*biz.Staff, error) {
	p, err := rp.data.db.Staff.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return entStaff(*p).BizStruct(), nil
}

func (rp *staffRepo) ListStaff(ctx context.Context, req *pb.ListStaffReq) ([]*biz.Staff, error) {
	query := rp.data.db.Staff.Query()

	ps, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	rv := make([]*biz.Staff, 0, len(ps))
	for _, p := range ps {
		rv = append(rv, entStaff(*p).BizStruct())
	}
	return rv, nil
}

func (rp *staffRepo) PageListStaff(ctx context.Context, req *pb.PageListStaffReq) ([]*biz.Staff, int64, error) {
	query := rp.data.db.Staff.Query()

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

	rv := make([]*biz.Staff, 0, len(ps))
	for _, p := range ps {
		rv = append(rv, entStaff(*p).BizStruct())
	}
	return rv, int64(total), nil
}
