package frontend

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	g "main/app/global"
	"main/app/internal/model"
	"main/app/internal/model/request"
	"main/app/internal/model/response"
	"main/app/internal/service"
	"main/utils"
	"main/utils/cookie"
	"net/http"
	"os"
	"strings"
	"time"
)

type AuthApi struct {
	BaseApi
}

var insAuth = AuthApi{}

type AccessTokenResponse struct {
	AccessToken string `json:"access_token" binding:"required"`
}

// Index
// @Tags Login
// @Summary 登录页面展示
// @Produce text/html
// @Router /login [get]
func (a *AuthApi) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "login/index.tmpl", gin.H{
		"prevPage": c.Request.Referer(),
	})
}

// Login
// @Tags Login
// @Summary 登录验证
// @Produce  application/json
// @Param data body request.LoginReq true "用户名, 密码"
// @Success 200 "签发用户token"
// @Failure 400 {object} response.Response{ok=bool,msg=string} "返回错误信息"
// @Header 200 {string} token "签发用户token"
// @Router /login [post]
func (a *AuthApi) Login(c *gin.Context) {
	var req request.LoginReq
	_ = c.ShouldBind(&req)
	g.Logger.Debugf("%v\n", req)
	user := &model.User{Phone: req.Phone, Password: req.Password}
	user, err := service.Frontend().Auth().Login(user)
	if err != nil {
		response.FailWithMessage(c, "用户不存在或错误")
	} else {
		a.tokenNext(c, *user)
	}
}

func (a *AuthApi) tokenNext(c *gin.Context, user model.User) {
	j := utils.NewJWT()
	claims := j.CreateClaims(utils.BaseClaims{
		Id:         user.Id,
		Phone:      user.Phone,
		LastIp:     user.LastIp,
		Email:      user.Email,
		Status:     user.Status,
		UpdateTime: user.UpdateTime,
		CreateTime: user.CreateTime,
	})
	tokenString, err := j.GenerateToken(claims)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "生成JWT失败",
		})
	}
	rs := utils.NewRedisStore(g.Redis, "jwt_", 24*7*time.Hour, c)
	var jwtStr string
	if hasJwtStr := rs.Get(rs.PreKey+user.Phone, &jwtStr); hasJwtStr == false {
		rs.Set(user.Phone, tokenString)
		response.OkWithData(c, response.LoginRes{
			User:      user,
			Token:     tokenString,
			ExpiresAt: g.Config.Auth.JWT.ExpiresTime,
		})
	} else {
		var blackJWT utils.JwtBlackList
		blackJWT.Id = user.Phone
		blackJWT.Jwt = jwtStr
		if err := j.JsonInBlackList(blackJWT); err != nil {
			response.FailWithMessage(c, "jwt作废失败")
			return
		}
		rs.Set(user.Phone, tokenString)
		//c.Header("x-token", tokenString)
		//c.Set("x-token", tokenString)
		cookie.Set(c, "x-token", tokenString)
		response.OkWithDetailed(c, "登录成功",
			response.LoginRes{
				User:      user,
				Token:     tokenString,
				ExpiresAt: g.Config.Auth.JWT.ExpiresTime,
			})
	}
}

// LogOut
// @Tags Login
// @Summary 登出
// @Produce  application/json
// @Router /logout [get]
func (a *AuthApi) LogOut(c *gin.Context) {
	//c.SetSameSite()
	//c.SetCookie()
	//http.SetCookie()
	//cookie.se
	//models.Cookie.Remove(c.Ctx, "userinfo", "")
	cookie.Remove(c, "x-token", "")
	c.Redirect(302, c.Request.Referer())
}

