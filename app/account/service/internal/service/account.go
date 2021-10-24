package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	pb "github.com/lalifeier/vgo/api/account/service/v1"
	"github.com/lalifeier/vgo/app/account/service/internal/biz"
	"github.com/micro/go-micro/errors"
	"google.golang.org/protobuf/types/known/emptypb"
	// "google.golang.org/protobuf/types/known/emptypb"
)

type bizAccountUser biz.AccountUser

func (b bizAccountUser) PbAccountUser() *pb.AccountUser {
	return &pb.AccountUser{
		Id:       b.Id,
		Username: b.Username,
		Password: b.Password,
		Phone:    b.Phone,
		Email:    b.Email,
	}
}

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
	_, err := s.auc.Login(ctx, req.Username, req.Password)
	if err != nil {
		return nil, errors.Unauthorized("UNAUTHORIZED", err.Error())
	}
	return &pb.LoginReply{}, nil
}

func (s *AccountService) PlatformLogin(ctx context.Context, req *pb.PlatformLoginRequest) (*pb.PlatformLoginReply, error) {
	return &pb.PlatformLoginReply{}, nil
}

func (s *AccountService) CreateAccountUser(ctx context.Context, req *pb.AccountUserRequest) (*pb.AccountUser, error) {
	au, err := s.auc.Create(ctx, &biz.AccountUser{})
	if err != nil {
		return nil, err
	}
	return bizAccountUser(*au).PbAccountUser(), nil
}

func (s *AccountService) UpdateAccountUser(ctx context.Context, req *pb.AccountUserRequest) (*pb.AccountUser, error) {
	au, err := s.auc.Update(ctx, req.Id, &biz.AccountUser{})
	if err != nil {
		return nil, err
	}
	return bizAccountUser(*au).PbAccountUser(), nil
}

func (s *AccountService) DeleteAccountUser(ctx context.Context, req *pb.DeleteAccountUserRequest) (*emptypb.Empty, error) {
	err := s.auc.Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *AccountService) GetAccountUser(ctx context.Context, req *pb.GetAccountUserRequest) (*pb.AccountUser, error) {
	au, err := s.auc.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return bizAccountUser(*au).PbAccountUser(), nil
}

func (s *AccountService) ListAccountUser(ctx context.Context, req *pb.AccountUserListRequest) (*pb.AccountUserList, error) {
	pos, err := s.auc.List(ctx, req.PageNum, req.PageSize)
	if err != nil {
		return nil, err
	}

	rv := make([]*pb.AccountUser, 0)
	for _, po := range pos {
		rv = append(rv, bizAccountUser(*po).PbAccountUser())
	}
	return &pb.AccountUserList{
		AccountUser: rv,
	}, nil
}
