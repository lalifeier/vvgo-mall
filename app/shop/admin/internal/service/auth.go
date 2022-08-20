package service

import (
	"context"

	pb "github.com/lalifeier/vvgo-mall/api/shop/admin/v1"
	"github.com/lalifeier/vvgo-mall/app/shop/admin/internal/biz"
)

func (s *ShopService) Register(ctx context.Context, req *pb.RegisterReq) (*pb.RegisterResp, error) {
	id, err := s.authUseCase.Register(ctx, &biz.AccountUser{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
		Phone:    req.Phone,
	})
	if err != nil {
		return nil, err
	}
	return &pb.RegisterResp{
		Id: id,
	}, nil
}

func (s *ShopService) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginResp, error) {
	user, err := s.authUseCase.Login(ctx, &biz.AccountUser{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
		Phone:    req.Phone,
	})
	if err != nil {
		return nil, err
	}

	return &pb.LoginResp{Id: user.Id, Username: user.Username, Email: user.Email, Phone: user.Phone, Token: user.Token}, nil
}
