package biz

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-oauth2/oauth2/v4/server"
)

type OauthUsecase struct {
	log    *log.Helper
	server *server.Server
}

func NewOauthUsecase(logger log.Logger, server *server.Server) *OauthUsecase {
	return &OauthUsecase{log: log.NewHelper(logger), server: server}
}
