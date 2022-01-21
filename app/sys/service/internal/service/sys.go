package service

import (
	"github.com/go-kratos/kratos/v2/log"

	pb "github.com/lalifeier/vvgo-mall/api/sys/service/v1"
	"github.com/lalifeier/vvgo-mall/app/sys/service/internal/biz"
)

type SysService struct {
	pb.UnimplementedSysServer

	log         *log.Helper
	dictUsecase *biz.DictUsecase
}

func NewSysService(dictUsecase *biz.DictUsecase, logger log.Logger) *SysService {
	return &SysService{
		log:         log.NewHelper(log.With(logger, "module", "sys-service/service")),
		dictUsecase: dictUsecase,
	}
}
