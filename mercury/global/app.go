package global

import (
	"github.com/spf13/viper"
	"mercury/mercury/config"
)

type Application struct {
	ConfigViper *viper.Viper
	Config      config.GlobalConfig
}

var App = new(Application)
