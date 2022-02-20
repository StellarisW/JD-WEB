package service

import (
	"github.com/google/wire"
)

var LoginSet = wire.NewSet(wire.Struct(new(sLogin), "*"))

type sLogin struct{}

var insLogin = sLogin{}

func Login() *sLogin {
	return &insLogin
}

//func (a *sLogin) GetCaptcha(ctx context.Context, length int) (*schema.LoginCaptcha, error) {
//	captchaID := captcha.NewLen(length)
//	item := &schema.LoginCaptcha{
//		CaptchaID: captchaID,
//	}
//	return item, nil
//}
