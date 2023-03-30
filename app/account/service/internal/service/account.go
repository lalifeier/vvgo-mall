package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/lalifeier/vvgo-mall/app/account/service/internal/biz"
	pb "github.com/lalifeier/vvgo-mall/gen/api/go/account/service/v1"
	v1 "github.com/lalifeier/vvgo-mall/gen/api/go/account/service/v1"
)

type AccountService struct {
	v1.UnimplementedAccountServiceServer
	log *log.Helper

	accountUserUseCase *biz.AccountUserUseCase

	authUseCase *biz.AuthUseCase
}

func NewAccountService(logger log.Logger, accountUserUseCase *biz.AccountUserUseCase, authUseCase *biz.AuthUseCase) *AccountService {
	return &AccountService{
		log:                log.NewHelper(log.With(logger, "module", "account-service/service")),
		accountUserUseCase: accountUserUseCase,
		authUseCase:        authUseCase,
	}
}

func (s *AccountService) Register(ctx context.Context, req *pb.RegisterReq) (*pb.RegisterResp, error) {
	return nil, nil
	// id, err := s.authUseCase.Register(ctx, &biz.AccountUser{
	// 	Username: req.Username,
	// 	Password: req.Password,
	// 	Email:    req.Email,
	// 	Phone:    req.Phone,
	// })
	// if err != nil {
	// 	return nil, err
	// }
	// return &pb.RegisterResp{
	// 	Id: id,
	// }, nil
}

func (s *AccountService) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginResp, error) {
	return nil, nil
	// u, err := s.authUseCase.Login(ctx, &biz.AccountUser{
	// 	Username: req.Username,
	// 	Password: req.Password,
	// 	Email:    req.Email,
	// 	Phone:    req.Phone,
	// })
	// if err != nil {
	// 	return nil, err
	// }
	// return &pb.LoginResp{
	// 	Id:       u.Id,
	// 	Username: req.Username,
	// 	Email:    req.Email,
	// 	Phone:    req.Phone,
	// }, nil
}
