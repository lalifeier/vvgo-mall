package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type DictType struct {
	Id     int64
	Name   string
	Type   string
	Status int8
	Remark string
}

type DictTypeListReq struct {
	Name   string
	Type   string
	Status int8
}

type DictTypePageListReq struct {
	PageNum  int64
	PageSize int64
}

type DictTypePageListResp struct {
	TotalPages  int64
	CurrentPage int64
	PageSize    int64
	Data        []*DictType
}

type DictTypeRepo interface {
	CreateDictType(ctx context.Context, dicttype *DictType) (int64, error)
	UpdateDictType(ctx context.Context, dicttype *DictType) error
	DeleteDictType(ctx context.Context, id int64) error
	GetDictType(ctx context.Context, id int64) (*DictType, error)
	ListDictType(ctx context.Context, req *DictTypeListReq) ([]*DictType, error)
	PageListDictType(ctx context.Context, req *DictTypePageListReq) (*DictTypePageListResp, error)
}

type DictTypeUsecase struct {
	repo DictTypeRepo
	log  *log.Helper
}

func NewDictTypeUsecase(repo DictTypeRepo, logger log.Logger) *DictTypeUsecase {
	return &DictTypeUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *DictTypeUsecase) CreateDictType(ctx context.Context, dicttype *DictType) (int64, error) {
	return uc.repo.CreateDictType(ctx, dicttype)
}

func (uc *DictTypeUsecase) UpdateDictType(ctx context.Context, dicttype *DictType) error {
	return uc.repo.UpdateDictType(ctx, dicttype)
}

func (uc *DictTypeUsecase) DeleteDictType(ctx context.Context, id int64) error {
	return uc.repo.DeleteDictType(ctx, id)
}

func (uc *DictTypeUsecase) GetDictType(ctx context.Context, id int64) (*DictType, error) {
	return uc.repo.GetDictType(ctx, id)
}

func (uc *DictTypeUsecase) ListDictType(ctx context.Context, req *DictTypeListReq) ([]*DictType, error) {
	return uc.repo.ListDictType(ctx, req)
}

func (uc *DictTypeUsecase) PageListDictType(ctx context.Context, req *DictTypePageListReq) (*DictTypePageListResp, error) {
	return uc.repo.PageListDictType(ctx, req)
}
