package bootstrap

import (
	"os"
	"path/filepath"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"

	g "Raven-Admin/global"
	"Raven-Admin/utils"
)

var (
	level   zapcore.Level
	options []zap.Option
)

func InitializeLog() *zap.Logger {
	createRootDir()
	setLogLevel()
	if g.Cof.Log.ShowLine {
		options = append(options, zap.AddCaller())
	}

	return zap.New(getZapCore(), options...)
}

func createRootDir() {
	if ok, _ := utils.PathExists(g.Cof.Log.RootDir); !ok {
		_ = os.Mkdir(g.Cof.Log.RootDir, os.ModePerm)
	}
}

func setLogLevel() {
	switch g.Cof.Log.Level {
	case "debug":
		level = zap.DebugLevel
		options = append(options, zap.AddStacktrace(level))
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
		options = append(options, zap.AddStacktrace(level))
	case "dpanic":
		level = zap.DPanicLevel
	case "panic":
		level = zap.PanicLevel
	case "fatal":
		level = zap.FatalLevel
	default:
		level = zap.InfoLevel
	}
}

func getZapCore() zapcore.Core {
	var encoder zapcore.Encoder

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(t.Format("[" + "2006/01/02 15:04:05" + "]"))
	}
	encoderConfig.EncodeLevel = func(l zapcore.Level, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(g.Cof.App.Env + "." + l.String())
	}

	if g.Cof.Log.Format == "json" {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	return zapcore.NewCore(encoder, getLogWriter(), level)
}

func getLogWriter() zapcore.WriteSyncer {
	file := &lumberjack.Logger{
		Filename:   filepath.Join(g.Cof.Log.RootDir, g.Cof.Log.Filename),
		MaxSize:    g.Cof.Log.MaxSize,
		MaxAge:     g.Cof.Log.MaxAge,
		MaxBackups: g.Cof.Log.MaxBackup,
		Compress:   g.Cof.Log.Compress,
	}

	return zapcore.AddSync(file)
}
