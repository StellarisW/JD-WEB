package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	g "main/app/global"
	"main/app/internal/model"
	"net/url"
	"strings"
)

func BackendAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		pathname := c.Request.URL.String()
		session := sessions.Default(c)
		IUserinfo := session.Get("userinfo") //.(model.Administrator)
		if IUserinfo == nil {
			if pathname != "/backend/login" &&
				pathname != "/backend/login/goLogin" &&
				pathname != "/backend/login/verifyCode" {
				c.Redirect(302, "/backend/login")
				c.Next()
				return
			}
		} else {
			userinfo := IUserinfo.(model.Administrator)
			pathname = strings.Replace(pathname, "/"+"backend", "", 1)
			urlPath, _ := url.Parse(pathname)
			if userinfo.IsSuper == 0 && !excludeAuthPath(urlPath.Path) {
				roleId := userinfo.RoleId
				var roleAuth []model.RoleAuth
				_ = g.DB.Select(&roleAuth, "select * from role_auth where role_id=?", roleId)
				roleAuthMap := make(map[int]int)
				for _, v := range roleAuth {
					roleAuthMap[v.AuthId] = v.AuthId
				}
				auth := model.Auth{}
				_ = g.DB.Get(&auth, "select * from auth where url=?", urlPath.Path)
				if _, ok := roleAuthMap[auth.Id]; !ok {
					c.Writer.WriteString("没有权限")
					c.Next()
					return
				}
			}
		}
	}
}

//检验路径权限
func excludeAuthPath(urlPath string) bool {
	excludeAuthPathSlice := strings.Split("/,/welcome,/login/logout", ",")
	for _, v := range excludeAuthPathSlice {
		if v == urlPath {
			return true
		}
	}
	return false
}
