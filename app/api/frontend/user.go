package frontend

import (
	"github.com/gin-gonic/gin"
	g "main/app/global"
	"main/app/internal/model"
	"main/app/internal/model/response"
	"main/app/internal/service"
	"main/utils"
	"net/http"
	"strconv"
)

type UserApi struct {
	BaseApi
}

var insUser = UserApi{}

// Index
// @Tags User
// @Summary 用户页面
// @Success 200 "页面展示"
// @Router /user [get]
func (a *UserApi) Index(c *gin.Context) {

	claims, _ := utils.GetUserInfo(c)
	user := service.Frontend().User().GetUserInfo(claims)

	HelloMsg := service.Frontend().User().GetTimeSlot()

	waitPay, waitRec := service.Frontend().User().GetOrderNum(user.Id)

	c.HTML(http.StatusOK, "user/index.tmpl",
		utils.MergeMaps(a.Init(c), gin.H{
			"user":     user,
			"Hello":    HelloMsg,
			"wait_pay": waitPay,
			"wait_rec": waitRec,
		}))
}

// GetOrderList
// @Tags User
// @Summary 获取订单
// @Param uid query string true "订单id"
// @Success 200 "展示订单"
// @Router /user/order [get]
func (a *UserApi) GetOrderList(c *gin.Context) {
	// 获取搜索关键词
	where := "uid=? "
	keywords := c.Query("keywords")
	if keywords != "" {
		service.Frontend().User().GetSearchResult(keywords, &where)
	}
	// 获取筛选条件
	orderStatus := c.Query("order_status")
	where += "and order_status=" + orderStatus

	// 获取当前用户
	userId := utils.GetUserID(c)
	// 获取当前用户的订单信息并分页
	pageString := c.Query("page")
	page, _ := strconv.Atoi(pageString)
	if page == 0 {
		page = 1
	}

	order, totalPages, page := service.Frontend().User().GetOrderList(userId, page, where)

	// 计算总数量
	c.HTML(http.StatusOK, "order.tmpl",
		utils.MergeMaps(a.Init(c), gin.H{
			"orderStatus": orderStatus,
			"keywords":    keywords,
			"order":       order,
			"totalPages":  totalPages,
			"page":        page,
		}))
}

// GetOrderInfo
// @Tags User
// @Summary 获取订单详细信息
// @Param uid query string true "订单id"
// @Success 200 "展示订单信息"
// @Router /user/orderInfo [get]
func (a *UserApi) GetOrderInfo(c *gin.Context) {
	idString := c.Query("id")
	id, _ := strconv.Atoi(idString)
	userId := utils.GetUserID(c)
	order := service.Frontend().User().GetOrderInfo(id, userId)
	if order.OrderId == "" {
		c.Redirect(302, "/")
	}
	c.HTML(http.StatusOK, "order_info.tmpl",
		utils.MergeMaps(a.Init(c), gin.H{
			"order": order,
		}))
}

// GetCollectProduct
// @Tags User
// @Summary 获取收藏商品
// @Param page query string false "页面数"
// @Success 200 "展示收藏商品"
// @Router /user/collect [get]
func (a *UserApi) GetCollectProduct(c *gin.Context) {
	// 获取当前用户
	userId := utils.GetUserID(c)

	// 获取当前用户的订单信息并分页
	pageString := c.Query("page")
	page, _ := strconv.Atoi(pageString)
	if page == 0 {
		page = 1
	}

	product, totalPages, page := service.Frontend().User().GetCollectProduct(userId, page)
	c.HTML(http.StatusOK, "collect.tmpl",
		utils.MergeMaps(a.Init(c), gin.H{
			"product":    product,
			"totalPages": totalPages,
			"page":       page,
		}))
}

// GetOneAddress
// @Tags User
// @Summary 获取一个收货地址
// @Param address_id query string true "收货地址id"
// @Success 200 {object} response.Response{code=bool,data=model.Address,msg=string} "展示收货地址"
// @Router /user/address/getOne [get]
func (a *UserApi) GetOneAddress(c *gin.Context) {
	addressIdString := c.Query("address_id")
	addressId, _ := strconv.Atoi(addressIdString)
	address := service.Frontend().User().GetOneAddress(addressId)

	response.OkWithData(c, gin.H{
		"result": address,
	})
}

// AddAddress
// @Tags User
// @Summary 增加收货地址
// @Param token header string true "jwt"
// @Success 200 {object} response.Response{code=bool,data=[]model.Address,msg=string} "增加收货地址"
// @Failure 400 {object} response.Response{code=bool,msg=string} "增加收货地址失败"
// @Router /user/address/add [post]
func (a *UserApi) AddAddress(c *gin.Context) {
	var toAdd model.Address
	userId := utils.GetUserID(c)
	_ = c.ShouldBind(&toAdd)
	toAdd.Uid = userId
	toAdd.DefaultAddress = 1
	g.Logger.Debugf("%v\n%v\n", userId, toAdd)
	ok := service.Frontend().User().IsAddressLimit(userId)
	if ok {
		response.FailWithMessage(c, "增加收货地址失败，收货地址数量超过限制")
	}

	allAddress := service.Frontend().User().AddAddress(userId, toAdd)

	response.OkWithData(c, gin.H{
		"result": allAddress,
	})
}

// EditAddress
// @Tags User
// @Summary 编辑收货地址
// @Param token header string true "jwt"
// @Success 200 {object} response.Response{code=bool,data=[]model.Address,msg=string} "增加收货地址成功"
// @Router /user/address/edit [post]
func (a *UserApi) EditAddress(c *gin.Context) {
	var address model.Address
	userId := utils.GetUserID(c)
	err := c.ShouldBind(&address)
	address.DefaultAddress = 1
	g.Logger.Debugf("%v\n%v\n%v\n", userId, address, err)

	allAddress := service.Frontend().User().EditAddress(userId, address)

	response.OkWithData(c, gin.H{
		"result": allAddress,
	})
}

// ChangeToDefaultAddress
// @Tags User
// @Summary 设置默认收货地址
// @Param token header string true "jwt"
// @Param address_id query string ture "收货地址id"
// @Success 200 {object} response.Response{code=bool,msg=string} "设置默认收货地址成功"
// @Router /user/address/changeDefault [post]
func (a *UserApi) ChangeToDefaultAddress(c *gin.Context) {
	userId := utils.GetUserID(c)
	addressIdString := c.Query("address_id")
	addressId, _ := strconv.Atoi(addressIdString)
	service.Frontend().User().ChangeToDefaultAddress(userId, addressId)

	response.OkWithMessage(c, "更新默认收获地址成功")
}
