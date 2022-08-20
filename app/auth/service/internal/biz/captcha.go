package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/lalifeier/vvgo-mall/pkg/captcha"
)

var (
	ErrCaptcha = errors.Unauthorized("UNAUTHORIZED", "验证码错误")
)

type CaptchaRepo interface {
	// SendSmsCode(ctx context.Context, phone string) error
}

type CaptchaUseCase struct {
	log *log.Helper
	// repo CaptchaRepo
}

func NewCaptchaUseCase(logger log.Logger) *CaptchaUseCase {
	return &CaptchaUseCase{log: log.NewHelper(logger)}
}

func (us *CaptchaUseCase) GetCaptcha(ctx context.Context) (captchaInfo *captcha.CaptchaInfo, err error) {
	captchaInfo, err = captcha.GetCaptcha()
	return
}

func (us *CaptchaUseCase) VerifyCaptcha(ctx context.Context, captchaId string, captchaCode string) (err error) {
	ok := captcha.VerifyCaptcha(captchaId, captchaCode)
	if !ok {
		return ErrCaptcha
	}
	return
}
