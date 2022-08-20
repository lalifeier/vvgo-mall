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

type DictTypeUseCase struct {
	repo DictTypeRepo
	log  *log.Helper
}

func NewDictTypeUseCase(repo DictTypeRepo, logger log.Logger) *DictTypeUseCase {
	return &DictTypeUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *DictTypeUseCase) CreateDictType(ctx context.Context, dicttype *DictType) (int64, error) {
	return uc.repo.CreateDictType(ctx, dicttype)
}

func (uc *DictTypeUseCase) UpdateDictType(ctx context.Context, dicttype *DictType) error {
	return uc.repo.UpdateDictType(ctx, dicttype)
}

func (uc *DictTypeUseCase) DeleteDictType(ctx context.Context, id int64) error {
	return uc.repo.DeleteDictType(ctx, id)
}

func (uc *DictTypeUseCase) GetDictType(ctx context.Context, id int64) (*DictType, error) {
	return uc.repo.GetDictType(ctx, id)
}

func (uc *DictTypeUseCase) ListDictType(ctx context.Context, req *DictTypeListReq) ([]*DictType, error) {
	return uc.repo.ListDictType(ctx, req)
}

func (uc *DictTypeUseCase) PageListDictType(ctx context.Context, req *DictTypePageListReq) (*DictTypePageListResp, error) {
	return uc.repo.PageListDictType(ctx, req)
}
