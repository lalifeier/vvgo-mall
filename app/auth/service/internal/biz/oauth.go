package biz

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-oauth2/oauth2/v4/server"
)

type OauthUseCase struct {
	log    *log.Helper
	server *server.Server
}

func NewOauthUseCase(logger log.Logger, server *server.Server) *OauthUseCase {
	return &OauthUseCase{log: log.NewHelper(logger), server: server}
}
