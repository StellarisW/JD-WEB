package utils

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	g "main/app/global"
	"main/utils/cookie"
)

type Claims struct {
	ID string
	jwt.RegisteredClaims
}

type CustomClaims struct {
	BaseClaims
	BufferTime int64
	jwt.RegisteredClaims
}

type BaseClaims struct {
	//UUID        uuid.UUID
	Id         int
	Phone      string
	LastIp     string
	Email      string
	Status     int
	UpdateTime string
	CreateTime string
	//Username    string
	//NickName    string
	//AuthorityId string
}

func GetClaims(c *gin.Context) (*CustomClaims, error) {
	//token := c.Request.Header.Get("x-token")
	var token string
	ok := cookie.Get(c, "x-token", &token)
	//token, err := c.Cookie("x-token")
	if !ok {
		err := errors.New("get token by cookie failed")
		g.Logger.Errorf("从cookie获取token失败, err: %v\n", err)
		return nil, err
	}
	j := NewJWT()
	claims, err := j.ParseToken(token)
	if err != nil {
		err := errors.New("parse token failed")
		g.Logger.Errorf("从Gin的Context中获取从jwt解析信息失败, 请检查请求头是否存在x-token且claims是否为规定结构")
		return nil, err
	}
	return claims, nil
}

// 从Gin的Context中获取从jwt解析出来的用户ID
func GetUserID(c *gin.Context) int {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return 0
		} else {
			return cl.BaseClaims.Id
		}
	} else {
		waitUse := claims.(*CustomClaims)
		return waitUse.BaseClaims.Id
	}
}

func GetUserPhone(c *gin.Context) string {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return ""
		} else {
			return cl.BaseClaims.Phone
		}
	} else {
		waitUse := claims.(*CustomClaims)
		return waitUse.BaseClaims.Phone
	}
}

//// 从Gin的Context中获取从jwt解析出来的用户UUID
//func GetUserUuid(c *gin.Context) uuid.UUID {
//	if claims, exists := c.Get("claims"); !exists {
//		if cl, err := GetClaims(c); err != nil {
//			return uuid.UUID{}
//		} else {
//			return cl.UUID
//		}
//	} else {
//		waitUse := claims.(*systemReq.CustomClaims)
//		return waitUse.UUID
//	}
//}

//// 从Gin的Context中获取从jwt解析出来的用户角色id
//func GetUserAuthorityId(c *gin.Context) string {
//	if claims, exists := c.Get("claims"); !exists {
//		if cl, err := GetClaims(c); err != nil {
//			return ""
//		} else {
//			return cl.AuthorityId
//		}
//	} else {
//		waitUse := claims.(*systemReq.CustomClaims)
//		return waitUse.AuthorityId
//	}
//}

// GetUserInfo 从Gin的Context中获取从jwt解析出来的用户角色id
func GetUserInfo(c *gin.Context) (*BaseClaims, error) {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return nil, err
		} else {
			return &cl.BaseClaims, nil
		}
	} else {
		waitUse := claims.(*CustomClaims)
		return &waitUse.BaseClaims, nil
	}
}
