package service

import (
	"context"

	pb "github.com/lalifeier/vvgo-mall/api/account/service/v1"
	"github.com/lalifeier/vvgo-mall/app/account/service/internal/biz"
	"google.golang.org/protobuf/types/known/emptypb"
)

type bizAccountUser biz.AccountUser

func (b bizAccountUser) protoStruct() *pb.AccountUser {
	return &pb.AccountUser{
		Id: b.Id,
	}
}

func (s *AccountService) CreateAccountUser(ctx context.Context, req *pb.AccountUser) (*pb.AccountUser, error) {
	p, err := s.accountUserUsecase.CreateAccountUser(ctx, &biz.AccountUser{})
	if err != nil {
		return nil, err
	}
	return bizAccountUser(*p).protoStruct(), nil
}
func (s *AccountService) UpdateAccountUser(ctx context.Context, req *pb.AccountUser) (*pb.AccountUser, error) {
	p, err := s.accountUserUsecase.UpdateAccountUser(ctx, &biz.AccountUser{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	return bizAccountUser(*p).protoStruct(), nil
}
func (s *AccountService) DeleteAccountUser(ctx context.Context, req *pb.DeleteAccountUserReq) (*emptypb.Empty, error) {
	err := s.accountUserUsecase.DeleteAccountUser(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
func (s *AccountService) GetAccountUser(ctx context.Context, req *pb.GetAccountUserReq) (*pb.AccountUser, error) {
	p, err := s.accountUserUsecase.GetAccountUser(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return bizAccountUser(*p).protoStruct(), nil
}
func (s *AccountService) ListAccountUser(ctx context.Context, req *pb.ListAccountUserReq) (*pb.ListAccountUserResp, error) {
	rv, err := s.accountUserUsecase.ListAccountUser(ctx, req)
	if err != nil {
		return nil, err
	}

	rs := make([]*pb.AccountUser, 0, len(rv))
	for _, x := range rv {
		rs = append(rs, bizAccountUser(*x).protoStruct())
	}

	return &pb.ListAccountUserResp{
		List: rs,
	}, nil
}
func (s *AccountService) PageListAccountUser(ctx context.Context, req *pb.PageListAccountUserReq) (*pb.PageListAccountUserResp, error) {
	rv, total, err := s.accountUserUsecase.PageListAccountUser(ctx, req)
	if err != nil {
		return nil, err
	}

	rs := make([]*pb.AccountUser, 0, len(rv))
	for _, x := range rv {
		rs = append(rs, bizAccountUser(*x).protoStruct())
	}

	return &pb.PageListAccountUserResp{
		List:     rs,
		Total:    total,
		Page:     req.PageNum,
		PageSize: req.PageSize,
	}, nil
}
