package router

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"main/app/global"
	"main/app/internal/middleware"
	_ "main/app/manifest/docs"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var files []string

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	g.Logger.Info("Starting initialize routers...")
	r := gin.Default()

	// 为session设置redis服务器
	store, _ := redis.NewStore(10, "tcp", g.Config.Redis.Addr, g.Config.Redis.Password, []byte(g.Config.Secret.Common))
	r.Use(sessions.Sessions("mySession", store))

	// 使用zap接收gin框架默认的日志并配置日志归档
	r.Use(middleware.ZapLogger(g.Logger), middleware.ZapRecovery(g.Logger, true))
	//r.POST("/auth", api.GetAuth)

	// Swagger接口文档,输入 /swagger/index.html访问
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//r.POST("/upload", api.UploadImage)

	// 此处可用Nginx代理页面
	// 告诉gin框架模板文件引用的静态文件去哪里找
	r.Static("/resource", "app/resource/public/resource")
	//r.StaticFile("/", "app/resource/template/index.html") // 首页入口
	// 告诉gin框架去哪里找模板文件
	if err := filepath.Walk("app/resource/template", func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".tmpl") {
			files = append(files, path)
		}
		return err
	}); err != nil {
		g.Logger.Error("Walk files failed, err: %v\n", err)
	}

	// 为模板映射函数
	SetFuncMap(r)

	r.LoadHTMLFiles(files...)
	//r.LoadHTMLGlob("app/resource/template/**/*")

	r.StaticFS(g.Config.Local.Path, http.Dir(g.Config.Local.Path)) //// 为用户头像和文件提供静态地址

	// 跨域
	r.Use(middleware.Cors()) // 直接放行全部跨域请求
	// r.Use(middleware.CorsByRules()) // 按照配置的规则放行跨域请求

	// 获取路由组实例
	frontendRouter := GroupApp.Frontend
	backendRouter := GroupApp.Backend

	// 不登陆可以访问的路由
	PublicGroup := r.Group("")
	{
		//健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(200, "ok")
		})
		frontendRouter.InitIndexRouter(PublicGroup)
		frontendRouter.InitAuthRouter(PublicGroup)
		frontendRouter.InitProductRouter(PublicGroup)
		frontendRouter.InitPayRouter(PublicGroup)
		//systemRouter.InitInitRouter(PublicGroup) // 自动初始化相关
		//systemRouter.InitBaseRouter(PublicGroup) // 注册基础功能路由 不做鉴权
	}

	// 登录之后才能访问的路由
	PrivateGroup := r.Group("")
	PrivateGroup.Use(middleware.JWTAuthMiddleware())
	{
		frontendRouter.InitUserRouter(PrivateGroup)
		frontendRouter.InitCartRouter(PrivateGroup)
		frontendRouter.InitBuyRouter(PrivateGroup)

	}

	// 后台控制路由
	BackendGroup := r.Group("/backend")
	BackendGroup.Use()
	{
		backendRouter.InitIndexRouter(BackendGroup)
		backendRouter.InitLoginRouter(BackendGroup)
	}

	g.Logger.Info("Initialize routers successfully!")
	return r
}
