package global

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"mercury/mercury/config"
)

type Application struct {
	ConfigViper *viper.Viper
	Config      config.GlobalConfig
	Log         *zap.Logger
	DB          *gorm.DB
}

var App = new(Application)
