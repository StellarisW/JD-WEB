package frontend

import (
	g "main/app/global"
	"main/app/internal/model"
	"main/utils"
)

type sAuth struct{}

var insAuth = sAuth{}

func (s *sAuth) Login(user *model.User) (*model.User, error) {
	user.Password = utils.Md5(user.Password)
	g.Logger.Debugf("%v\n", user.Password)
	g.DB.Where("phone=? AND password=?", user.Phone, user.Password).Find(&user)
	return user, nil
}
