package config

type Zap struct {
	Format       string `mapstructure:"format" json:"format" yaml:"format"`
	Director     string `mapstructure:"director" json:"director"  yaml:"director"`
	EncodeLevel  string `mapstructure:"encode-level" json:"encode-Level" yaml:"encode-level"`
	EncodeCaller string `mapstructure:"encode-caller" json:"encode-caller" yaml:"encode-caller"`
}
