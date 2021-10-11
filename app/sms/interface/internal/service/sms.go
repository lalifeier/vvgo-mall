package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	pb "github.com/lalifeier/vgo/api/sms/interface/v1"
	"github.com/lalifeier/vgo/app/sms/interface/internal/biz"
)

type SmsService struct {
	pb.UnimplementedSmsServer
	uc  *biz.SmsUsecase
	log *log.Helper
}

func NewSmsService(uc *biz.SmsUsecase, logger log.Logger) *SmsService {
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
