package utils

import (
	"fmt"
	"io/ioutil"
	"time"
)

func GetUnix() int64 {
	fmt.Println(time.Now().Unix())
	return time.Now().Unix()
}

// GetUnixNano 获取时间戳Nano时间
func GetUnixNano() int64 {
	return time.Now().UnixNano()
}

// Mul 乘法的函数
func Mul(price float64, num int) float64 {
	return price * float64(num)
}

func SendMsg(str string) {
	// 短信验证码需要到相关网站申请
	// 目前先固定一个值
	ioutil.WriteFile("test_send.txt", []byte(str), 06666)
}

func GenerateOrderId() string {
	template := "200601021504"
	return time.Now().Format(template) + GetRandomNum()
}
