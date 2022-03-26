package utils

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	g "main/app/global"
	"time"
)

type JWT struct {
	SigningKey []byte
}

type JwtBlackList struct {
	g.Model
	Id  string
	Jwt string //`gorm:"type:text;comment:jwt"`
}

var (
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("Couldn't handle this token:")
)

func NewJWT() *JWT {
	return &JWT{
		[]byte(g.Config.Secret.JWT),
	}
}

func (j *JWT) CreateClaims(baseClaims BaseClaims) CustomClaims {
	config := g.Config.Auth.JWT
	claims := CustomClaims{
		BaseClaims: baseClaims,
		BufferTime: config.BufferTime, // 缓冲时间1天 缓冲时间内会获得新的token刷新令牌 此时一个用户会存在两个有效令牌 但是前端只留一个 另一个会丢失
		RegisteredClaims: jwt.RegisteredClaims{
			NotBefore: jwt.NewNumericDate(time.Now().Truncate(time.Second)),                                // 签名生效时间
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(config.ExpiresTime) * time.Second)), // 过期时间 7天  配置文件
			Issuer:    config.Issuer,                                                                       // 签名的发行者
		},
	}
	return claims
}

func (j *JWT) GenerateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

//func (j *JWT) CreateTokenByOldToken(oldToken string, claims CustomClaims) (string, error) {
//	v, err, _ := g.ConcurrencyControl.Do("JWT_"+oldToken, func() (interface{}, error) {
//		return j.GenerateToken(claims)
//	})
//	return v.(string), err
//}

// ParseToken 解析JWT
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid
	} else {
		return nil, TokenInvalid
	}
}

func IsBlacklist(jwt string) bool {
	val, err := g.Redis.Get(context.Background(), jwt).Result()
	if err != nil || val == "" {
		return false
	}
	return true
	// err := global.GVA_DB.Where("jwt = ?", jwt).First(&system.JwtBlacklist{}).Error
	// isNotFound := errors.Is(err, gorm.ErrRecordNotFound)
	// return !isNotFound
}

func (j *JWT) JsonInBlackList(jwtStr JwtBlackList) error {
	g.DB.Select("jwt").Create(jwtStr)
	g.Redis.Set(context.Background(), "jwt_"+jwtStr.Id, jwtStr.Jwt, 0)
	return nil
}
