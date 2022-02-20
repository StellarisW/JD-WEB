package frontend

import (
	g "main/app/global"
	"main/app/internal/model"
	"main/utils"
	"math"
	"time"
)

type sUser struct{}

var insUser = sUser{}

func (s *sUser) GetUserInfo(claims *utils.BaseClaims) model.User {
	user := model.User{
		Phone:  claims.Phone,
		LastIp: claims.LastIp,
		Email:  claims.Email,
		Status: claims.Status,
	}
	user.Id = claims.Id
	user.UpdateTime = claims.UpdateTime
	user.CreateTime = claims.CreateTime
	return user
}

func (s *sUser) GetTimeSlot() string {
	var msg string
	time := time.Now().Hour()
	if time >= 12 && time <= 18 {
		msg = "尊敬的用户下午好"
	} else if time >= 6 && time < 12 {
		msg = "尊敬的用户上午好！"
	} else {
		msg = "深夜了，注意休息哦！"
	}
	return msg
}

func (s *sUser) GetOrderNum(userId int) (int, int) {
	var order []model.Order
	var waitPay, waitRec int
	_ = g.DB.Select(&order, "select * from `order` where uid=?", userId)
	for i := 0; i < len(order); i++ {
		if order[i].PayStatus == 0 {
			waitPay++
		}
		if order[i].OrderStatus >= 2 && order[i].OrderStatus < 4 {
			waitRec++
		}
	}
	return waitPay, waitRec
}

func (s *sUser) GetSearchResult(keywords string, sqlStr *string) {
	where := *sqlStr
	var orderItem []model.OrderItem
	g.DB.Select(&orderItem, "select * from order_item where product_title like ?", "%"+keywords+"%")
	var str string
	for i := 0; i < len(orderItem); i++ {
		if i == 0 {
			str += orderItem[i].OrderId
		} else {
			str += "," + orderItem[i].OrderId
		}
	}
	where += "and id in (" + str + ")"
}

func (s *sUser) GetOrderList(userId, page int, where string) ([]model.Order, float64, int) {
	pageSize := 2
	var count int
	_ = g.DB.QueryRow("select count(*) from `order` where "+where, userId).Scan(&count)
	var order []model.Order
	var orderItem []model.OrderItem
	_ = g.DB.Select(&order, "select * from `order` where "+where+" order by create_time desc limit ? offset ?",
		userId, pageSize, (page-1)*pageSize)

	g.Logger.Debugf("%v\n%v\n", order, orderItem)
	for k, o := range order {
		_ = g.DB.Select(&orderItem, "select * from order_item where order_id=? order by create_time desc limit ? offset ?",
			o.OrderId, pageSize, (page-1)*pageSize)
		order[k].OrderItem = orderItem
	}
	g.Logger.Debugf("%v\n", order)
	return order, math.Ceil(float64(count) / float64(pageSize)), page
}

func (s *sUser) GetOrderInfo(id, userId int) model.Order {
	var order model.Order
	var orderItem []model.OrderItem
	_ = g.DB.Get(&order, "select * from `order` where id=? and uid=?", id, userId)
	_ = g.DB.Select(&orderItem, "select * from order_item where order_id=? and uid=?", order.OrderId, userId)
	order.OrderItem = orderItem
	return order
}

func (s *sUser) GetCollectProduct(userId, page int) ([]model.Product, float64, int) {
	pageSize := 2
	var count int
	_ = g.DB.QueryRow("select count(*) from product_collect where user_id=?", userId).Scan(&count)
	var productIds []int
	var productList []model.Product
	var product model.Product
	_ = g.DB.Select(&productIds, "select product_id from product_collect where user_id=? order by create_time desc limit ? offset ?",
		userId, pageSize, (page-1)*pageSize)
	for _, id := range productIds {
		_ = g.DB.Get(&product, "select * from product where id=?", id)
		productList = append(productList, product)
	}
	return productList, math.Ceil(float64(count) / float64(pageSize)), page
}

func (s *sUser) IsAddressLimit(userId int) bool {
	var count int
	_ = g.DB.QueryRow("select count(*) from address where uid=?", userId).Scan(&count)
	if count > 10 {
		return true
	}
	return false
}

func (s *sUser) GetOneAddress(addressId int) model.Address {
	var address model.Address
	_ = g.DB.Get(&address, "select * from address where id=?", addressId)
	return address
}

func (s *sUser) AddAddress(userId int, address model.Address) []model.Address {
	_, err := g.DB.Exec("update address set default_address=0 where uid=?", userId)
	g.Logger.Debugf("%v\n", err)
	_, err = g.DB.Exec("insert into address(uid, phone, name, zipcode, address,default_address)values (?,?,?,?,?,?)",
		address.Uid, address.Phone, address.Name, address.Zipcode, address.Address, address.DefaultAddress)
	g.Logger.Debugf("%v\n", err)

	var allAddressResult []model.Address
	_ = g.DB.Select(&allAddressResult, "select * from address where uid=?", userId)

	return allAddressResult
}

func (s *sUser) EditAddress(userId int, address model.Address) []model.Address {
	g.Logger.Debugf("%v\n", address)
	_, err := g.DB.Exec("update address set default_address=0 where uid=?", userId)
	g.Logger.Debugf("%v\n", err)
	_, err = g.DB.Exec("update address set phone=?,name=?,zipcode=?,address=?,default_address=? where id=?",
		address.Phone, address.Name, address.Zipcode, address.Address, address.DefaultAddress, address.Id)
	g.Logger.Debugf("%v\n", err)

	var allAddressResult []model.Address
	err = g.DB.Select(&allAddressResult, "select * from address where uid=? order by default_address desc",
		userId)
	g.Logger.Debugf("%v\n%v\n", err, allAddressResult)
	return allAddressResult
}

func (s *sUser) ChangeToDefaultAddress(userId, addressId int) {
	_, _ = g.DB.Exec("update address set default_address=0 where uid=?", userId)
	_, _ = g.DB.Exec("update address set default_address=1 where id=?", addressId)
}
