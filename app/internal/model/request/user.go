package request

type LoginReq struct {
	Phone    string `json:"phone" form:"phone" binding:"required"`       // 用户名
	Password string `json:"password" form:"password" binding:"required"` // 密码
}
