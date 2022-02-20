package model

import g "main/app/global"

type Role struct {
	g.Model
	Title       string
	Description string
	Status      int
}

type RoleAuth struct {
	AuthId int `db:"auth_id"`
	RoleId int `db:"role_id"`
}

func (Role) TableName() string {
	return "role"
}

func (RoleAuth) TableName() string {
	return "role_auth"
}
