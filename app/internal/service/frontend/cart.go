package frontend

import (
	g "main/app/global"
	"main/app/internal/model"
)

type sCart struct{}

var insCart = sCart{}

func (s *sCart) AddPruduct(productId, colorId int) (currentData model.Cart, product model.Product) {
	var productColor model.ProductColor

	_ = g.DB.Get(&product, "select * from product where id=?", productId)
	_ = g.DB.Get(&productColor, "select * from product_color where id=?", colorId)

	// 获取增加购物车的商品数据
	currentData = model.Cart{
		Title:          product.Title,
		Price:          product.Price,
		ProductVersion: product.ProductVersion,
		Num:            1,
		ProductColor:   productColor.ColorName,
		ProductImg:     product.ProductImg,
		ProductGift:    product.ProductGift, //赠品
		ProductAttr:    "",                  //根据自己的需求拓展
		Checked:        true,                //默认选中
	}
	currentData.Id = productId
	return currentData, product
}

func (s *sCart) UpdateCookie(PCartList *[]model.Cart, currentData model.Cart) {
	cartList := *PCartList
	// 判断购物车有没有数据（cookie）
	if len(cartList) > 0 { //购物车有数据
		// 判断购物车有没有当前数据
		if model.CartHasData(cartList, currentData) {
			// 已经有该商品则数量+1
			for i := 0; i < len(cartList); i++ {
				if cartList[i].Id == currentData.Id && cartList[i].ProductColor == currentData.ProductColor && cartList[i].ProductAttr == currentData.ProductAttr {
					cartList[i].Num = cartList[i].Num + 1
				}
			}
		} else {
			// 没有直接在cookie追加数据
			cartList = append(cartList, currentData)
		}

	} else {
		// 如果购物车没有任何数据，直接把当前数据写入cookie
		cartList = append(cartList, currentData)
	}
	*PCartList = cartList
}

func (s *sCart) CalAllPrice(cartList []model.Cart) (allPrice float64) {
	for i := 0; i < len(cartList); i++ {
		if cartList[i].Checked {
			allPrice += cartList[i].Price * float64(cartList[i].Num)
		}
	}
	return allPrice
}

func (s *sCart) IncItem(PCartList *[]model.Cart, productId int, productColor string, productAttr string) (bool, int, float64, float64) {
	var flag bool
	var num int
	var currentAllPrice float64
	var allPrice float64
	cartList := *PCartList
	for i := 0; i < len(cartList); i++ {
		if cartList[i].Id == productId && cartList[i].ProductColor == productColor && cartList[i].ProductAttr == productAttr {
			cartList[i].Num = cartList[i].Num + 1
			flag = true
			num = cartList[i].Num
			currentAllPrice = cartList[i].Price * float64(cartList[i].Num)
		}
		if cartList[i].Checked {
			allPrice += cartList[i].Price * float64(cartList[i].Num)
		}
	}
	*PCartList = cartList
	return flag, num, currentAllPrice, allPrice
}

func (s *sCart) DecItem(PCartList *[]model.Cart, productId int, productColor string, productAttr string) (bool, int, float64, float64) {
	var flag bool
	var num int
	var currentAllPrice float64
	var allPrice float64
	cartList := *PCartList
	for i := 0; i < len(cartList); i++ {
		if cartList[i].Id == productId && cartList[i].ProductColor == productColor && cartList[i].ProductAttr == productAttr {
			if cartList[i].Num > 1 {
				cartList[i].Num = cartList[i].Num - 1
			}
			flag = true
			num = cartList[i].Num
			currentAllPrice = cartList[i].Price * float64(cartList[i].Num)
		}
		if cartList[i].Checked {
			allPrice += cartList[i].Price * float64(cartList[i].Num)
		}
	}
	*PCartList = cartList
	return flag, num, currentAllPrice, allPrice
}

func (s *sCart) DelItem(PCartList *[]model.Cart, productId int, productColor string, productAttr string) {
	cartList := *PCartList

	for i := 0; i < len(cartList); i++ {
		if cartList[i].Id == productId && cartList[i].ProductColor == productColor && cartList[i].ProductAttr == productAttr {
			//执行删除
			cartList = append(cartList[:i], cartList[(i+1):]...)
		}
	}

	*PCartList = cartList
}

func (s *sCart) SelectOneItem(PCartList *[]model.Cart, productId int, productColor string, productAttr string) (bool, float64) {
	var flag bool
	var allPrice float64
	cartList := *PCartList
	for i := 0; i < len(cartList); i++ {
		if cartList[i].Id == productId && cartList[i].ProductColor == productColor && cartList[i].ProductAttr == productAttr {
			cartList[i].Checked = !cartList[i].Checked
			flag = true
		}
		if cartList[i].Checked {
			allPrice += cartList[i].Price * float64(cartList[i].Num)
		}
	}
	*PCartList = cartList
	return flag, allPrice
}

func (s *sCart) SelectAllItem(PCartList *[]model.Cart, flag int) float64 {
	var allPrice float64
	cartList := *PCartList
	for i := 0; i < len(cartList); i++ {
		if flag == 1 {
			cartList[i].Checked = true
		} else {
			cartList[i].Checked = false
		}
		//计算总价
		if cartList[i].Checked {
			allPrice += cartList[i].Price * float64(cartList[i].Num)
		}
	}
	*PCartList = cartList
	return allPrice
}
