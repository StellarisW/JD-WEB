package config

type Cookie struct {
	MaxAge   int  `mapstructure:"max_age" json:"max_age" yaml:"max_age"`
	Secure   bool `mapstructure:"secure" json:"secure" yaml:"secure"`
	HttpOnly bool `mapstructure:"httponly" json:"httponly" yaml:"httponly"`
}
