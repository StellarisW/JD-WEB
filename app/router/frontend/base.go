package frontend

import (
	"github.com/gin-gonic/gin"
	"main/app/api"
)

type IndexRouter struct{}

func (r *IndexRouter) InitIndexRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	indexRouter := Router.Group("/")
	indexApi := api.Frontend().Index()
	{
		indexRouter.GET("/", indexApi.Get)
		//baseRouter.POST("login", baseApi.Login)
		//baseRouter.POST("captcha", baseApi.Captcha)
	}
	return indexRouter
}