// OauthHandler
// @Tags Login
// @Summary Oauth认证
// @Success 200 "重定向"
// @Header 200 {string} Location "/login/welcome/?access_token="
// @Router /login/oauth [get]
func (a *AuthApi) OauthHandler(c *gin.Context) {
	config := g.Config.Auth.Oauth
	r := c.Request
	w := c.Writer
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(os.Stdout, "could not parse query: %v", err)
		w.WriteHeader(http.StatusBadRequest)
	}
	code := r.FormValue("code")

	reqURL := fmt.Sprintf("https://github.com/login/oauth/access_token?"+
		"client_id=%s&client_secret=%s&code=%s", config.ClientID, config.ClientSecret, code)
	g.Logger.Debug(reqURL)
	req, err := http.NewRequest(http.MethodPost, reqURL, nil)
	if err != nil {
		fmt.Fprintf(os.Stdout, "could not create HTTP request: %v", err)
		w.WriteHeader(http.StatusBadRequest)
	}
	req.Header.Set("accept", "application/json")
	httpClient := http.Client{}
	res, err := httpClient.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stdout, "could not send HTTP request: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	defer res.Body.Close()

	var t AccessTokenResponse
	if err := json.NewDecoder(res.Body).Decode(&t); err != nil {
		fmt.Fprintf(os.Stdout, "could not parse JSON response: %v", err)
		w.WriteHeader(http.StatusBadRequest)
	}
	w.Header().Set("Location", "/login/welcome/?access_token="+t.AccessToken)
	w.WriteHeader(http.StatusFound)
}

// OauthWelcome
// @Tags Login
// @Summary Oauth欢迎页面
// @Produce text/html
// @Router /login/welcome [get]
func (a *AuthApi) OauthWelcome(c *gin.Context) {
	c.HTML(http.StatusOK, "login/welcome.tmpl", nil)
}

// RegisterStep1
// @Tags Register
// @Summary 注册页面展示
// @Produce text/html
// @Router /register_step1 [get]
func (a *AuthApi) RegisterStep1(c *gin.Context) {
	c.HTML(http.StatusOK, "register_step1.tmpl", nil)
}

// RegisterStep2
// @Tags Register
// @Summary 验证图形验证码
// @Accept application/json
// @Produce text/html
// @Param captcha query string true "图形验证码答案"
// @Success 200 "显示页面"
// @Failure 302 "重定向"
// @Header 302 {string} Location "/register_step1"
// @Router /register_step2 [get]
func (a *AuthApi) RegisterStep2(c *gin.Context) {
	sign := c.Query("sign")
	captchaAnswer := c.Query("captcha_answer")
	////验证图形验证码和前面是否正确
	//sessionPhotoCode := c.GetSession("phone_code")
	//if phone_code != sessionPhotoCode {
	//	c.Redirect("/auth/registerStep1", 302)
	//	return
	//}
	var userTemp []model.UserSms
	g.DB.Select(&userTemp, "select * from user_sms where sign=?", sign)
	if len(userTemp) > 0 {
		c.HTML(http.StatusOK, "register_step2.tmpl", gin.H{
			"sign":           sign,
			"captcha_answer": captchaAnswer,
			"phone":          userTemp[0].Phone,
		})
	} else {
		c.Redirect(http.StatusFound, "/register_step1")
		return
	}
}

// RegisterStep3
// @Tags Register
// @summary 验证手机验证码
// @Product text/html
// @Param sms_code query string true "手机验证码"
// @Success 200 "显示页面"
// @Failure 302 "重定向"
// @Header 302 {string} Location "/registerStep1"
// @register_step3 [get]
func (a *AuthApi) RegisterStep3(c *gin.Context) {
	sign := c.Query("sign")
	sms_code := c.Query("sms_code")
	//sessionSmsCode := c.GetSession("sms_code")
	//if sms_code != sessionSmsCode && sms_code != "5259" {
	//	c.Redirect("/auth/registerStep1", 302)
	//	return
	//}
	var userTemp []model.UserSms
	g.DB.Select(&userTemp, "select * from user_sms where sign=?", sign)
	if len(userTemp) > 0 {
		c.HTML(http.StatusOK, "register_step3.tmpl", gin.H{
			"sign":     sign,
			"sms_code": sms_code,
		})
	} else {
		c.Redirect(302, "/registerStep1")
		return
	}
}

// GoRegister
// @Tags Register
// @Summary 验证注册提交表单
// @Param Form formData string true "表单"
// @Success 200 "注册用户到数据库"
// @Failure 302 "重定向"
// @Header 302 {string} Location "/register_step1"
// @Router /auth/doRegister [post]
func (a *AuthApi) GoRegister(c *gin.Context) {
	sign := c.PostForm("sign")
	password := c.PostForm("password")
	var userTemp []model.UserSms
	g.DB.Select(&userTemp, "select * from user_sms where sign=?", sign)
	ip := strings.Split(c.Request.RemoteAddr, ":")[0]
	if len(userTemp) > 0 {
		user := model.User{
			Phone:    userTemp[0].Phone,
			Password: utils.Md5(password),
			LastIp:   ip,
		}
		g.DB.Table("user").Select("phone", "password", "last_ip").Create(&user)
		c.Redirect(302, "/")
	} else {
		c.Redirect(302, "/register_step1")
	}
}

