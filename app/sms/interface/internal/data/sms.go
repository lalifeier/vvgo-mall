package data

import (
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/lalifeier/vgo/app/sms/interface/internal/biz"
)

var _ biz.SmsRepo = (*smsRepo)(nil)

type smsRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewSmsRepo(data *Data, logger log.Logger) biz.SmsRepo {
	return &smsRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r smsRepo) SendSmsCode(mobile, code string) error {
	return r.data.smsClient.SendCode(mobile, code)
}

func (r smsRepo) Set(key string, value interface{}, expiration time.Duration) error {
	if _, err := r.data.rdb.Set(key, value, expiration).Result(); err != nil {
		return err
	}
	return nil
}

func (r smsRepo) Get(key string) (string, error) {
	return r.data.rdb.Get(key).Result()
}

func (r smsRepo) Del(key string) (int64, error) {
	return r.data.rdb.Del(key).Result()
}
