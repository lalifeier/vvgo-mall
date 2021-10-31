package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"
	pb "github.com/lalifeier/vgo/api/shop/admin/v1"
	"github.com/lalifeier/vgo/app/shop/admin/internal/biz/ums"
	"google.golang.org/protobuf/types/known/emptypb"
)

type UmsService struct {
	pb.UnimplementedUmsServer

	log                *log.Helper
	accountUserUsecase *ums.AccountUserUsecase
}

func NewUmsService(logger log.Logger, accountUserUsecase *ums.AccountUserUsecase) *UmsService {
	return &UmsService{
		log:                log.NewHelper(log.With(logger, "module", "shop-admin/service/")),
		accountUserUsecase: accountUserUsecase,
	}
}

func (s *UmsService) CreateAccountUser(ctx context.Context, req *pb.CreateAccountUserReq) (*pb.CreateAccountUserResp, error) {
	id, err := s.accountUserUsecase.CreateAccountUser(ctx, &ums.AccountUser{})
	if err != nil {
		return nil, err
	}
	return &pb.CreateAccountUserResp{
		Id: id,
	}, nil
}
func (s *UmsService) UpdateAccountUser(ctx context.Context, req *pb.UpdateAccountUserReq) (*emptypb.Empty, error) {
	err := s.accountUserUsecase.UpdateAccountUser(ctx, &ums.AccountUser{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
func (s *UmsService) DeleteAccountUser(ctx context.Context, req *pb.DeleteAccountUserReq) (*emptypb.Empty, error) {
	err := s.accountUserUsecase.DeleteAccountUser(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
func (s *UmsService) GetAccountUser(ctx context.Context, req *pb.GetAccountUserReq) (*pb.GetAccountUserResp, error) {
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
func (s *UmsService) ListAccountUser(ctx context.Context, req *pb.ListAccountUserReq) (*pb.ListAccountUserResp, error) {
	resp, err := s.accountUserUsecase.ListAccountUser(ctx, &ums.AccountUserListReq{
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
	})
	if err != nil {
		return nil, err
	}

	accountUsers := make([]*pb.AccountUser, 0)
	err = copier.Copy(&accountUsers, &resp.Data)
	if err != nil {
		return nil, err
	}

	return &pb.ListAccountUserResp{
		Total:       resp.Total,
		CurrentPage: resp.CurrentPage,
		PageSize:    resp.PageSize,
		Data:        accountUsers,
	}, nil
}
