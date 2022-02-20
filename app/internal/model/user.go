package model

import g "main/app/global"

type User struct {
	g.Model
	Phone    string `json:"phone" form:"phone"`
	Password string `json:"password" form:"password"`
	LastIp   string `json:"last_ip" form:"last_ip" db:"last_ip"`
	Email    string `json:"email" form:"email"`
	Status   int    `json:"status" form:"status"`
}

type UserSms struct {
	g.Model
	Ip        string
	Phone     string
	SendCount int    `json:"send_count" form:"send_count" db:"send_count"`
	AddDay    string `json:"add_day" form:"add_day" db:"add_day"`
	Sign      string
}

func (User) TableName() string {
	return "user"
}

func (UserSms) TableName() string {
	return "user_sms"
}

func (u User) Verify() bool {
	var user User
	_ = g.DB.Get(&user, "select * from user where phone=?", u.Phone)
	if user.Id == 0 {
		return false
	}
	return true
}
