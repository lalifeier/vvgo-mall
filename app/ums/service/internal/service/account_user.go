package service

import (
	"context"
	"fmt"

	"github.com/jinzhu/copier"
	pb "github.com/lalifeier/vvgo-mall/api/ums/service/v1"
	"github.com/lalifeier/vvgo-mall/app/ums/service/internal/biz"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *UmsService) CreateAccountUser(ctx context.Context, req *pb.CreateAccountUserReq) (*pb.CreateAccountUserResp, error) {
	fmt.Println(req)
	id, err := s.accountUserUseCase.Create(ctx, &biz.AccountUser{
		Username: req.Username,
		Password: req.Password,
		Phone:    req.Phone,
		Email:    req.Email,
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateAccountUserResp{
		Id: id,
	}, nil
}

func (s *UmsService) UpdateAccountUser(ctx context.Context, req *pb.UpdateAccountUserReq) (*emptypb.Empty, error) {
	err := s.accountUserUseCase.Update(ctx, &biz.AccountUser{
		Id:       req.Id,
		Username: req.Username,
		Password: req.Password,
		Phone:    req.Phone,
		Email:    req.Email,
	})
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *UmsService) DeleteAccountUser(ctx context.Context, req *pb.DeleteAccountUserReq) (*emptypb.Empty, error) {
	err := s.accountUserUseCase.Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *UmsService) GetAccountUser(ctx context.Context, req *pb.GetAccountUserReq) (*pb.GetAccountUserResp, error) {
	au, err := s.accountUserUseCase.Get(ctx, req.Id)
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
	pos, err := s.accountUserUseCase.List(ctx, &biz.AccountUserListReq{PageNum: req.PageNum, PageSize: req.PageSize})
	if err != nil {
		return nil, err
	}

	accountUsers := make([]*pb.AccountUser, 0)
	err = copier.Copy(&accountUsers, &pos.Data)
	if err != nil {
		return nil, err
	}

	return &pb.ListAccountUserResp{
		TotalPages:  pos.TotalPages,
		CurrentPage: pos.CurrentPage,
		PageSize:    pos.PageSize,
		Data:        accountUsers,
	}, nil
}
