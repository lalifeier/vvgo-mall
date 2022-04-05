package service

import (
	"context"

	pb "github.com/lalifeier/vvgo-mall/api/shop/admin/v1"
)

func (s *ShopService) Register(ctx context.Context, req *pb.RegisterReq) (*pb.RegisterResp, error) {
	return &pb.RegisterResp{}, nil
}
func (s *ShopService) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginResp, error) {
	// return &pb.LoginResp{}, nil
	return s.authUsecase.Login(ctx, req)
}
