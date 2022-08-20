package biz

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"

	sms "github.com/lalifeier/vvgo-mall/pkg/sms/aliyun"
)

var (
	codePrefix = "sms_code_"
	codeExpire = 5 * time.Minute

	ErrSmsCode = errors.InternalServer("ERROR", "Sms Code ERROR")
)

type Sms struct {
	Mobile string
}

type SmsRepo interface {
	SendSmsCode(mobile, code string) error

	Set(key string, value interface{}, expiration time.Duration) error
	Get(key string) (string, error)
	Del(key string) (int64, error)
}

type SmsUseCase struct {
	repo SmsRepo
	log  *log.Helper
}

func NewSmsUseCase(repo SmsRepo, logger log.Logger) *SmsUseCase {
	return &SmsUseCase{repo: repo, log: log.NewHelper(logger)}
}

func generateCodeKey(mobile string) string {
	return codePrefix + mobile
}

func (uc *SmsUseCase) SendCode(ctx context.Context, mobile string) error {
	code := sms.Code()
	err := uc.repo.SendSmsCode(mobile, code)

	if err != nil {
		return errors.InternalServer(err.Error(), err.Error())
	}

	key := generateCodeKey(mobile)

	return uc.repo.Set(key, code, codeExpire)
}

func (uc *SmsUseCase) VerifyCode(ctx context.Context, mobile, code string) error {
	key := generateCodeKey(mobile)
	val, err := uc.repo.Get(key)
	if err != nil {
		return errors.InternalServer(err.Error(), err.Error())
	}

	if val != code {
		return ErrSmsCode
	}

	uc.repo.Del(key)

	return nil
}
