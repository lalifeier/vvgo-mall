package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"
	pb "github.com/lalifeier/vvgo-mall/api/shop/admin/v1"
	"github.com/lalifeier/vvgo-mall/app/shop/admin/internal/biz"
)

type ShopService struct {
	pb.UnimplementedShopServer

	log *log.Helper

	authUsecase *biz.AuthUsecase
	dictUsecase *biz.DictUsecase
}

func NewShopService(logger log.Logger, authUsecase *biz.AuthUsecase, dictUsecase *biz.DictUsecase) *ShopService {
	return &ShopService{
		log:         log.NewHelper(log.With(logger, "module", "shop-admin/service")),
		authUsecase: authUsecase,
		dictUsecase: dictUsecase,
	}
}

func (s *ShopService) CreateDict(ctx context.Context, req *pb.CreateDictReq) (*pb.CreateDictResp, error) {
	id, err := s.dictUsecase.CreateDict(ctx, &biz.Dict{
		DictTypeId: req.DictTypeId,
		Label:      req.Label,
		Value:      req.Value,
		Status:     int8(req.Status),
		Remark:     req.Remark,
		Sort:       int8(req.Sort),
		IsDefault:  int8(req.IsDefault),
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateDictResp{
		Id: id,
	}, nil
}
func (s *ShopService) UpdateDict(ctx context.Context, req *pb.UpdateDictReq) (*pb.UpdateDictResp, error) {
	err := s.dictUsecase.UpdateDict(ctx, &biz.Dict{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	return &pb.UpdateDictResp{}, nil
}
func (s *ShopService) DeleteDict(ctx context.Context, req *pb.DeleteDictReq) (*pb.DeleteDictResp, error) {
	err := s.dictUsecase.DeleteDict(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteDictResp{}, nil
}
func (s *ShopService) ListDict(ctx context.Context, req *pb.ListDictReq) (*pb.ListDictResp, error) {
	rv, err := s.dictUsecase.ListDict(ctx, &biz.DictListReq{})
	if err != nil {
		return nil, err
	}

	dicts := make([]*pb.Dict, 0)
	err = copier.Copy(&dicts, &rv)
	if err != nil {
		return nil, err
	}

	return &pb.ListDictResp{
		Data: dicts,
	}, nil
}
func (s *ShopService) PageListDict(ctx context.Context, req *pb.PageListDictReq) (*pb.PageListDictResp, error) {
	pos, err := s.dictUsecase.PageListDict(ctx, &biz.DictPageListReq{PageNum: req.PageNum, PageSize: req.PageSize})
	if err != nil {
		return nil, err
	}

	dicts := make([]*pb.Dict, 0)
	err = copier.Copy(&dicts, &pos.Data)
	if err != nil {
		return nil, err
	}

	return &pb.PageListDictResp{
		TotalPages:  pos.TotalPages,
		CurrentPage: req.PageSize,
		PageSize:    req.PageNum,
		Data:        dicts,
	}, nil
}
func (s *ShopService) GetDict(ctx context.Context, req *pb.GetDictReq) (*pb.GetDictResp, error) {
	au, err := s.dictUsecase.GetDict(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	dict := pb.GetDictResp{}
	err = copier.Copy(&dict, au)
	if err != nil {
		return nil, err
	}

	return &dict, nil
}
