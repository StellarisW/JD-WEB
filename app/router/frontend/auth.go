package frontend

import (
	"github.com/gin-gonic/gin"
	"main/app/api"
)

type AuthRouter struct{}

func (r *AuthRouter) InitAuthRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	authRouter := Router.Group("/")
	authApi := api.Frontend().Auth()
	{
		authRouter.GET("/login", authApi.Index)
		authRouter.POST("/login", authApi.Login)
		authRouter.GET("/logout", authApi.LogOut)
		authRouter.GET("/login/oauth", authApi.OauthHandler)
		authRouter.GET("/login/welcome", authApi.OauthWelcome)
		authRouter.GET("/register_step1", authApi.RegisterStep1)
		authRouter.GET("/register_step2", authApi.RegisterStep2)
		authRouter.GET("/register_step3", authApi.RegisterStep3)
		authRouter.GET("/auth/sendCode", authApi.SendCode)
		authRouter.GET("/auth/validateSmsCode", authApi.ValidateSmsCode)
		authRouter.POST("/auth/doRegister", authApi.GoRegister)
		//baseRouter.POST("login", baseApi.Login)
		//baseRouter.POST("captcha", baseApi.Captcha)
	}
	return authRouter
}
