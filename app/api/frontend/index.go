package frontend

import (
	"github.com/gin-gonic/gin"
	g "main/app/global"
	"main/app/internal/service"
	"main/utils"
	"net/http"
	"time"
)

type IndexApi struct {
	BaseApi
}

var insIndex = IndexApi{}

// Get
// InitIndexRouter
// @Tags Base
// @Summary 主页展示
// @Router / [get]
func (a *IndexApi) Get(c *gin.Context) {
	rs := utils.NewRedisStore(g.Redis, "page_", 24*time.Hour, c)
	//开始时间

	banner, err := service.Frontend().Index().GetBanner(rs)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "获取轮播图失败",
		})
	}

	PhoneProduct, err := service.Frontend().Index().GetPhoneProduct(rs)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "获取手机商品列表失败",
		})
	}

	TvProduct, err := service.Frontend().Index().GetTvProduct(rs)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "获取电视商品列表失败",
		})
	}
	c.HTML(http.StatusOK, "index/index.tmpl", utils.MergeMaps(a.Init(c), gin.H{
		"bannerList": banner,
		"phoneList":  PhoneProduct,
		"tvList":     TvProduct,
	}))

}
