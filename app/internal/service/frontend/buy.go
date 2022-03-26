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
	g.DB.Where("uid=? AND default_address=1", userId).Find(&address)
	return address
}

func (s *sBuy) GenerateSign() string {
	orderSign := utils.EncryptBySHA256(utils.GetRandomNum())
	return orderSign
}

func (s *sBuy) GenerateOrderInfo(userId int, address model.Address, orderList []model.Cart, allPrice float64) string {
	orderId := utils.GenerateOrderId()
	order := model.Order{
		OrderId:     utils.GenerateOrderId(),
		Uid:         userId,
		AllPrice:    allPrice,
		Phone:       address.Phone,
		Name:        address.Name,
		Address:     address.Address,
		Zipcode:     address.Zipcode,
		PayStatus:   0,
		PayType:     0,
		OrderStatus: 0,
	}
	g.DB.Create(&order)
	for i := 0; i < len(orderList); i++ {
		orderItem := model.OrderItem{
			OrderId:        orderId,
			Uid:            userId,
			ProductTitle:   orderList[i].Title,
			ProductId:      orderList[i].Id,
			ProductImg:     orderList[i].ProductImg,
			ProductPrice:   orderList[i].Price,
			ProductNum:     orderList[i].Num,
			ProductVersion: orderList[i].ProductVersion,
			ProductColor:   orderList[i].ProductColor,
		}
		g.DB.Create(&orderItem)
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
	var order model.Order
	orderId := strings.Split(temp, "_")[1]
	g.DB.Where("order_id=?", orderId).Find(&order)
	order.PayStatus = 1
	order.OrderStatus = 1
	g.DB.Save(&order)
}

func (s *sBuy) GetOrderInfo(orderId string) model.Order {
	var orderInfo model.Order
	g.DB.Where("order_id=?", orderId).Find(&orderInfo)
	return orderInfo
}

func (s *sBuy) GetOrderItem(orderId string) []model.OrderItem {
	var orderItem []model.OrderItem
	_ = g.DB.Select(&orderItem, "select * from order_item where order_id=?", orderId)
	return orderItem
}
