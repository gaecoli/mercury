package global

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"mercury/mercury/config"
)

type Application struct {
	ConfigViper *viper.Viper
	Config      config.GlobalConfig
	Log         *zap.Logger
}

var App = new(Application)
