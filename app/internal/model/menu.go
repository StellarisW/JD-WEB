package model

import g "main/app/global"

type Menu struct {
	g.Model
	Title       string
	Link        string
	Position    int
	IsOpennew   int    `db:"is_opennew"`
	Relation    string // 关联商品
	Sort        int
	Status      int
	ProductItem []Product // `gorm:"-"`
}

func (Menu) TableName() string {
	return "menu"
}
