package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/lalifeier/vvgo-mall/app/taoke/interface/internal/biz"
	pb "github.com/lalifeier/vvgo-mall/gen/api/go/taoke/interface/v1"
)

type TaoKeService struct {
	pb.UnimplementedTaoKeServer

	log *log.Helper

	taoBaoUseCase *biz.TaoBaoUseCase
}

func NewTaoKeService(logger log.Logger, taoBaoUseCase *biz.TaoBaoUseCase) *TaoKeService {
	return &TaoKeService{
		log: log.NewHelper(log.With(logger, "module", "taoke-interface/service")),

		taoBaoUseCase: taoBaoUseCase,
	}
}
