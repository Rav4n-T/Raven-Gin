package bootstrap

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"go.uber.org/zap"
	"gopkg.in/natefinch/lumberjack.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"Raven-gin/app/models"
	g "Raven-gin/global"
)

func InitializeDB() *gorm.DB {
	switch g.Cof.Database.Driver {
	case "mysql":
		return initMysqlGorm()
	default:
		return initMysqlGorm()
	}
}

func initMysqlGorm() *gorm.DB {
	dbConfig := g.Cof.Database
	if dbConfig.Database == "" {
		return nil
	}
	dsn := dbConfig.Username + ":" + dbConfig.Password + "@tcp(" + dbConfig.Host + ":" + strconv.Itoa(dbConfig.Port) + ")/" + dbConfig.Database + "?charset=" + dbConfig.Charset + "&parseTime=True&loc=Local"
	mysqlConfig := mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         191,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}

	if db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger:                                   getGormLogger(),
	}); err != nil {
		g.Log.Error("MySQL启动异常", zap.Any("err", err))
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(dbConfig.MaxIdleConns)
		sqlDB.SetMaxOpenConns(dbConfig.MaxOpenConns)
		initMysqlTables(db)
		return db
	}

}

func initMysqlTables(db *gorm.DB) {
	err := db.AutoMigrate(
		models.User{},
	)

	if err != nil {
		g.Log.Error("自动迁移失败", zap.Any("err", err))
		os.Exit(0)
	}
}

func getGormLogWriter() logger.Writer {
	var writer io.Writer

	if g.Cof.Database.EnableFileLogWriter {
		writer = &lumberjack.Logger{
			Filename:   filepath.Join(g.Cof.Log.RootDir, g.Cof.Database.LogFilename),
			MaxSize:    g.Cof.Log.MaxSize,
			MaxBackups: g.Cof.Log.MaxBackup,
			MaxAge:     g.Cof.Log.MaxAge,
			Compress:   g.Cof.Log.Compress,
		}
	} else {
		writer = os.Stdout
	}

	return log.New(writer, "\n", log.LstdFlags)
}

func getGormLogger() logger.Interface {
	var logMode logger.LogLevel
	switch g.Cof.Database.LogMode {
	case "silent":
		logMode = logger.Silent
	case "error":
		logMode = logger.Error
	case "warn":
		logMode = logger.Warn
	case "info":
		logMode = logger.Info
	}

	return logger.New(getGormLogWriter(), logger.Config{
		SlowThreshold:             200 * time.Millisecond,
		LogLevel:                  logMode,
		IgnoreRecordNotFoundError: false,
		Colorful:                  !g.Cof.Database.EnableFileLogWriter,
	})
}
