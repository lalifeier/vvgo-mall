package service

import (
	"context"

	"github.com/jinzhu/copier"
	pb "github.com/lalifeier/vvgo-mall/api/taoke/interface/v1"
)

func (s *TaoKeService) GetCarouseList(ctx context.Context, req *pb.GetCarouseListReq) (*pb.GetCarouseListResp, error) {
	var (
		params = make(map[string]string)
	)
	carouseList, err := s.taoBaoUseCase.CarouseList(ctx, params)
	if err != nil {
		return nil, err
	}

	data := make([]*pb.Carouse, 0)
	err = copier.Copy(&data, &carouseList)
	if err != nil {
		return nil, err
	}

	return &pb.GetCarouseListResp{
		Data: data,
	}, nil
}

func (s *TaoKeService) GetSuperCategory(ctx context.Context, req *pb.GetSuperCategoryReq) (*pb.GetSuperCategoryResp, error) {
	var (
		params = make(map[string]string)
	)
	categoriesList, err := s.taoBaoUseCase.GetSuperCategory(ctx, params)
	if err != nil {
		return nil, err
	}

	data := make([]*pb.Categories, 0)
	err = copier.Copy(&data, &categoriesList)
	if err != nil {
		return nil, err
	}

	return &pb.GetSuperCategoryResp{
		Data: data,
	}, nil
}
