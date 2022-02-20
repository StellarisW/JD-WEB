package frontend

import (
	"github.com/gin-gonic/gin"
	g "main/app/global"
	"main/app/internal/service"
	"main/utils"
	"main/utils/captcha"
	"net/http"
	"net/url"
	"time"
)

type BaseApi struct{}

var insBase = BaseApi{}

var rs = captcha.NewDefaultRedisStore()

func (a *BaseApi) Init(c *gin.Context) gin.H {
	rs := utils.NewRedisStore(g.Redis, "page_", 24*time.Hour, c)
	// 获取顶部导航
	menu, err := service.Frontend().Base().InitTopMenu(rs)
	if err != nil {
		g.Logger.Errorf("Get topMenu failed, err: %v\n", err)
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "获取顶部导航失败",
		})
	}
	// 获取左侧分类栏
	productCate, err := service.Frontend().Base().InitProductCate(rs)
	if err != nil {
		g.Logger.Errorf("Get productCate failed, err: %v\n", err)
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "获取左侧分类栏失败",
		})
	}
	// 获取中间导航栏
	middleMenu, err := service.Frontend().Base().InitMiddleMenu(rs)
	if err != nil {
		g.Logger.Errorf("Get middleMenu failed, err: %v\n", err)
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "获取中间导航数据失败",
		})
	}
	// 获取用户信息
	userinfo := service.Frontend().Base().IsLogin(c)
	urlPath, _ := url.Parse(c.Request.URL.String())
	return gin.H{
		"topMenuList":     menu,
		"productCateList": productCate,
		"middleMenuList":  middleMenu,
		"userinfo":        userinfo,
		"pathname":        urlPath.Path,
	}
	//response.OkWithData(c, menu, productCate, middleMenu, userinfo, urlPath)
}
