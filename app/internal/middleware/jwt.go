package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	g "main/app/global"
	"main/app/internal/model/response"
	"main/utils"
	"main/utils/cookie"
	"time"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		// 这里的具体实现方式要依据你的实际业务情况决定
		var token string
		ok := cookie.Get(c, "x-token", &token)
		if token == "" || !ok {
			response.FailWithDetailed(c, "未登录或非法访问", gin.H{"reload": true})
			c.Abort()
			//c.Next()
			return
		}

		if utils.IsBlacklist(token) {
			response.FailWithDetailed(c, "您的帐户异地登陆或令牌失效", gin.H{"reload": true})
			c.Abort()
			//c.Next()
			return
		}

		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		// parseToken 解析token包含的信息
		j := utils.NewJWT()
		mc, err := j.ParseToken(token)
		if err != nil {
			if err == utils.TokenExpired {
				response.FailWithDetailed(c, "授权已过期", gin.H{"reload": true})
				c.Abort()
				//c.Next()
				return
			}
			response.FailWithDetailed(c, err.Error(), gin.H{"reload": true})
			c.Abort()
			//c.Next()
			return
		}
		// 用户被删除的逻辑 需要优化 此处比较消耗性能 如果需要 请自行打开
		//if err, _ = userService.FindUserByUuid(claims.UUID.String()); err != nil {
		//	_ = jwtService.JsonInBlacklist(system.JwtBlacklist{Jwt: token})
		//	response.FailWithDetailed(gin.H{"reload": true}, err.Error(), c)
		//	c.Abort()
		//}
		if mc.ExpiresAt.Unix()-time.Now().Unix() < mc.BufferTime {
			mc.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Duration(g.Config.Auth.JWT.ExpiresTime) * time.Second))
			newToken, _ := j.GenerateToken(*mc)
			newClaims, _ := j.ParseToken(newToken)
			//c.Header("x-token", newToken)
			cookie.Set(c, "x-token", newToken)
			//c.Header("new-expires-at", strconv.FormatInt(newClaims.ExpiresAt.Unix(), 10))
			val, err := g.Redis.Get(c, newClaims.BaseClaims.Phone).Result()
			if err != nil {
				g.Logger.Errorf("Get jwt by redis failed, err: %v\n", err)
			} else {
				// 当之前的取成功时才进行拉黑操作
				j.JsonInBlackList(utils.JwtBlackList{Jwt: val})
			}
			// 无论如何都要记录当前的活跃状态
			_ = g.Redis.Set(c, newToken, newClaims.BaseClaims.Phone, time.Duration(g.Config.Auth.JWT.ExpiresTime))
		}
		c.Set("claims", mc)
		// 将当前请求的username信息保存到请求的上下文c上
		//c.Set("phone", mc.BaseClaims.ID)
		c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
	}
}
