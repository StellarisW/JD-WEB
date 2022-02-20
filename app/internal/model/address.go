package model

import (
	g "main/app/global"
)

type Address struct {
	g.Model
	Uid            int    `json:"uid" form:"uid"`
	Phone          string `json:"phone" form:"phone"`
	Name           string `json:"name" form:"name"`
	Address        string `json:"address" form:"address"`
	Zipcode        string `json:"zipcode" form:"zipcode"`
	DefaultAddress int    `json:"default_address" form:"default_address" db:"default_address"`
}

func (Address) TableName() string {
	return "address"
}
