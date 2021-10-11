package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	pb "github.com/lalifeier/vgo/api/account/interface/v1"
	"github.com/lalifeier/vgo/app/account/interface/internal/biz"
)

type UserService struct {
	pb.UnimplementedUserInterfaceServer

	uc *biz.UserUsecase

	log *log.Helper
}

func NewUserService(uc *biz.UserUsecase, logger log.Logger) *UserService {
	return &UserService{
		uc:  uc,
		log: log.NewHelper(log.With(logger, "module", "user-interface/service")),
	}
}

func (s *UserService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterReply, error) {
	rv, err := s.uc.Register(ctx, &biz.User{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	return &pb.RegisterReply{
		Id: rv.Id,
	}, err
}

func (s *UserService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginReply, error) {
	rv, err := s.uc.Login(ctx, &biz.User{
		Username: req.Username,
		Password: req.Password,
	})
	return &pb.LoginReply{
		Token: rv,
	}, err
}

func (s *UserService) Logout(ctx context.Context, req *pb.LogoutRequest) (*pb.LogoutReply, error) {
	err := s.uc.Logout(ctx, &biz.User{})
	return &pb.LogoutReply{}, err
}
