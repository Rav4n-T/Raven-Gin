package middleware

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"

	"Raven-gin/app/common/response"
	g "Raven-gin/global"
)

func CustomRecovery() gin.HandlerFunc {
	return gin.RecoveryWithWriter(
		&lumberjack.Logger{
			Filename:   g.Cof.Log.RootDir + "/" + g.Cof.Log.Filename,
			MaxSize:    g.Cof.Log.MaxSize,
			MaxBackups: g.Cof.Log.MaxBackup,
			MaxAge:     g.Cof.Log.MaxAge,
			Compress:   g.Cof.Log.Compress,
		},
		response.FailByServerError)
}
