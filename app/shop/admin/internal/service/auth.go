package service

import (
	"context"

	"github.com/lalifeier/vvgo-mall/app/shop/admin/internal/biz"
	pb "github.com/lalifeier/vvgo-mall/gen/api/go/shop/admin/v1"
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
