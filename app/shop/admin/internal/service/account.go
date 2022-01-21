package service

import (
	"context"

	pb "github.com/lalifeier/vvgo-mall/api/shop/admin/v1"
)

type AccountService struct {
	pb.UnimplementedAccountServer
}

func NewAccountService() *AccountService {
	return &AccountService{}
}

func (s *AccountService) Register(ctx context.Context, req *pb.RegisterReq) (*pb.RegisterResp, error) {
	return &pb.RegisterResp{}, nil
}
func (s *AccountService) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginResp, error) {
	return &pb.LoginResp{}, nil
}
func (s *AccountService) CreateAccountUser(ctx context.Context, req *pb.CreateAccountUserReq) (*pb.CreateAccountUserResp, error) {
	return &pb.CreateAccountUserResp{}, nil
}
func (s *AccountService) UpdateAccountUser(ctx context.Context, req *pb.UpdateAccountUserReq) (*pb.UpdateAccountUserResp, error) {
	return &pb.UpdateAccountUserResp{}, nil
}
func (s *AccountService) DeleteAccountUser(ctx context.Context, req *pb.DeleteAccountUserReq) (*pb.DeleteAccountUserResp, error) {
	return &pb.DeleteAccountUserResp{}, nil
}
func (s *AccountService) GetAccountUser(ctx context.Context, req *pb.GetAccountUserReq) (*pb.GetAccountUserResp, error) {
	return &pb.GetAccountUserResp{
		Id: req.Id,
	}, nil
}
func (s *AccountService) ListAccountUser(ctx context.Context, req *pb.ListAccountUserReq) (*pb.ListAccountUserResp, error) {
	return &pb.ListAccountUserResp{}, nil
}
func (s *AccountService) PageListAccountUser(ctx context.Context, req *pb.PageListAccountUserReq) (*pb.PageListAccountUserResp, error) {
	return &pb.PageListAccountUserResp{}, nil
}
func (s *AccountService) CreateStaff(ctx context.Context, req *pb.CreateStaffReq) (*pb.CreateStaffResp, error) {
	return &pb.CreateStaffResp{}, nil
}
func (s *AccountService) UpdateStaff(ctx context.Context, req *pb.UpdateStaffReq) (*pb.UpdateStaffResp, error) {
	return &pb.UpdateStaffResp{}, nil
}
func (s *AccountService) DeleteStaff(ctx context.Context, req *pb.DeleteStaffReq) (*pb.DeleteStaffResp, error) {
	return &pb.DeleteStaffResp{}, nil
}
func (s *AccountService) GetStaff(ctx context.Context, req *pb.GetStaffReq) (*pb.GetStaffResp, error) {
	return &pb.GetStaffResp{}, nil
}
func (s *AccountService) ListStaff(ctx context.Context, req *pb.ListStaffReq) (*pb.ListStaffResp, error) {
	return &pb.ListStaffResp{}, nil
}
func (s *AccountService) PageListStaff(ctx context.Context, req *pb.PageListStaffReq) (*pb.PageListStaffResp, error) {
	return &pb.PageListStaffResp{}, nil
}
