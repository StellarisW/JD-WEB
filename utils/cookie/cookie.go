package cookie

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	g "main/app/global"
	"net/http"
	"strconv"
	"strings"
	"time"
)

//写入数据的方法
func Set(c *gin.Context, key string, value interface{}) {
	bytes, _ := json.Marshal(value)
	setSecureCookie(c, key, string(bytes))
}

//获取数据的方法
func Get(c *gin.Context, key string, obj interface{}) bool {
	tempData, ok := getSecureCookie(c, g.Config.Secret.Common, key)
	if !ok {
		return false
	}
	json.Unmarshal([]byte(tempData), obj)
	return true
}

//删除数据的方法
func Remove(c *gin.Context, key string, value interface{}) {
	bytes, _ := json.Marshal(value)
	setSecureCookie(c, key, string(bytes))
}

func setSecureCookie(c *gin.Context, name, value string) {
	vs := base64.URLEncoding.EncodeToString([]byte(value))
	timestamp := strconv.FormatInt(time.Now().UnixNano(), 10)
	h := hmac.New(sha256.New, []byte(g.Config.Secret.Common))
	fmt.Fprintf(h, "%s%s", vs, timestamp)
	sig := fmt.Sprintf("%02x", h.Sum(nil))
	cookie := strings.Join([]string{vs, timestamp, sig}, "|")
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     name,
		Value:    cookie,
		MaxAge:   g.Config.Cookie.MaxAge,
		Path:     "/",
		Domain:   g.Config.App.Domain,
		SameSite: http.SameSite(1),
		Secure:   g.Config.Cookie.Secure,
		HttpOnly: g.Config.Cookie.HttpOnly,
	})
}

// GetSecureCookie Get secure cookie from request by a given key.
func getSecureCookie(c *gin.Context, Secret, key string) (string, bool) {
	val, err := c.Cookie(key)
	if val == "" || err != nil {
		return "", false
	}

	parts := strings.SplitN(val, "|", 3)

	if len(parts) != 3 {
		return "", false
	}

	vs := parts[0]
	timestamp := parts[1]
	sig := parts[2]

	h := hmac.New(sha256.New, []byte(Secret))
	fmt.Fprintf(h, "%s%s", vs, timestamp)

	if fmt.Sprintf("%02x", h.Sum(nil)) != sig {
		return "", false
	}
	res, _ := base64.URLEncoding.DecodeString(vs)
	return string(res), true
}
