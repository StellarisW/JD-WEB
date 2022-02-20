package frontend

import (
	"github.com/gin-gonic/gin"
	"main/app/api"
)

type PayRouter struct{}

func (r *PayRouter) InitPayRouter(Router *gin.RouterGroup) (R *gin.IRoutes) {
	payRouter := Router.Group("/pay")
	payApi := api.Frontend().Pay()
	{
		payRouter.GET("/alipay", payApi.Alipay)
		payRouter.GET("/alipayNotify", payApi.AlipayNotify)
		payRouter.GET("/alipayReturn", payApi.AlipayReturn)
	}
	return
}
