package service

import (
	"context"

	"github.com/lalifeier/vvgo-mall/app/account/service/internal/biz"
	pb "github.com/lalifeier/vvgo-mall/gen/api/account/service/v1"
)

func (s *AccountService) Register(ctx context.Context, req *pb.RegisterReq) (*pb.RegisterResp, error) {
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

func (s *AccountService) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginResp, error) {
	u, err := s.authUseCase.Login(ctx, &biz.AccountUser{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
		Phone:    req.Phone,
	})
	if err != nil {
		return nil, err
	}
	return &pb.LoginResp{
		Id:       u.Id,
		Username: req.Username,
		Email:    req.Email,
		Phone:    req.Phone,
	}, nil
}
