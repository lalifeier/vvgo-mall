package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/lalifeier/vvgo-mall/app/account/service/internal/biz"
	pb "github.com/lalifeier/vvgo-mall/gen/api/go/account/service/v1"
	v1 "github.com/lalifeier/vvgo-mall/gen/api/go/account/service/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

type AccountUserService struct {
	v1.UnimplementedAccountUserServiceServer
	log *log.Helper

	accountUserUseCase *biz.AccountUserUseCase
}

func NewAccountUserService(logger log.Logger, accountUserUseCase *biz.AccountUserUseCase) *AccountUserService {
	return &AccountUserService{
		log:                log.NewHelper(log.With(logger, "module", "account-service/service")),
		accountUserUseCase: accountUserUseCase,
	}
}

type bizAccountUser biz.AccountUser

func (b bizAccountUser) protoStruct() *pb.AccountUser {
	return &pb.AccountUser{
		Id: b.Id,
	}
}

func (s *AccountService) CreateAccountUser(ctx context.Context, req *pb.AccountUser) (*pb.AccountUser, error) {
	p, err := s.accountUserUseCase.CreateAccountUser(ctx, &biz.AccountUser{})
	if err != nil {
		return nil, err
	}
	return bizAccountUser(*p).protoStruct(), nil
}
func (s *AccountService) UpdateAccountUser(ctx context.Context, req *pb.AccountUser) (*pb.AccountUser, error) {
	p, err := s.accountUserUseCase.UpdateAccountUser(ctx, &biz.AccountUser{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	return bizAccountUser(*p).protoStruct(), nil
}
func (s *AccountService) DeleteAccountUser(ctx context.Context, req *pb.DeleteAccountUserReq) (*emptypb.Empty, error) {
	err := s.accountUserUseCase.DeleteAccountUser(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
func (s *AccountService) GetAccountUser(ctx context.Context, req *pb.GetAccountUserReq) (*pb.AccountUser, error) {
	p, err := s.accountUserUseCase.GetAccountUser(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return bizAccountUser(*p).protoStruct(), nil
}
func (s *AccountService) ListAccountUser(ctx context.Context, req *pb.ListAccountUserReq) (*pb.ListAccountUserResp, error) {
	rv, err := s.accountUserUseCase.ListAccountUser(ctx, req)
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
	rv, total, err := s.accountUserUseCase.PageListAccountUser(ctx, req)
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
