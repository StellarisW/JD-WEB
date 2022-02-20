package boot

import (
	"github.com/gin-gonic/gin"
	g "main/app/global"
	"main/app/router"
	"net/http"
	"time"
)

func ServerSetup() {

	config := g.Config.Server
	gin.SetMode(config.Mode)           // 服务器运行模式(Debug/Release)
	routersInit := router.InitRouter() //初始化路由
	server := &http.Server{
		Addr:           config.Addr(),
		Handler:        routersInit,
		ReadTimeout:    config.ReadTimeOut * time.Second,
		WriteTimeout:   config.WriteTimeout * time.Second,
		MaxHeaderBytes: 1 << 20, // 请求头最大大小 16MB
	}

	g.Logger.Infof("Server running on %s ...\n", config.Addr())
	g.Logger.Error(server.ListenAndServe().Error())

}
