package captcha

import (
	"github.com/mojocn/base64Captcha"
)

var captchaStore = base64Captcha.DefaultMemStore

type CaptchaInfo struct {
	CaptchaId string
	ImgBytes  string
}

func GetCaptcha() (*CaptchaInfo, error) {
	driver := base64Captcha.NewDriverDigit(80, 250, 5, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, captchaStore)
	id, b64s, err := cp.Generate()
	if err != nil {
		return nil, err
	}

	return &CaptchaInfo{
		CaptchaId: id,
		ImgBytes:  b64s,
	}, nil
}

func VerifyCaptcha(captchaId string, captchaCode string) bool {
	return captchaStore.Verify(captchaId, captchaCode, true)
}
