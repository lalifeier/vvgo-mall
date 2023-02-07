package service

import (
	"context"

	// "github.com/go-kratos/kratos/v2/transport"
	pb "github.com/lalifeier/vvgo-mall/gen/api/go/auth/service/v1"
)

func (s *AuthService) GetCaptcha(ctx context.Context, req *pb.GetCaptchaReq) (*pb.GetCaptchaResp, error) {

	// tr, _ := transport.FromServerContext(ctx)
	// tr.RequestHeader()
	captchaInfo, err := s.captchaUseCase.GetCaptcha(ctx)
	if err != nil {
		return nil, err
	}
	return &pb.GetCaptchaResp{
		CaptchaId: captchaInfo.CaptchaId,
		ImgBytes:  captchaInfo.ImgBytes,
	}, nil
}

func (s *AuthService) VerifyCaptcha(ctx context.Context, req *pb.VerifyCaptchaReq) (*pb.VerifyCaptchaResp, error) {
	err := s.captchaUseCase.VerifyCaptcha(ctx, req.CaptchaId, req.CaptchaCode)
	if err != nil {
		return nil, err
	}
	return &pb.VerifyCaptchaResp{}, nil
}
