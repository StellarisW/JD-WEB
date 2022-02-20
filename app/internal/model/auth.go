package model

import (
	g "main/app/global"
)

type Auth struct {
	g.Model
	ModuleName  string `json:"module_name" form:"module_name" db:"module_name"` //模块名称
	ActionName  string `json:"action_name" form:"action_name" db:"action_name"` //操作名称
	Type        int    //节点类型 :  1、表示模块    2、表示菜单     3、操作
	Url         string //路由跳转地址
	ModuleId    int    `json:"module_id" form:"module_id" db:"module_id"` //此module_id和当前模型的_id关联      module_id= 0 表示模块
	Sort        int
	Description string
	Status      int
	AuthItem    []Auth //`gorm:"foreignkey:ModuleId;association_foreignkey:Id"`
	Checked     bool   //`gorm:"-"` // 忽略本字段
}

func (Auth) TableName() string {
	return "auth"
}
