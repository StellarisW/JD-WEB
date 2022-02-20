package frontend

import (
	"github.com/gin-gonic/gin"
	"main/app/internal/model"
	"main/app/internal/model/response"
	"main/app/internal/service"
	"main/utils"
	"main/utils/cookie"
	"net/http"
	"strconv"
)

type CartApi struct {
	BaseApi
}

var insCart = CartApi{}

// Get 购物车展示
// @Tags Cart
// @Summary 购物车展示
// @Param cartList header string true "cartList"
// @Success 200 "展示购物车"
// @Router /cart [get]
func (a *CartApi) Get(c *gin.Context) {

	var cartList []model.Cart
	cookie.Get(c, "cartList", &cartList)

	//执行计算总价
	allPrice := service.Frontend().Cart().CalAllPrice(cartList)
	c.HTML(http.StatusOK, "cart/index.tmpl",
		utils.MergeMaps(a.Init(c), gin.H{
			"cartList": cartList,
			"allPrice": allPrice,
		}))
}

// Add
// @Tags Cart
// @Summary 购物车增加商品
// @Param product_id query string ture "商品id"
// @Param color_id query string ture "颜色id"
// @Success 200 "增加商品成功"
// @Router /cart/add [get]
func (a *CartApi) Add(c *gin.Context) {
	productIdString := c.Query("product_id")
	colorIdString := c.Query("color_id")
	colorId, _ := strconv.Atoi(colorIdString)
	productId, _ := strconv.Atoi(productIdString)

	currentData, product := service.Frontend().Cart().AddPruduct(productId, colorId)

	//2.判断购物车有没有数据（cookie）
	var cartList []model.Cart
	cookie.Get(c, "cartList", &cartList)
	service.Frontend().Cart().UpdateCookie(&cartList, currentData)
	cookie.Set(c, "cartList", cartList)

	c.HTML(http.StatusOK, "cart/add_success.tmpl",
		utils.MergeMaps(a.Init(c), gin.H{
			"product": product,
		}))
}

// Inc
// @Tags Cart
// @Summary 购物车增加商品数量
// @Param product_id query string ture "商品id"
// @Param product_color query string ture "商品颜色"
// @Param cartList header string ture "cartList"
// @Success 200 {object} response.Response{code=bool,data=response.CartProductRes,msg=string} "增加商品数量成功"
// @Router /cart/inc [get]
func (a *CartApi) Inc(c *gin.Context) {
	var flag bool

	productIdString := c.Query("product_id")
	productId, _ := strconv.Atoi(productIdString)
	productColor := c.Query("product_color")
	productAttr := ""

	var cartList []model.Cart
	cookie.Get(c, "cartList", &cartList)

	// num: 商品数量 currentAllPrice: 商品总价 allPrice: 所有商品总价
	flag, num, currentAllPrice, allPrice := service.Frontend().Cart().IncItem(&cartList, productId, productColor, productAttr)

	if flag {
		cookie.Set(c, "cartList", cartList)
		response.OkWithDetailed(c, "修改数量成功", response.CartProductRes{
			Num:             num,
			CurrentAllPrice: currentAllPrice,
			AllPrice:        allPrice,
		})
	} else {
		response.FailWithMessage(c, "传入参数错误")
	}
}

// Dec
// @Tags Cart
// @Summary 购物车减少商品数量
// @Param product_id query string ture "商品id"
// @Param product_color query string ture "商品颜色"
// @Param cartList header string ture "cartList"
// @Success 200 {object} response.Response{code=bool,data=response.CartProductRes,msg=string} "减少商品数量成功"
// @Router /cart/dec [get]
func (a *CartApi) Dec(c *gin.Context) {
	var flag bool

	productIdString := c.Query("product_id")
	productId, _ := strconv.Atoi(productIdString)
	productColor := c.Query("product_color")
	productAttr := ""

	var cartList []model.Cart
	cookie.Get(c, "cartList", &cartList)

	flag, num, currentAllPrice, allPrice := service.Frontend().Cart().DecItem(&cartList, productId, productColor, productAttr)

	if flag {
		cookie.Set(c, "cartList", cartList)
		response.OkWithDetailed(c, "修改数量成功", gin.H{
			"num":             num,
			"currentAllPrice": currentAllPrice,
			"allPrice":        allPrice,
		})
	} else {
		response.FailWithMessage(c, "传入参数错误")
	}
}

// Del
// @Tags Cart
// @Summary 删除购物车商品
// @Param product_id query string ture "商品id"
// @Param product_color query string ture "商品颜色"
// @Param cartList header string ture "cartList"
// @Success 200 "刷新页面"
// @Router /cart/del [get]
func (a *CartApi) Del(c *gin.Context) {
	productIdString := c.Query("product_id")
	productId, _ := strconv.Atoi(productIdString)
	productColor := c.Query("product_color")
	productAttr := ""

	var cartList []model.Cart
	cookie.Get(c, "cartList", &cartList)
	service.Frontend().Cart().DelItem(&cartList, productId, productColor, productAttr)
	cookie.Set(c, "cartList", cartList)

	c.Redirect(302, "/cart")
}

// SelectOne
// @Tags Cart
// @Summary 选择一个购物车商品
// @Param product_id query string ture "商品id"
// @Param product_color query string ture "商品颜色"
// @Param cartList header string ture "cartList"
// @Success 200 {object} response.Response{code=bool,data=float64,msg=string} "选择一个购物车商品成功"
// @Failure 400 {object} response.Response{code=bool,msg=string} "选择一个购物车商品失败"
// @Router /cart/selectOne [get]
func (a *CartApi) SelectOne(c *gin.Context) {

	productIdString := c.Query("product_id")
	productId, _ := strconv.Atoi(productIdString)
	productColor := c.Query("product_color")
	productAttr := ""

	var cartList []model.Cart
	cookie.Get(c, "cartList", &cartList)

	flag, allPrice := service.Frontend().Cart().SelectOneItem(&cartList, productId, productColor, productAttr)
	if flag {
		cookie.Set(c, "cartList", cartList)
		response.OkWithDetailed(c, "修改状态成功", gin.H{
			"allPrice": allPrice,
		})
	} else {
		response.FailWithMessage(c, "传入参数错误")
	}
}

// SelectAll 全选/反选
// @Tags Cart
// @Summary 选择全部购物车商品
// @Param cartList header string ture "cartList"
// @Success 200 {object} response.Response{code=bool,data=float64,msg=string} "选择全部购物车商品成功"
// @Router /cart/selectAll [get]
func (a *CartApi) SelectAll(c *gin.Context) {
	flagString := c.Query("flag")
	flag, _ := strconv.Atoi(flagString)

	var cartList []model.Cart
	cookie.Get(c, "cartList", &cartList)

	allPrice := service.Frontend().Cart().SelectAllItem(&cartList, flag)

	cookie.Set(c, "cartList", cartList)

	response.OkWithData(c, gin.H{
		"allPrice": allPrice,
	})
}
