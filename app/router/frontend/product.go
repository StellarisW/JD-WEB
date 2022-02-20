package frontend

import (
	"github.com/gin-gonic/gin"
	"main/app/api"
)

type ProductRouter struct{}

func (r *ProductRouter) InitProductRouter(Router *gin.RouterGroup) (R *gin.IRoutes) {
	productRouter := Router.Group("/")
	productApi := api.Frontend().Product()
	{
		productRouter.GET("/category/:id", productApi.GetCateList)
		productRouter.GET("/item/:id", productApi.GetProductItem)
		productRouter.GET("/product/collect", productApi.CollectProduct)
		productRouter.GET("/product/getImgList", productApi.GetImgList)
	}
	return
}
