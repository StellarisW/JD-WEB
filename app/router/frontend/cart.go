package frontend

import (
	"github.com/gin-gonic/gin"
	"main/app/api"
)

type CartRouter struct{}

func (r *CartRouter) InitCartRouter(Router *gin.RouterGroup) (R *gin.IRoutes) {
	cartRouter := Router.Group("/cart")
	cartApi := api.Frontend().Cart()
	{
		cartRouter.GET("/", cartApi.Get)
		cartRouter.GET("/add", cartApi.Add)
		cartRouter.GET("/inc", cartApi.Inc)
		cartRouter.GET("/dec", cartApi.Dec)
		cartRouter.GET("/del", cartApi.Del)
		cartRouter.GET("/selectOne", cartApi.SelectOne)
		cartRouter.GET("/selectAll", cartApi.SelectAll)
	}
	return
}
