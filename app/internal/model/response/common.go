package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Ok   bool        `json:"success"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	SUCCESS = true
	ERROR   = false
)

func Result(c *gin.Context, ok bool, msg string, data interface{}) {
	if ok == true {
		c.JSON(http.StatusOK, Response{
			ok,
			data,
			msg,
		})
	} else {
		c.JSON(http.StatusBadRequest, Response{
			ok,
			data,
			msg,
		})
	}
}

func Ok(c *gin.Context) {
	Result(c, SUCCESS, "操作成功", map[string]interface{}{})
}

func OkWithMessage(c *gin.Context, message string) {
	Result(c, SUCCESS, message, map[string]interface{}{})
}

func OkWithData(c *gin.Context, data interface{}) {
	Result(c, SUCCESS, "操作成功", data)
}

func OkWithDetailed(c *gin.Context, message string, data interface{}) {
	Result(c, SUCCESS, message, data)
}

func Fail(c *gin.Context) {
	Result(c, ERROR, "操作失败", map[string]interface{}{})
}

func FailWithMessage(c *gin.Context, message string) {
	Result(c, ERROR, message, map[string]interface{}{})
}

func FailWithDetailed(c *gin.Context, message string, data interface{}) {
	Result(c, ERROR, message, data)
}
