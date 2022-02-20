package frontend

import (
	"github.com/gin-gonic/gin"
	"main/app/internal/model"
	"main/app/internal/model/response"
	"main/app/internal/service"
	"main/utils"
	"math"
	"net/http"
	"strconv"
)

type ProductApi struct {
	BaseApi
}

var insProduct = ProductApi{}

// GetCateList 商品列表展示
// @Tags Product
// @Summary 商品列表展示
// @Param id path string true "商品id"
// @Param page query string true "页面数"
// @Success 200 "返回页面"
// @Router /category/:id [get]
func (a *ProductApi) GetCateList(c *gin.Context) {
	id := c.Param("id")
	cateId, _ := strconv.Atoi(id)
	currentProductCate, subProductCate := service.Frontend().Product().GetCateList(cateId)
	var tempSlice []int
	if currentProductCate.Pid == 0 { // 顶级分类
		// 二级分类
		for i := 0; i < len(subProductCate); i++ {
			tempSlice = append(tempSlice, subProductCate[i].Id)
		}
	}
	tempSlice = append(tempSlice, cateId)

	//当前页
	pageString := c.Query("page")
	page, _ := strconv.Atoi(pageString)
	if page == 0 {
		page = 1
	}
	//每一页显示的数量
	pageSize := 5
	products, count := service.Frontend().Product().GetProductByCate(tempSlice, page, pageSize)
	tpl := currentProductCate.Template
	if tpl == "" {
		tpl = "product/list.tmpl"
	}
	c.HTML(http.StatusOK, tpl,
		utils.MergeMaps(a.Init(c), gin.H{
			"currentProductCate": currentProductCate,
			"subProductCate":     subProductCate,
			"productList":        products,
			"totalPages":         math.Ceil(float64(count) / float64(pageSize)),
			"page":               page,
		}))
}

// GetProductItem
// @Tags Product
// @Summary 商品详细页
// @Param id path string true "商品id"
// @Success 200 "返回页面"
// @Router /item/:id [get]
func (a *ProductApi) GetProductItem(c *gin.Context) {

	id := c.Param("id")
	// 获取商品信息
	productId, _ := strconv.Atoi(id)
	product := service.Frontend().Product().GetProductById(productId)

	// 获取关联商品  RelationProduct
	relationProduct := service.Frontend().Product().GetRelationProduct(product)

	// 获取关联赠品 ProductGift
	productGift := service.Frontend().Product().GetGift(product)

	// 获取关联颜色 ProductColor
	productColor := service.Frontend().Product().GetColor(product)

	// 获取关联配件 ProductFitting
	productFitting := service.Frontend().Product().GetFitting(product)

	// 获取商品关联的图片 ProductImage
	productImage := service.Frontend().Product().GetImg(product)

	// 获取规格参数信息 ProductAttr
	productAttr := service.Frontend().Product().GetAttr(product)

	c.HTML(http.StatusOK, "product/item.tmpl",
		utils.MergeMaps(a.Init(c), gin.H{
			"product":         product,
			"relationProduct": relationProduct,
			"productGift":     productGift,
			"productColor":    productColor,
			"productFitting":  productFitting,
			"productImage":    productImage,
			"productAttr":     productAttr,
		}))
}

// CollectProduct 收藏商品
// @Tags Product
// @Summary 收藏商品
// @Param product_id query string true "商品id"
// @Success 200 {object} response.Response{code=bool,msg=string}
// @Failure 400 {object} response.Response{code=bool,msg=string}
// @Router /product/collect [get]
func (a *ProductApi) CollectProduct(c *gin.Context) {
	productIdString := c.Query("product_id")
	productId, _ := strconv.Atoi(productIdString)

	claims, err := utils.GetUserInfo(c)
	if err != nil {
		response.FailWithMessage(c, "请先登录")
		return
	}

	userinfo := model.User{
		Phone:  claims.Phone,
		LastIp: claims.LastIp,
		Email:  claims.Email,
		Status: claims.Status,
	}
	userinfo.Id = claims.Id
	userinfo.UpdateTime = claims.UpdateTime
	userinfo.CreateTime = claims.CreateTime

	isExist := userinfo.Verify()
	if isExist == false {
		response.FailWithMessage(c, "非法用户")
		return
	}

	ok := service.Frontend().Product().CollectProduct(userinfo.Id, productId)
	if ok {
		response.OkWithMessage(c, "收藏成功")
	} else {
		response.OkWithMessage(c, "取消收藏成功")
	}
}

// GetImgList
// @Tags Product
// @Summary 查询商品图库
// @Param product_id query string true "商品id"
// @Param color_id query string true "颜色id"
// @Success 200 {object} response.Response{code=bool,data=model.ProductImage,msg=string}
// @Router /product/getImgList [get]
func (a *ProductApi) GetImgList(c *gin.Context) {

	productIdString := c.Query("product_id")
	colorIdString := c.Query("color_id")
	productId, _ := strconv.Atoi(productIdString)
	colorId, _ := strconv.Atoi(colorIdString)

	//查询商品图库信息
	productImageList := service.Frontend().Product().GetImgList(productId, colorId)

	response.OkWithData(c, gin.H{
		"result": productImageList,
	})
}
