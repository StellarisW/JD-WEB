package frontend

import (
	"github.com/gin-gonic/gin"
	"main/app/api"
)

type BuyRouter struct{}

func (r *BuyRouter) InitBuyRouter(Router *gin.RouterGroup) (R *gin.IRoutes) {
	buyRouter := Router.Group("/buy")
	buyApi := api.Frontend().Buy()
	{
		buyRouter.GET("/checkout", buyApi.Index)
		buyRouter.POST("/doOrder", buyApi.Process)
		buyRouter.GET("/confirm", buyApi.Confirm)
		buyRouter.GET("/orderPayStatus", buyApi.GetOrderStatus)
	}
	return
}
