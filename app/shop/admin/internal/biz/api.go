package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type Api struct {
	Id int64
}

type ApiListReq struct {
	ApiTypeTd int64
}

type ApiPageListReq struct {
	PageNum  int64
	PageSize int64
}

type ApiPageListResp struct {
	TotalPages  int64
	CurrentPage int64
	PageSize    int64
	Data        []*Api
}

type ApiRepo interface {
	CreateApi(ctx context.Context, a *Api) (int64, error)
	UpdateApi(ctx context.Context, a *Api) error
	DeleteApi(ctx context.Context, id int64) error
	GetApi(ctx context.Context, id int64) (*Api, error)
	ListApi(ctx context.Context, req *ApiListReq) ([]*Api, error)
	PageListApi(ctx context.Context, req *ApiPageListReq) (*ApiPageListResp, error)
}

type ApiUseCase struct {
	repo ApiRepo
	log  *log.Helper
}

func NewApiUseCase(repo ApiRepo, logger log.Logger) *ApiUseCase {
	return &ApiUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *ApiUseCase) CreateApi(ctx context.Context, a *Api) (int64, error) {
	return uc.repo.CreateApi(ctx, a)
}

func (uc *ApiUseCase) UpdateApi(ctx context.Context, a *Api) error {
	return uc.repo.UpdateApi(ctx, a)
}

func (uc *ApiUseCase) DeleteApi(ctx context.Context, id int64) error {
	return uc.repo.DeleteApi(ctx, id)
}

func (uc *ApiUseCase) GetApi(ctx context.Context, id int64) (*Api, error) {
	return uc.repo.GetApi(ctx, id)
}

func (uc *ApiUseCase) ListApi(ctx context.Context, req *ApiListReq) ([]*Api, error) {
	return uc.repo.ListApi(ctx, req)
}

func (uc *ApiUseCase) PageListApi(ctx context.Context, req *ApiPageListReq) (*ApiPageListResp, error) {
	return uc.repo.PageListApi(ctx, req)
}
