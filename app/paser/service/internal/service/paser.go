package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/lalifeier/vvgo-mall/app/paser/service/internal/biz"
	pb "github.com/lalifeier/vvgo-mall/gen/api/paser/service/v1"
)

type PaserService struct {
	pb.UnimplementedPaserServer

	log          *log.Helper
	paserUseCase *biz.PaserUseCase
}

func NewPaserService(paserUseCase *biz.PaserUseCase, logger log.Logger) *PaserService {
	return &PaserService{
		log:          log.NewHelper(log.With(logger, "module", "ums-service/service")),
		paserUseCase: paserUseCase,
	}
}

func (s *PaserService) Paser(ctx context.Context, req *pb.PaserReq) (*pb.PaserReply, error) {
	return s.paserUseCase.Paser(req.Url)
}
