package config

type Auth struct {
	Captcha Captcha `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
	Oauth   Oauth   `mapstructure:"oauth" json:"oauth" yaml:"oauth"`
	JWT     JWT     `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
}

type Captcha struct {
	KeyLong   int `mapstructure:"key-long" json:"keyLong" yaml:"key-long"`       // 验证码长度
	ImgWidth  int `mapstructure:"img-width" json:"imgWidth" yaml:"img-width"`    // 验证码宽度
	ImgHeight int `mapstructure:"img-height" json:"imgHeight" yaml:"img-height"` // 验证码高度
}

type Oauth struct {
	ClientID     string `mapstructure:"clientid" json:"clientid" yaml:"clientid"`
	ClientSecret string `mapstructure:"clientsecret" json:"clientsecret" yaml:"clientsecret"`
}

type JWT struct {
	ExpiresTime int64  `mapstructure:"expires-time" json:"expiresTime" yaml:"expires-time"` // 过期时间
	BufferTime  int64  `mapstructure:"buffer-time" json:"bufferTime" yaml:"buffer-time"`    // 缓冲时间
	Issuer      string `mapstructure:"issuer" json:"issuer" yaml:"issuer"`                  // 签发者
}
