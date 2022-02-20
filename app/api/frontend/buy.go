package frontend

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	g "main/app/global"
	"main/app/internal/model"
	"main/app/internal/model/response"
	"main/app/internal/service"
	"main/utils"
	"main/utils/cookie"
	"net/http"
)

type BuyApi struct {
	BaseApi
}

var insBuy = BuyApi{}

// Index
// @Tags Buy
// @Summary 展示结算页面
// @Param cartList header string ture "cartList"
// @Success 200 "展示结算页面"
// @Router /buy/checkout [get]
func (a *BuyApi) Index(c *gin.Context) {
	//1.获取要结算的商品
	var cartList []model.Cart
	cookie.Get(c, "cartList", &cartList)

	//2.执行计算总价
	flag, orderList, allPrice := service.Frontend().Buy().Cal(cartList)
	if flag == false {
		c.Redirect(302, "/")
		return
	}

	//3.获取收货地址
	userId := utils.GetUserID(c)
	addressList := service.Frontend().Buy().GetAddress(userId)

	//4.防止重复提交订单 生成签名
	orderSign := service.Frontend().Buy().GenerateSign()
	g.Logger.Debugf("%v\n", orderSign)

	session := sessions.Default(c)
	session.Set("orderSign", orderSign)
	session.Save()

	c.HTML(http.StatusOK, "buy/index.tmpl",
		utils.MergeMaps(a.Init(c), gin.H{
			"orderList":   orderList,
			"allPrice":    allPrice,
			"addressList": addressList,
			"orderSign":   orderSign,
		}))
}

// Process
// 提交订单
// 1、获取收货地址信息
// 2、获取购买商品的信息
// 3、把订单信息放在订单表，把商品信息放在商品表
// 4、删除购物车里面的选中数据
// @Tags Buy
// @Summary 处理订单
// @Param orderSign formData string true "订单签名"
// @Param cartList header string ture "cartList"
// @Success 200
// @Header 200 {string} Location "/buy/confirm?id="
// @Router /buy/doOrder [post]
func (a *BuyApi) Process(c *gin.Context) {
	session := sessions.Default(c)
	//0、防止重复提交订单
	orderSign := c.PostForm("orderSign")
	sessionOrderSign := session.Get("orderSign")
	if sessionOrderSign != orderSign {
		g.Logger.Error("orderSigns do not match")
		c.Redirect(302, "/")
		return
	}
	session.Delete("orderSign")

	// 1、获取收货地址信息
	userId := utils.GetUserID(c)
	address := service.Frontend().Buy().GetDefaultAddress(userId)
	if address.Id == 0 {
		g.Logger.Error("get address failed")
		c.Redirect(302, "/")
		return
	}

	// 2、获取购买商品的信息   orderList就是要购买的商品信息
	var cartList []model.Cart
	cookie.Get(c, "cartList", &cartList)

	flag, orderList, allPrice := service.Frontend().Buy().Cal(cartList)
	if flag == false {
		g.Logger.Error("get orderList failed")
		c.Redirect(302, "/")
		return
	}

	// 3、把订单信息放在订单表，把商品信息放在商品表
	orderId := service.Frontend().Buy().GenerateOrderInfo(userId, address, orderList, allPrice)

	// 4、删除购物车里面的选中数据
	service.Frontend().Buy().UpdateCart(&cartList)
	cookie.Set(c, "cartList", cartList)
	//TODO
	c.Redirect(302, "/buy/confirm?id="+orderId)
}

// Confirm
// @Tags Buy
// @Summary 展示结算确认页面
// @Param orderId query string true "订单id"
// @Success 200 "展示页面"
// @Failure 400 "重定向到主页"
// @Header 400 {string} Location "/"
// @Router /buy/confirm [get]
func (a *BuyApi) Confirm(c *gin.Context) {
	orderId := c.Query("id")
	//获取用户信息
	userId := utils.GetUserID(c)

	//获取主订单信息
	order := service.Frontend().Buy().GetOrderInfo(orderId)

	//判断当前数据是否合法
	g.Logger.Debugf("%v\n%v\n", userId, order.Uid)
	if userId != order.Uid {
		g.Logger.Error("order data is illegal")
		c.Redirect(302, "/")
	}

	//获取主订单下面的商品信息
	orderItem := service.Frontend().Buy().GetOrderItem(orderId)

	c.HTML(http.StatusOK, "buy/confirm.tmpl",
		utils.MergeMaps(a.Init(c), gin.H{
			"order":     order,
			"orderItem": orderItem,
		}))
}

// GetOrderStatus
// @Tags Buy
// @Summary 获取订单状态
// @Param orderId query string true "订单id"
// @Success 200 {object} response.Response{code=bool,msg=string} "获取订单状态成功"
// @Failure 400 {object} response.Response{code=bool,msg=string} "获取订单状态失败"
// @Router /buy/orderPayStatus [get]
func (a *BuyApi) GetOrderStatus(c *gin.Context) {
	//1、获取订单号
	orderId := c.Query("id")

	//2、查询订单
	order := service.Frontend().Buy().GetOrderInfo(orderId)

	//3、判断当前数据是否合法
	userId := utils.GetUserID(c)
	if userId != order.Uid {
		response.FailWithMessage(c, "传入参数错误")
		return
	}

	//4、判断订单的支付状态
	if order.PayStatus == 1 && order.OrderStatus == 1 {
		response.OkWithMessage(c, "已支付")
	} else {
		response.FailWithMessage(c, "未支付")
	}
}
