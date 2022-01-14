package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	pb "github.com/lalifeier/vvgo/api/paser/service/v1"
	"github.com/lalifeier/vvgo/app/paser/service/internal/biz"
)

type PaserService struct {
	pb.UnimplementedPaserServer

	log          *log.Helper
	paserUsecase *biz.PaserUsecase
}

func NewPaserService(paserUsecase *biz.PaserUsecase, logger log.Logger) *PaserService {
	return &PaserService{
		log:          log.NewHelper(log.With(logger, "module", "ums-service/service")),
		paserUsecase: paserUsecase,
	}
}

func (s *PaserService) Paser(ctx context.Context, req *pb.PaserReq) (*pb.PaserReply, error) {
	return s.paserUsecase.Paser(req.Url)
}
