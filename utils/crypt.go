package utils

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	g "main/app/global"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func EncryptBySHA256(str string) string {
	vs := base64.URLEncoding.EncodeToString([]byte(str))
	timestamp := strconv.FormatInt(time.Now().UnixNano(), 10)
	h := hmac.New(sha256.New, []byte(g.Config.Secret.Common))
	fmt.Fprintf(h, "%s%s", vs, timestamp)
	sig := fmt.Sprintf("%02x", h.Sum(nil))
	crypto := strings.Join([]string{vs, timestamp, sig}, "|")
	return crypto
}

func DecryptBySHA256(str string) (string, bool) {
	parts := strings.SplitN(str, "|", 3)
	if len(parts) != 3 {
		return "", false
	}

	vs := parts[0]
	timestamp := parts[1]
	sig := parts[2]

	h := hmac.New(sha256.New, []byte(g.Config.Secret.Common))
	fmt.Fprintf(h, "%s%s", vs, timestamp)

	if fmt.Sprintf("%02x", h.Sum(nil)) != sig {
		return "", false
	}
	res, _ := base64.URLEncoding.DecodeString(vs)
	return string(res), true
}

func Md5(str string) string {
	m := md5.New()
	m.Write([]byte(str))
	return string(hex.EncodeToString(m.Sum(nil)))
}

func GetRandomNum() string {
	var str string
	for i := 0; i < 10; i++ {
		current := rand.Intn(10) //0-9   "math/rand"
		str += strconv.Itoa(current)
	}
	return str
}