// SendCode
// @Tags Register
// @Summary 发送验证码
// @Param phone query string true "手机号"
// @success 200 {object} response.Response{code=bool,data=response.SmsRes,msg=string} "发送成功"
// @Failure 400 {object} response.Response{code=bool,msg=string} "发送失败"
// @Router /auth/sendCode [get]
func (a *AuthApi) SendCode(c *gin.Context) {
	phone := c.Query("phone")
	var user []model.User
	g.DB.Where("phone=?", phone).Find(&user)
	if len(user) > 0 {
		response.FailWithMessage(c, "此用户已存在")
		return
	}

	addDay := utils.FormatDay()
	ip := strings.Split(c.Request.RemoteAddr, ":")[0]
	sign := utils.Md5(phone + addDay) //签名
	smsCode := utils.GetRandomNum()
	var userTemp []model.UserSms
	g.DB.Select(&userTemp, "select * from user_sms where add_day=? and phone=?", addDay, phone)
	var sendCount int
	g.DB.Table("user_sms").Where("add_day=? AND ip=?").Find(&sendCount)
	//验证IP地址今天发送的次数是否合法
	if sendCount <= 10 {
		if len(userTemp) > 0 {
			//验证当前手机号今天发送的次数是否合法
			if userTemp[0].SendCount < 5 {
				utils.SendMsg(smsCode)
				//a.SetSession("sms_code", sms_code)
				oneUserSms := model.UserSms{}
				g.DB.Where("id=?", userTemp[0].Id).Find(&oneUserSms)
				oneUserSms.SendCount += 1
				g.DB.Exec("update user_sms set send_count=? where id=?", oneUserSms.SendCount, oneUserSms.Id)
				response.OkWithDetailed(c, "短信发送成功", response.SmsRes{
					Sign:    sign,
					SmsCode: smsCode,
				})
				return
			} else {
				response.FailWithMessage(c, "当前手机号今天发送短信数已达上线")
				return
			}

		} else {
			utils.SendMsg(smsCode)
			//a.SetSession("sms_code", sms_code)
			//发送验证码 并给userTemp写入数据
			g.DB.Exec("insert into user_sms(ip,phone,send_count,add_day,sign) values (?,?,?,?,?)",
				ip, phone, 1, addDay, sign)
			response.OkWithDetailed(c, "短信发送成功", response.SmsRes{
				Sign:    sign,
				SmsCode: smsCode,
			})
			return
		}
	} else {
		response.FailWithMessage(c, "此IP今天发送次数已经达到上限，明天再试")
		return
	}
}

// ValidateSmsCode
// @Tags Register
// @Summary 验证手机验证码
// @Param sms_code query string true "验证"
// @Success 200 {object} response.Response{code=bool,msg=string} "验证成功"
// @Failure 400 {object} response.Response{code=bool,msg=string} "发送失败"
// @Router /auth/validateSmsCode [get]
func (a *AuthApi) ValidateSmsCode(c *gin.Context) {
	sign := c.Query("sign")
	sms_code := c.Query("sms_code")

	userTemp := []model.UserSms{}
	g.DB.Select(&userTemp, "select * from user_sms where sign=?", sign)
	if len(userTemp) == 0 {
		response.FailWithMessage(c, "参数错误")
		return
	}

	//sessionSmsCode := c.GetSession("sms_code")
	//if sessionSmsCode != sms_code && sms_code != "5259" {
	//	c.Data["json"] = map[string]interface{}{
	//		"success": false,
	//		"msg":     "输入的短信验证码错误",
	//	}
	//	c.ServeJSON()
	//	return
	//}

	if sms_code != "5259" {
		response.FailWithMessage(c, "输入的短信验证码错误")
	}

	if userTemp[0].GetUpdatedTime().Add(time.Minute).After(time.Now()) {
		response.FailWithMessage(c, "验证码已过期")
		return
	}
	response.OkWithMessage(c, "验证成功")
}
