package model

import g "main/app/global"

type Administrator struct {
	g.Model
	Username string
	Password string
	Mobile   string
	Email    string
	Status   int
	RoleId   int //`gorm:"roleid"`
	IsSuper  int
	Role     Role //`gorm:"foreignkey:Id;association_foreignkey:RoleId"`
}

func (Administrator) TableName() string {
	return "administrator"
}
