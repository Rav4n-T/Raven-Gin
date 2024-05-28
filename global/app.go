package g

import (
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"Raven-gin/config"
)

var Vp = new(*viper.Viper)

var Cof = new(config.Configuration)

var Log = new(zap.Logger)

var DB = new(gorm.DB)

var Redis = new(redis.Client)
