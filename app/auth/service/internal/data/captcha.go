package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/lalifeier/vvgo-mall/app/auth/service/internal/biz"
)

type captchaRepo struct {
	data *Data
	log  *log.Helper
}

func NewCaptchaRepo(data *Data, logger log.Logger) biz.CaptchaRepo {
	return &captchaRepo{data: data, log: log.NewHelper(logger)}
}

func (r *captchaRepo) GetCaptcha(ctx context.Context, uid string) (img []byte, err error) {
	return nil, nil
}
