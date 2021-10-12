package service

import (
	"context"

	pb "github.com/lalifeier/vgo/api/account/service/v1"
	"github.com/lalifeier/vgo/app/account/service/internal/biz"
)

func (s *AccountService) CreateAccountUser(ctx context.Context, req *pb.CreateAccountUserRequest) (*pb.CreateAccountUserReply, error) {
	rv, err := s.auc.Create(ctx, &biz.AccountUser{
		Email:    req.Email,
		Phone:    req.Phone,
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateAccountUserReply{
		Id: rv.Id,
	}, err
}

func (s *AccountService) UpdateAccountUser(ctx context.Context, req *pb.UpdateAccountUserRequest) (*pb.UpdateAccountUserReply, error) {
	return &pb.UpdateAccountUserReply{}, nil
}
func (s *AccountService) DeleteAccountUser(ctx context.Context, req *pb.DeleteAccountUserRequest) (*pb.DeleteAccountUserReply, error) {
	return &pb.DeleteAccountUserReply{}, nil
}
func (s *AccountService) GetAccountUser(ctx context.Context, req *pb.GetAccountUserRequest) (*pb.GetAccountUserReply, error) {
	return &pb.GetAccountUserReply{}, nil
}
func (s *AccountService) ListAccountUser(ctx context.Context, req *pb.ListAccountUserRequest) (*pb.ListAccountUserReply, error) {
	return &pb.ListAccountUserReply{}, nil
}
