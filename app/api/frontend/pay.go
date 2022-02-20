package frontend

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/smartwalle/alipay/v3"
	g "main/app/global"
	"main/app/internal/service"
	"time"
)

type PayApi struct{}

var insPay = PayApi{}

// Alipay
// @Tags Pay
// @Summary alipay
// @Param aliId query string true "从支付宝接口返回的id"
// @Success 200 "重定向"
// @Header 200 {string} Location "/user/order"
// @Router /alipay [get]
func (a *PayApi) Alipay(c *gin.Context) {
	aliId := c.Query("id")

	orderItem := service.Frontend().Buy().GetOrderItem(aliId)
	g.Logger.Debugf("%v\n", aliId)
	var privateKey = g.Config.Secret.Private // 必须，上一步中使用 RSA签名验签工具 生成的私钥
	var client, err = alipay.New("2021000119617698", privateKey, false)
	client.LoadAppPublicCertFromFile("cert/appCertPublicKey_2021000119617698.crt") // 加载应用公钥证书
	client.LoadAliPayRootCertFromFile("cert/alipayRootCert.crt")                   // 加载支付宝根证书
	client.LoadAliPayPublicCertFromFile("cert/alipayCertPublicKey_RSA2.crt")       // 加载支付宝公钥证书

	// 将 key 的验证调整到初始化阶段
	if err != nil {
		fmt.Println(err)
		return
	}

	//计算总价格
	var TotalAmount float64
	for i := 0; i < len(orderItem); i++ {
		TotalAmount = TotalAmount + orderItem[i].ProductPrice
	}
	totalAmount := fmt.Sprintf("%.2f", TotalAmount)
	var p = alipay.TradePagePay{}
	p.NotifyURL = "http://localhost:9090/pay/alipayNotify" // 支付成功后阿里云服务器发送请求到地址
	p.ReturnURL = "http://localhost:9090/pay/alipayNotify" // 支付成功后用户返回到该地址
	p.TotalAmount = totalAmount
	p.Subject = "订单order——" + time.Now().Format("200601021504")
	p.OutTradeNo = "WF" + time.Now().Format("200601021504") + "_" + aliId
	p.ProductCode = "FAST_INSTANT_TRADE_PAY"

	var url, err4 = client.TradePagePay(p)
	if err4 != nil {
		fmt.Println(err4)
	}
	var payURL = url.String()
	c.Redirect(302, payURL)
}

// AlipayNotify
// @Tags Pay
// @Summary 验证支付token
// @Param sign query string true "支付token"
// @Success 200 "更新订单状态"
// @Router /alipayNotify [get]
func (a *PayApi) AlipayNotify(c *gin.Context) {
	var privateKey = g.Config.Secret.Private // 必须，上一步中使用 RSA签名验签工具 生成的私钥
	var client, err = alipay.New("2021000119617698", privateKey, false)

	client.LoadAppPublicCertFromFile("cert/appCertPublicKey_2021000119617698.crt") // 加载应用公钥证书
	client.LoadAliPayRootCertFromFile("cert/alipayRootCert.crt")                   // 加载支付宝根证书
	client.LoadAliPayPublicCertFromFile("cert/alipayCertPublicKey_RSA2.crt")       // 加载支付宝公钥证书

	if err != nil {
		fmt.Println(err)
		return
	}

	req := c.Request
	req.ParseForm()
	g.Logger.Debugf("%v\n", req.Form)
	ok, err := client.VerifySign(req.Form)
	if !ok || err != nil {
		c.Redirect(302, c.Request.Referer())
	}
	rep := c.Writer
	var noti, _ = client.GetTradeNotification(req)
	if noti != nil {
		service.Frontend().Buy().UpdateOrder(noti.OutTradeNo)
		if string(noti.TradeStatus) == "TRADE_SUCCESS" {
			g.Logger.Debugf("success")
		}
	}
	alipay.AckNotification(rep) // 确认收到通知消息
	c.Redirect(302, "/user/order")
}

// AlipayReturn
// @Tags Pay
// @Summary 支付重定向
// @Success 200 "重定向"
// @Header 200 {string} Location "/user/order"
// @Router /alipayReturn [get]
func (a *PayApi) AlipayReturn(c *gin.Context) {
	c.Redirect(302, "/user/order")
}
