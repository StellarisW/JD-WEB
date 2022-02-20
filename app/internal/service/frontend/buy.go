package frontend

import (
	g "main/app/global"
	"main/app/internal/model"
	"main/utils"
	"strings"
)

type sBuy struct{}

var insBuy = sBuy{}

func (s *sBuy) Cal(cartList []model.Cart) (bool, []model.Cart, float64) {
	var orderList []model.Cart
	var allPrice float64
	//2.执行计算总价
	for i := 0; i < len(cartList); i++ {
		if cartList[i].Checked {
			allPrice += cartList[i].Price * float64(cartList[i].Num)
			orderList = append(orderList, cartList[i])
		}
	}
	if len(orderList) == 0 {
		return false, nil, 0
	}
	return true, orderList, allPrice
}

func (s *sBuy) GetAddress(userId int) []model.Address {
	var addressList []model.Address
	_ = g.DB.Select(&addressList, "select * from address where uid=? order by default_address desc", userId)
	return addressList
}

func (s *sBuy) GetDefaultAddress(userId int) model.Address {
	var address model.Address
	_ = g.DB.Get(&address, "select * from address where uid=? and default_address=1", userId)
	return address
}

func (s *sBuy) GenerateSign() string {
	orderSign := utils.EncryptBySHA256(utils.GetRandomNum())
	return orderSign
}

func (s *sBuy) GenerateOrderInfo(userId int, address model.Address, orderList []model.Cart, allPrice float64) string {
	orderId := utils.GenerateOrderId()
	_, _ = g.DB.Exec("insert into `order`(order_id, uid, all_price, phone, name, address, zipcode, pay_status, pay_type, order_status) values (?,?,?,?,?,?,?,?,?,?)",
		orderId, userId, allPrice, address.Phone, address.Name, address.Address, address.Zipcode, 0, 0, 0)
	var order model.Order
	_ = g.DB.Get(&order, "select * from `order` where order_id=?", orderId)
	for i := 0; i < len(orderList); i++ {
		_, _ = g.DB.Exec("insert into `order_item`(order_id, uid, product_title, product_id, product_img, product_price, product_num, product_version, product_color) values (?,?,?,?,?,?,?,?,?)",
			orderId, userId, orderList[i].Title, orderList[i].Id, orderList[i].ProductImg, orderList[i].Price, orderList[i].Num, orderList[i].ProductVersion, orderList[i].ProductColor)
	}
	return orderId
}

func (s *sBuy) UpdateCart(PCartList *[]model.Cart) {
	var noSelectedCartList []model.Cart
	cartList := *PCartList
	for i := 0; i < len(cartList); i++ {
		if !cartList[i].Checked {
			noSelectedCartList = append(noSelectedCartList, cartList[i])
		}
	}
	*PCartList = noSelectedCartList
}

func (s *sBuy) UpdateOrder(temp string) {
	orderId := strings.Split(temp, "_")[1]
	_, _ = g.DB.Exec("update `order` set pay_status=1,order_status=1 where order_id=?", orderId)
}

func (s *sBuy) GetOrderInfo(orderId string) model.Order {
	var orderInfo model.Order
	_ = g.DB.Get(&orderInfo, "select * from `order` where order_id=?", orderId)
	return orderInfo
}

func (s *sBuy) GetOrderItem(orderId string) []model.OrderItem {
	var orderItem []model.OrderItem
	_ = g.DB.Select(&orderItem, "select * from order_item where order_id=?", orderId)
	return orderItem
}
