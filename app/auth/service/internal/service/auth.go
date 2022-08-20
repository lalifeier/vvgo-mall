package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-oauth2/oauth2/v4/server"

	pb "github.com/lalifeier/vvgo-mall/api/auth/service/v1"
	"github.com/lalifeier/vvgo-mall/app/auth/service/internal/biz"
)

type AuthService struct {
	pb.UnimplementedAuthServer
	log *log.Helper

	authUseCase    *biz.AuthUseCase
	captchaUseCase *biz.CaptchaUseCase

	server *server.Server
}

func NewAuthService(logger log.Logger, authUseCase *biz.AuthUseCase, captchaUseCase *biz.CaptchaUseCase, server *server.Server) *AuthService {
	return &AuthService{
		log:            log.NewHelper(log.With(logger, "module", "auth-service/service")),
		authUseCase:    authUseCase,
		captchaUseCase: captchaUseCase,
		server:         server,
	}
}

// func (s *AuthService) Authorize(ctx context.Context, req *pb.AuthorizeReq) (*pb.AuthorizeResp, error) {
// 	return &pb.AuthorizeResp{}, nil
// }
// func (s *AuthService) Token(ctx context.Context, req *pb.TokenReq) (*pb.TokenResp, error) {
// 	return &pb.TokenResp{}, nil
// }
// func (s *AuthService) Verify(ctx context.Context, req *pb.VerifyReq) (*pb.VerifyResp, error) {
// 	return &pb.VerifyResp{}, nil
// }
// func (s *AuthService) RefreshToken(ctx context.Context, req *pb.RefreshTokenReq) (*pb.RefreshTokenResp, error) {
// 	return &pb.RefreshTokenResp{}, nil
// }
// func (s *AuthService) Logout(ctx context.Context, req *pb.LogoutReq) (*pb.LogoutResp, error) {
// 	return &pb.LogoutResp{}, nil
// }
