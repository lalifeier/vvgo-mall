package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"
	pb "github.com/lalifeier/vvgo-mall/api/account/service/v1"
	"github.com/lalifeier/vvgo-mall/app/account/service/internal/biz"
)

type AccountService struct {
	pb.UnimplementedAccountServer

	log *log.Helper

	accountUserUsecase *biz.AccountUserUsecase
}

func NewAccountService(logger log.Logger, accountUserUsecase *biz.AccountUserUsecase) *AccountService {
	return &AccountService{
		log:                log.NewHelper(log.With(logger, "module", "account-service/service")),
		accountUserUsecase: accountUserUsecase,
	}
}

func (s *AccountService) Register(ctx context.Context, req *pb.RegisterReq) (*pb.RegisterResp, error) {
	return &pb.RegisterResp{}, nil
}
func (s *AccountService) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginResp, error) {
	return &pb.LoginResp{}, nil
}
func (s *AccountService) CreateAccountUser(ctx context.Context, req *pb.CreateAccountUserReq) (*pb.CreateAccountUserResp, error) {
	_, err := s.accountUserUsecase.CreateAccountUser(ctx, &biz.AccountUser{})
	if err != nil {
		return nil, err
	}
	return &pb.CreateAccountUserResp{}, nil
}
func (s *AccountService) UpdateAccountUser(ctx context.Context, req *pb.UpdateAccountUserReq) (*pb.UpdateAccountUserResp, error) {
	err := s.accountUserUsecase.UpdateAccountUser(ctx, &biz.AccountUser{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	return &pb.UpdateAccountUserResp{}, nil
}
func (s *AccountService) DeleteAccountUser(ctx context.Context, req *pb.DeleteAccountUserReq) (*pb.DeleteAccountUserResp, error) {
	err := s.accountUserUsecase.DeleteAccountUser(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteAccountUserResp{}, nil
}
func (s *AccountService) GetAccountUser(ctx context.Context, req *pb.GetAccountUserReq) (*pb.GetAccountUserResp, error) {
	au, err := s.accountUserUsecase.GetAccountUser(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	accountUser := pb.GetAccountUserResp{}
	err = copier.Copy(&accountUser, au)
	if err != nil {
		return nil, err
	}
	return &accountUser, nil
}
func (s *AccountService) ListAccountUser(ctx context.Context, req *pb.ListAccountUserReq) (*pb.ListAccountUserResp, error) {
	rv, err := s.accountUserUsecase.ListAccountUser(ctx, req)
	if err != nil {
		return nil, err
	}

	return &pb.ListAccountUserResp{
		List: make([]*pb.AccountUser, 0, len(rv)),
	}, nil
}
func (s *AccountService) PageListAccountUser(ctx context.Context, req *pb.PageListAccountUserReq) (*pb.PageListAccountUserResp, error) {
	pos, total, err := s.accountUserUsecase.PageListAccountUser(ctx, req)
	if err != nil {
		return nil, err
	}

	return &pb.PageListAccountUserResp{
		List:     make([]*pb.AccountUser, 0, len(pos)),
		Total:    total,
		Page:     req.PageSize,
		PageSize: req.PageNum,
	}, nil
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
