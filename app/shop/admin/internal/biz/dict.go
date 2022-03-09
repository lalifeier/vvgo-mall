package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type Dict struct {
	Id         int64
	DictTypeId int64
	Label      string
	Value      string
	IsDefault  int8
	Sort       int8
	Remark     string
	Status     int8
}

type DictListReq struct {
	DictTypeTd int64
}

type DictPageListReq struct {
	PageNum  int64
	PageSize int64
}

type DictPageListResp struct {
	TotalPages  int64
	CurrentPage int64
	PageSize    int64
	Data        []*Dict
}

type DictRepo interface {
	CreateDict(ctx context.Context, dict *Dict) (int64, error)
	UpdateDict(ctx context.Context, dict *Dict) error
	DeleteDict(ctx context.Context, id int64) error
	GetDict(ctx context.Context, id int64) (*Dict, error)
	ListDict(ctx context.Context, req *DictListReq) ([]*Dict, error)
	PageListDict(ctx context.Context, req *DictPageListReq) (*DictPageListResp, error)
}

type DictUsecase struct {
	repo DictRepo
	log  *log.Helper
}

func NewDictUsecase(repo DictRepo, logger log.Logger) *DictUsecase {
	return &DictUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *DictUsecase) CreateDict(ctx context.Context, dict *Dict) (int64, error) {
	return uc.repo.CreateDict(ctx, dict)
}

func (uc *DictUsecase) UpdateDict(ctx context.Context, dict *Dict) error {
	return uc.repo.UpdateDict(ctx, dict)
}

func (uc *DictUsecase) DeleteDict(ctx context.Context, id int64) error {
	return uc.repo.DeleteDict(ctx, id)
}

func (uc *DictUsecase) GetDict(ctx context.Context, id int64) (*Dict, error) {
	return uc.repo.GetDict(ctx, id)
}

func (uc *DictUsecase) ListDict(ctx context.Context, req *DictListReq) ([]*Dict, error) {
	return uc.repo.ListDict(ctx, req)
}

func (uc *DictUsecase) PageListDict(ctx context.Context, req *DictPageListReq) (*DictPageListResp, error) {
	return uc.repo.PageListDict(ctx, req)
}
