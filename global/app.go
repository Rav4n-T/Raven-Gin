package g

import (
	"Raven-gin/config"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var Vp = new(*viper.Viper)
var Cof = new(config.Configuration)
var Log = new(zap.Logger)
var DB = new(gorm.DB)
