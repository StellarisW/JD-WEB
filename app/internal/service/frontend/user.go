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
	var count int64
	g.DB.Where(where, userId).Table("order").Count(&count)
	var order []model.Order
	g.DB.Where(where, userId).Offset((page - 1) * pageSize).Limit(pageSize).Preload("OrderItem").
		Order("create_time desc").Find(&order)
	g.Logger.Debugf("%v\n", order)
	return order, math.Ceil(float64(count) / float64(pageSize)), page
}

func (s *sUser) GetOrderInfo(id, userId int) model.Order {
	var order model.Order
	g.DB.Where("id=? AND uid=?", id, userId).Preload("OrderItem").Find(&order)
	return order
}

func (s *sUser) GetCollectProduct(userId, page int) ([]model.Product, float64, int) {
	pageSize := 2
	var count int64
	g.DB.Where("user_id=?").Table("product_collect").Count(&count)
	var productIds []int
	var productList []model.Product
	var product model.Product
	g.DB.Where("user_id=?", userId).Offset((page - 1) * pageSize).Limit(pageSize).Order("create_time desc").Find(&productIds)
	for _, id := range productIds {
		g.DB.Where("id=?", id).Find(&product)
		productList = append(productList, product)
	}
	return productList, math.Ceil(float64(count) / float64(pageSize)), page
}

func (s *sUser) IsAddressLimit(userId int) bool {
	var count int64
	g.DB.Where("uid=?", userId).Count(&count)
	if count > 10 {
		return true
	}
	return false
}

func (s *sUser) GetOneAddress(addressId int) model.Address {
	var address model.Address
	g.DB.Where("id=?", addressId).Find(&address)
	return address
}

func (s *sUser) AddAddress(userId int, address model.Address) []model.Address {
	g.DB.Table("address").Where("uid=?", userId).Updates(map[string]interface{}{"default_address": 0})
	addressResult := model.Address{
		Uid:            address.Uid,
		Name:           address.Phone,
		Phone:          address.Phone,
		Address:        address.Address,
		Zipcode:        address.Zipcode,
		DefaultAddress: address.DefaultAddress,
	}
	g.DB.Create(&addressResult)
	var allAddressResult []model.Address
	g.DB.Where("uid=?", userId).Find(&allAddressResult)
	return allAddressResult
}

func (s *sUser) EditAddress(userId int, address model.Address) []model.Address {
	g.DB.Table("address").Where("uid=?", userId).Updates(map[string]interface{}{"default_address": 0})
	addressModel := model.Address{}
	g.DB.Where("id=?", address.Id).Find(&addressModel)
	addressModel.Name = address.Name
	addressModel.Phone = address.Phone
	addressModel.Zipcode = address.Zipcode
	addressModel.Address = address.Address
	addressModel.DefaultAddress = address.DefaultAddress
	g.DB.Save(&addressModel)
	var allAddressResult []model.Address
	g.DB.Where("uid=?", userId).Find(&allAddressResult)
	return allAddressResult
}

func (s *sUser) ChangeToDefaultAddress(userId, addressId int) {
	g.DB.Table("address").Where("uid=?", userId).Updates(map[string]interface{}{"default_address": 0})
	g.DB.Table("address").Where("id=?", addressId).Updates(map[string]interface{}{"default_address": 1})
}
