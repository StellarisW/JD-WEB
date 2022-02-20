package frontend

import (
	"github.com/gin-gonic/gin"
	"main/app/api"
)

type UserRouter struct{}

func (r *UserRouter) InitUserRouter(Router *gin.RouterGroup) (R *gin.IRoutes) {
	userRouter := Router.Group("/user")
	userApi := api.Frontend().User()
	{
		userRouter.GET("/", userApi.Index)

		userRouter.GET("/collect", userApi.GetCollectProduct)
		userRouter.GET("/order", userApi.GetOrderList)
		userRouter.GET("/orderInfo", userApi.GetOrderInfo)

		userRouter.GET("/address/getOne", userApi.GetOneAddress)
		userRouter.POST("/address/add", userApi.AddAddress)
		userRouter.POST("/address/edit", userApi.EditAddress)
		userRouter.POST("/address/changeDefault", userApi.ChangeToDefaultAddress)
	}
	return
}
