package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type DictData struct {
	Id         int64
	DictTypeId int64
	Label      string
	Value      string
	IsDefault  int8
	Sort       int8
	Remark     string
	Status     int8
}

type DictDataListReq struct {
	DictTypeId int64
}

type DictDataPageListReq struct {
	PageNum  int64
	PageSize int64
}

type DictDataPageListResp struct {
	TotalPages  int64
	CurrentPage int64
	PageSize    int64
	Data        []*DictData
}

type DictDataRepo interface {
	CreateDictData(ctx context.Context, dict *DictData) (int64, error)
	UpdateDictData(ctx context.Context, dict *DictData) error
	DeleteDictData(ctx context.Context, id int64) error
	GetDictData(ctx context.Context, id int64) (*DictData, error)
	ListDictData(ctx context.Context, req *DictDataListReq) ([]*DictData, error)
	PageListDictData(ctx context.Context, req *DictDataPageListReq) (*DictDataPageListResp, error)
}

type DictDataUseCase struct {
	repo DictDataRepo
	log  *log.Helper
}

func NewDictDataUseCase(repo DictDataRepo, logger log.Logger) *DictDataUseCase {
	return &DictDataUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *DictDataUseCase) CreateDictData(ctx context.Context, dict *DictData) (int64, error) {
	return uc.repo.CreateDictData(ctx, dict)
}

func (uc *DictDataUseCase) UpdateDictData(ctx context.Context, dict *DictData) error {
	return uc.repo.UpdateDictData(ctx, dict)
}

func (uc *DictDataUseCase) DeleteDictData(ctx context.Context, id int64) error {
	return uc.repo.DeleteDictData(ctx, id)
}

func (uc *DictDataUseCase) GetDictData(ctx context.Context, id int64) (*DictData, error) {
	return uc.repo.GetDictData(ctx, id)
}

func (uc *DictDataUseCase) ListDictData(ctx context.Context, req *DictDataListReq) ([]*DictData, error) {
	return uc.repo.ListDictData(ctx, req)
}

func (uc *DictDataUseCase) PageListDictData(ctx context.Context, req *DictDataPageListReq) (*DictDataPageListResp, error) {
	return uc.repo.PageListDictData(ctx, req)
}
