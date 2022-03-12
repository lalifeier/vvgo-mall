package biz

import (
	"context"

	"github.com/golang-jwt/jwt"
	pb "github.com/lalifeier/vvgo-mall/api/shop/admin/v1"
	"github.com/lalifeier/vvgo-mall/app/shop/admin/internal/conf"
)

type AuthUsecase struct {
	key string
}

func NewAuthUsecase(conf *conf.Auth) *AuthUsecase {
	return &AuthUsecase{
		key: conf.ApiKey,
	}
}

func (uc *AuthUsecase) Register(ctx context.Context, req *pb.RegisterReq) (*pb.RegisterResp, error) {
	return &pb.RegisterResp{}, nil
}
func (uc *AuthUsecase) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginResp, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{})
	signedString, err := claims.SignedString([]byte(uc.key))
	if err != nil {
		return nil, err
	}
	return &pb.LoginResp{Token: signedString}, nil
}
