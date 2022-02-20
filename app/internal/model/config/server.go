package config

import "time"

type Server struct {
	Env          string        `mapstructure:"env" json:"env" yaml:"env"`
	Mode         string        `mapstructure:"mode" json:"mode" yaml:"mode"`
	Port         string        `mapstructure:"port" json:"port" yaml:"port"`
	ReadTimeOut  time.Duration `mapstructure:"read-timeout" json:"read-timeout" yaml:"read-timeout"`
	WriteTimeout time.Duration `mapstructure:"write-timeout" json:"write-timeout" yaml:"write-timeout"`
	//DbType        string `mapstructure:"db-type" json:"dbType" yaml:"db-type"`
	//OssType       string `mapstructure:"oss-type" json:"ossType" yaml:"oss-type"`
	//UseMultipoint bool   `mapstructure:"use-multipoint" json:"useMultipoint" yaml:"use-multipoint"`
	//LimitCountIP  int    `mapstructure:"iplimit-count" json:"iplimitCount" yaml:"iplimit-count"`
	//LimitTimeIP   int    `mapstructure:"iplimit-time" json:"iplimitTime" yaml:"iplimit-time"`
}

func (s *Server) Addr() string {
	return ":" + s.Port
}
