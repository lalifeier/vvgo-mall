package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/lalifeier/vvgo-mall/app/sms/interface/internal/biz"
	pb "github.com/lalifeier/vvgo-mall/gen/api/go/sms/interface/v1"
)

type SmsService struct {
	pb.UnimplementedSmsServer
	uc  *biz.SmsUseCase
	log *log.Helper
}

func NewSmsService(uc *biz.SmsUseCase, logger log.Logger) *SmsService {
	return &SmsService{
		uc: uc, log: log.NewHelper(logger),
	}
}

func (s *SmsService) SendSmsCode(ctx context.Context, req *pb.SendSmsCodeRequest) (*pb.SendSmsCodeReply, error) {
	err := s.uc.SendCode(ctx, req.Mobile)
	if err != nil {
		return nil, err
	}
	return &pb.SendSmsCodeReply{}, nil
}

func (s *SmsService) VerifySmsCode(ctx context.Context, req *pb.VerifySmsCodeRequest) (*pb.VerifySmsCodeReply, error) {
	err := s.uc.VerifyCode(ctx, req.Mobile, req.Code)
	if err != nil {
		return nil, err
	}
	return &pb.VerifySmsCodeReply{}, nil
}
