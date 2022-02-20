package model

import g "main/app/global"

type Cart struct {
	g.Model
	Title          string
	Price          float64
	ProductVersion string `db:"product_version"`
	Num            int
	ProductGift    string `db:"product_gift"`
	ProductFitting string `db:"product_fitting"`
	ProductColor   string `db:"product_color"`
	ProductImg     string `db:"product_img"`
	ProductAttr    string `db:"product_attr"`
	Checked        bool   //`gorm:"-"` // 忽略本字段
}

func (Cart) TableName() string {
	return "cart"
}

// CartHasData 判断购物车里面有没有当前数据
func CartHasData(cartList []Cart, currentData Cart) bool {
	for i := 0; i < len(cartList); i++ {
		if cartList[i].Id == currentData.Id &&
			cartList[i].ProductColor == currentData.ProductColor &&
			cartList[i].ProductAttr == currentData.ProductAttr {
			return true
		}
	}
	return false
}
