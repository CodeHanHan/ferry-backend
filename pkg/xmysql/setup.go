package xmysql

import (
	"context"
	"time"

	gormLogger "gorm.io/gorm/logger"

	dbLogger "github.com/CodeHanHan/ferry-backend/db/logger"
	"github.com/CodeHanHan/ferry-backend/pkg/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetUp(register func(db *gorm.DB), dsn string, loggerLevel int) error {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: dbLogger.
			NewGormLogger(time.Millisecond * 500).
			LogMode(gormLogger.LogLevel(loggerLevel)),
	})
	if err != nil {
		logger.Critical(context.Background(), "failed to connect to database: %v", err)
		return err
	}

	register(db)
	return nil
}
