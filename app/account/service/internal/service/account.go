package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	pb "github.com/lalifeier/vgo/api/account/service/v1"
	"github.com/lalifeier/vgo/app/account/service/internal/biz"
)

type AccountService struct {
	pb.UnimplementedAccountServer

	auc *biz.AccountUserUsecase
	log *log.Helper
}

func NewAccountService(auc *biz.AccountUserUsecase, logger log.Logger) *AccountService {
	return &AccountService{
		auc: auc,
		log: log.NewHelper(log.With(logger, "module", "account-service/service")),
	}
}

func (s *AccountService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterReply, error) {
	return &pb.RegisterReply{}, nil
}
func (s *AccountService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginReply, error) {
	return &pb.LoginReply{}, nil
}
func (s *AccountService) PlatformLogin(ctx context.Context, req *pb.PlatformLoginRequest) (*pb.PlatformLoginReply, error) {
	return &pb.PlatformLoginReply{}, nil
}
