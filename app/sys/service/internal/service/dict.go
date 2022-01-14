package service

import (
	"context"

	"github.com/jinzhu/copier"
	pb "github.com/lalifeier/vvgo/api/sys/service/v1"
	"github.com/lalifeier/vvgo/app/sys/service/internal/biz"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *SysService) CreateDict(ctx context.Context, req *pb.CreateDictReq) (*pb.CreateDictResp, error) {
	id, err := s.dictUsecase.CreateDict(ctx, &biz.Dict{})
	if err != nil {
		return nil, err
	}
	return &pb.CreateDictResp{
		Id: id,
	}, nil
}

func (s *SysService) UpdateDict(ctx context.Context, req *pb.UpdateDictReq) (*emptypb.Empty, error) {
	err := s.dictUsecase.UpdateDict(ctx, &biz.Dict{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *SysService) DeleteDict(ctx context.Context, req *pb.DeleteDictReq) (*emptypb.Empty, error) {
	err := s.dictUsecase.DeleteDict(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *SysService) GetDict(ctx context.Context, req *pb.GetDictReq) (*pb.GetDictResp, error) {
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

func (s *SysService) ListDict(ctx context.Context, req *pb.ListDictReq) (*pb.ListDictResp, error) {
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

func (s *SysService) PageListDict(ctx context.Context, req *pb.PageListDictReq) (*pb.PageListDictResp, error) {
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
		CurrentPage: pos.CurrentPage,
		PageSize:    pos.PageSize,
		Data:        dicts,
	}, nil
}
