package captcha

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	g "main/app/global"
	"main/utils/ginx"
	"net/http"
	"time"
)

const (
	// default captcha attributes
	expiration       = 600 * time.Second
	fieldIDName      = "captcha_id"
	fieldCaptchaName = "captcha"
	cachePrefix      = "captcha_"
	urlPrefix        = "/captcha"
)

// Captcha struct
type Captcha struct {
	// beego cache store
	Store  *RedisStore
	driver *base64Captcha.DriverDigit
	Cpt    *base64Captcha.Captcha

	// captcha expiration seconds
	Expiration time.Duration

	// cache key prefix
	CachePrefix string
	UrlPrefix   string
}

type Req struct {
	Id     string `json:"id" form:"id" binding:"required"`
	Answer string `json:"answer" form:"answer" binding:"required"`
}

type LoginDto struct {
	Username string `json:"username" binding:"required" msg:"用户名不能为空"`
	Password string `json:"password" binding:"min=3,max=6" msg:"密码长度不能小于3大于6"`
	Email    string `json:"email" binding:"email" msg:"邮箱地址格式不正确"`
}

func NewCaptcha() *Captcha {
	config := g.Config.Auth.Captcha
	c := &Captcha{}
	c.Store = NewDefaultRedisStore()
	c.driver = base64Captcha.NewDriverDigit(config.ImgHeight, config.ImgWidth, config.KeyLong, 0.7, 80)
	c.Cpt = base64Captcha.NewCaptcha(c.driver, c.Store) // v8下使用redis

	c.Expiration = expiration
	c.CachePrefix = cachePrefix
	c.UrlPrefix = urlPrefix

	return c
}

//
func (c *Captcha) Handler(ct *gin.Context) {
	if len(ginx.Input().Query(ct, "load")) > 0 {
		id, b64s, err := c.Cpt.Generate()
		if err != nil {
			ginx.Output().SetStatus(500)
			ct.Writer.WriteString("captcha reload error")
			g.Logger.Errorf("Reload Create Captcha failed, err: %v\n", err)
			return
		}
		ct.JSON(http.StatusOK, gin.H{
			"msg": "操作成功",
			"data": gin.H{
				"captcha_id":  id,
				"captcha_img": b64s,
			},
		})
	} else {
		//userDto := &LoginDto{} // 要指针类型
		//if err := ct.ShouldBindJSON(userDto); err != nil {
		//	ct.JSON(200, gin.H{"message": err.Error()})
		//} else {
		//	ct.JSON(200, gin.H{
		//		"message": userDto,
		//	})
		//}
		var req Req
		err := ct.ShouldBindJSON(&req)
		if err != nil {
			g.Logger.Errorf("Get captcha request failed, err: %v\n", err)
			ct.JSON(http.StatusOK, gin.H{
				"code": "1",
				"msg":  "操作失败",
			})
		} else {
			if isCorrect := c.Store.Verify(req.Id, req.Answer, false); isCorrect == true {
				ct.JSON(http.StatusOK, gin.H{
					"code": "0",
					"msg":  "验证成功",
				})
			} else {
				ct.JSON(http.StatusOK, gin.H{
					"code": "1",
					"msg":  "验证失败",
				})
			}
		}
	}
}

// Key generate key string
func (c *Captcha) Key(id string) string {
	return c.CachePrefix + id
}
