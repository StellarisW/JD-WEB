package config

// `mapstructure:"" json:"" yaml:""`

type Get struct {
	App    App    `mapstructure:"app" json:"app" yaml:"app"`
	Server Server `mapstructure:"server" json:"server" yaml:"server"`
	Cookie Cookie `mapstructure:"cookie" json:"cookie" yaml:"cookie"`

	Secret Secret `mapstructure:"secret" json:"secret" yaml:"secret"`

	// SQL
	Mysql Mysql `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Redis Redis `mapstructure:"redis" json:"redis" yaml:"redis"`
	//DBList []DB  `mapstructure:"db-list" json:"db-list" yaml:"db-list"`

	Auth Auth `mapstructure:"auth" json:"auth" yaml:"auth"`

	Zap    Zap    `mapstructure:"zap" json:"zap" yaml:"zap"`
	Email  Email  `mapstructure:"email" json:"email" yaml:"email"`
	Casbin Casbin `mapstructure:"casbin" json:"casbin" yaml:"casbin"`
	// auto
	AutoCode Autocode `mapstructure:"autoCode" json:"autoCode" yaml:"autoCode"`

	// oss
	Local      Local      `mapstructure:"local" json:"local" yaml:"local"`
	Qiniu      Qiniu      `mapstructure:"qiniu" json:"qiniu" yaml:"qiniu"`
	AliyunOSS  AliyunOSS  `mapstructure:"aliyun-oss" json:"aliyunOSS" yaml:"aliyun-oss"`
	HuaWeiObs  HuaWeiObs  `mapstructure:"hua-wei-obs" json:"huaWeiObs" yaml:"hua-wei-obs"`
	TencentCOS TencentCOS `mapstructure:"tencent-cos" json:"tencentCOS" yaml:"tencent-cos"`
	AwsS3      AwsS3      `mapstructure:"aws-s3" json:"awsS3" yaml:"aws-s3"`

	Excel Excel `mapstructure:"excel" json:"excel" yaml:"excel"`
	Timer Timer `mapstructure:"timer" json:"timer" yaml:"timer"`

	// 跨域配置
	Cors CORS `mapstructure:"cors" json:"cors" yaml:"cors"`
}
