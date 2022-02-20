package backend

import (
	"github.com/gin-gonic/gin"
	"main/app/api"
)

type IndexRouter struct{}

func (r *IndexRouter) InitIndexRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	indexRouter := Router.Group("/")
	indexApi := api.Backend().Index()
	{
		indexRouter.GET("/", indexApi.Index)
		indexRouter.GET("/welcome", indexApi.Welcome)
		indexRouter.GET("/changeStatus", indexApi.ChangeStatus)
		indexRouter.GET("/editNum", indexApi.Index)
	}
	return
}
