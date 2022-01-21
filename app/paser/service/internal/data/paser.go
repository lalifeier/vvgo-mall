package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/lalifeier/vvgo-mall/app/paser/service/internal/biz"
)

type paserRepo struct {
	data *Data
	log  *log.Helper
}

func NewPaserRepo(data *Data, logger log.Logger) biz.PaserRepo {
	return &paserRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
