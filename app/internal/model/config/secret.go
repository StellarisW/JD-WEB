package config

type Secret struct {
	Common  string `mapstructure:"common" json:"common" yaml:"common"`
	Private string `mapstructure:"private" json:"private" yaml:"private"`
	JWT     string `mapstructure:"jwt" json:"jwt" yaml:"jwt"` // jwt签名
}
