package backend

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginApi struct{}

var insLogin = LoginApi{}

func (a *LoginApi) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "backend/login/index.tmpl", nil)
}
