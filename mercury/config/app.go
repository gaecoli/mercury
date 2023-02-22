package config

type App struct {
	Env     string `mapstructure:"env" json:"env" yaml:"env"`
	Port    string `mapstructure:"port" json:"port" yaml:"port"`
	AppName string `mapstructure:"app_name" json:"app_name" yaml:"appName"`
	AppUrl  string `mapstructure:"app_url" json:"app_url" yaml:"appUrl"`
}
