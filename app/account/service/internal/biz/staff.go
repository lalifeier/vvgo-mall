package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	pb "github.com/lalifeier/vvgo-mall/api/account/service/v1"
)

type Staff struct {
	Id int64
}

type StaffRepo interface {
	CreateStaff(ctx context.Context, staff *Staff) (*Staff, error)
	UpdateStaff(ctx context.Context, staff *Staff) (*Staff, error)
	DeleteStaff(ctx context.Context, id int64) error
	GetStaff(ctx context.Context, id int64) (*Staff, error)
	ListStaff(ctx context.Context, req *pb.ListStaffReq) ([]*Staff, error)
	PageListStaff(ctx context.Context, req *pb.PageListStaffReq) ([]*Staff, int64, error)
}

type StaffUsecase struct {
	repo StaffRepo
	log  *log.Helper
}

func NewStaffUsecase(repo StaffRepo, logger log.Logger) *StaffUsecase {
	return &StaffUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *StaffUsecase) CreateStaff(ctx context.Context, staff *Staff) (*Staff, error) {
	return uc.repo.CreateStaff(ctx, staff)
}

func (uc *StaffUsecase) UpdateStaff(ctx context.Context, staff *Staff) (*Staff, error) {
	return uc.repo.UpdateStaff(ctx, staff)
}

func (uc *StaffUsecase) DeleteStaff(ctx context.Context, id int64) error {
	return uc.repo.DeleteStaff(ctx, id)
}

func (uc *StaffUsecase) GetStaff(ctx context.Context, id int64) (*Staff, error) {
	return uc.repo.GetStaff(ctx, id)
}

func (uc *StaffUsecase) ListStaff(ctx context.Context, req *pb.ListStaffReq) ([]*Staff, error) {
	return uc.repo.ListStaff(ctx, req)
}

func (uc *StaffUsecase) PageListStaff(ctx context.Context, req *pb.PageListStaffReq) ([]*Staff, int64, error) {
	return uc.repo.PageListStaff(ctx, req)
}
