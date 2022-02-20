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
	err := g.DB.Get(user,
		"select * from user where phone=? and password=?", user.Phone, user.Password)
	if err != nil {
		g.Logger.Debugf("%v\n", err)
		return nil, err
	} else {
		return user, nil
	}
}
