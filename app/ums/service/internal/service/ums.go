package service

import (
	"context"
	"errors"

	"github.com/go-kratos/kratos/v2/log"

	pb "github.com/lalifeier/vvgo-mall/api/ums/service/v1"
	"github.com/lalifeier/vvgo-mall/app/ums/service/internal/biz"
)

// type bizAccountUser biz.AccountUser

// func (b bizAccountUser) PbAccountUser() *pb.AccountUser {
// 	return &pb.AccountUser{
// 		Id:       b.Id,
// 		Username: b.Username,
// 		Password: b.Password,
// 		Phone:    b.Phone,
// 		Email:    b.Email,
// 	}
// }

type UmsService struct {
	pb.UnimplementedUmsServer

	log                *log.Helper
	accountUserUseCase *biz.AccountUserUseCase
}

func NewUmsService(accountUserUseCase *biz.AccountUserUseCase, logger log.Logger) *UmsService {
	return &UmsService{
		log:                log.NewHelper(log.With(logger, "module", "ums-service/service")),
		accountUserUseCase: accountUserUseCase,
	}
}

func (s *UmsService) Register(ctx context.Context, req *pb.RegisterReq) (*pb.RegisterResp, error) {
	return &pb.RegisterResp{}, nil
}

func (s *UmsService) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginResp, error) {
	_, err := s.accountUserUseCase.Login(ctx, req.Username, req.Password)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return &pb.LoginResp{}, nil
}

func (s *UmsService) Logout(ctx context.Context, req *pb.LogoutReq) (*pb.LogoutResp, error) {
	return &pb.LogoutResp{}, nil
}
