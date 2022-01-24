package service

import (
	"context"

	pb "github.com/lalifeier/vvgo-mall/api/account/service/v1"
	"github.com/lalifeier/vvgo-mall/app/account/service/internal/biz"
	"google.golang.org/protobuf/types/known/emptypb"
)

type bizStaff biz.Staff

func (b bizStaff) protoStruct() *pb.Staff {
	return &pb.Staff{
		Id: b.Id,
	}
}

func (s *AccountService) CreateStaff(ctx context.Context, req *pb.Staff) (*pb.Staff, error) {
	p, err := s.staffUsecase.CreateStaff(ctx, &biz.Staff{})
	if err != nil {
		return nil, err
	}
	return bizStaff(*p).protoStruct(), nil
}
func (s *AccountService) UpdateStaff(ctx context.Context, req *pb.Staff) (*pb.Staff, error) {
	p, err := s.staffUsecase.UpdateStaff(ctx, &biz.Staff{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	return bizStaff(*p).protoStruct(), nil
}
func (s *AccountService) DeleteStaff(ctx context.Context, req *pb.DeleteStaffReq) (*emptypb.Empty, error) {
	err := s.staffUsecase.DeleteStaff(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
func (s *AccountService) GetStaff(ctx context.Context, req *pb.GetStaffReq) (*pb.Staff, error) {
	p, err := s.staffUsecase.GetStaff(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return bizStaff(*p).protoStruct(), nil
}
func (s *AccountService) ListStaff(ctx context.Context, req *pb.ListStaffReq) (*pb.ListStaffResp, error) {
	rv, err := s.staffUsecase.ListStaff(ctx, req)
	if err != nil {
		return nil, err
	}

	rs := make([]*pb.Staff, 0, len(rv))
	for _, x := range rv {
		rs = append(rs, bizStaff(*x).protoStruct())
	}

	return &pb.ListStaffResp{
		List: rs,
	}, nil
}
func (s *AccountService) PageListStaff(ctx context.Context, req *pb.PageListStaffReq) (*pb.PageListStaffResp, error) {
	rv, total, err := s.staffUsecase.PageListStaff(ctx, req)
	if err != nil {
		return nil, err
	}

	rs := make([]*pb.Staff, 0, len(rv))
	for _, x := range rv {
		rs = append(rs, bizStaff(*x).protoStruct())
	}

	return &pb.PageListStaffResp{
		List:     rs,
		Total:    total,
		Page:     req.PageNum,
		PageSize: req.PageSize,
	}, nil
}
