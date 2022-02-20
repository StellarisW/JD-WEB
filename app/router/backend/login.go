package backend

import (
	"github.com/gin-gonic/gin"
	"main/app/api"
)

type LoginRouter struct{}

func (r *IndexRouter) InitLoginRouter(Router *gin.RouterGroup) (R *gin.IRoutes) {
	loginRouter := Router.Group("/login")
	loginApi := api.Backend().Login()
	{
		loginRouter.GET("/", loginApi.Index)
	}
	return
}
