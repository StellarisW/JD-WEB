package backend

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	g "main/app/global"
	"main/app/internal/model"
	"main/app/internal/model/response"
	"net/http"
)

type IndexApi struct{}

var insIndex = IndexApi{}

func (a *IndexApi) Index(c *gin.Context) {
	session := sessions.Default(c)
	IUserinfo := session.Get("userinfo") //.(model.Administrator)
	if IUserinfo != nil {
		userinfo := IUserinfo.(model.Administrator)
		g.Logger.Debugf("%v\n", userinfo)
		roleId := userinfo.RoleId
		var auth []model.Auth
		g.DB.Select(&auth, "select * from auth where module_id=? order by sort desc", 0)
		for k := range auth {
			var authItem []model.Auth
			g.DB.Select(&authItem, "select * from auth sort by auth.sort desc")
			auth[k].AuthItem = authItem
		}
		//获取当前部门拥有的权限，并把权限ID放在一个MAP对象里面
		var roleAuth []model.RoleAuth
		g.DB.Select(&roleAuth, "select * from role_auth where role_id=?", roleId)
		roleAuthMap := make(map[int]int)
		for _, v := range roleAuth {
			roleAuthMap[v.AuthId] = v.AuthId
		}
		for i := 0; i < len(auth); i++ {
			if _, ok := roleAuthMap[auth[i].Id]; ok {
				auth[i].Checked = true
			}
			for j := 0; j < len(auth[i].AuthItem); j++ {
				if _, ok := roleAuthMap[auth[i].AuthItem[j].Id]; ok {
					auth[i].AuthItem[j].Checked = true
				}
			}
		}

		c.HTML(http.StatusOK, "backend/index/index.tmpl", gin.H{
			"username": userinfo.Username,
			"authList": auth,
			"isSuper":  userinfo.IsSuper,
		})
	} else {
		c.HTML(http.StatusOK, "backend/index/index.tmpl", gin.H{
			"username": nil,
			"authList": nil,
			"isSuper":  nil,
		})
	}
}

func (a *IndexApi) Welcome(c *gin.Context) {
	c.HTML(http.StatusOK, "backend/index/welcome.tmpl", nil)
}

func (a *IndexApi) ChangeStatus(c *gin.Context) {
	id := c.Query("id") // TODO:不确定
	if id != "" {
		response.FailWithMessage(c, "非法请求")
		return
	}
	table := c.Query("table") // TODO:不确定
	field := c.Query("field") // TODO:不确定
	_, err := g.DB.Exec("update "+table+" set "+field+"=ABS("+field+"-1) where id=?", id)
	if err != nil {
		response.FailWithMessage(c, "更新数据失败")
		return
	}
	response.OkWithMessage(c, "更新数据成功")
}

func (a *IndexApi) EditNum(c *gin.Context) {
	id := c.Query("id")       // TODO:不确定
	table := c.Query("table") // TODO:不确定
	field := c.Query("field") // TODO:不确定
	num := c.Query("num")     // TODO:不确定
	_, err := g.DB.Exec("update " + table + " set " + field + "=" + num + " where id=" + id)
	if err != nil {
		response.FailWithMessage(c, "修改数量失败")
		return
	}

	response.OkWithMessage(c, "修改数量成功")
}
