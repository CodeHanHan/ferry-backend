package logger

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/CodeHanHan/ferry-backend/pkg/logger"
	"gorm.io/gorm/utils"

	gorm_logger "gorm.io/gorm/logger"
	gorm_utils "gorm.io/gorm/utils"
)

type gormLogger struct {
	SlowThreshold time.Duration
}

func NewGormLogger(slowThreshold time.Duration) gorm_logger.Interface {
	return &gormLogger{
		SlowThreshold: slowThreshold,
	}
}

func (l *gormLogger) LogMode(level gorm_logger.LogLevel) gorm_logger.Interface {
	newlogger := *l
	return &newlogger
}

func (l gormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	logger.Info(ctx, msg, append([]interface{}{gorm_utils.FileWithLineNum()}, data...)...)
}

func (l gormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	logger.Warn(ctx, msg, append([]interface{}{gorm_utils.FileWithLineNum()}, data...)...)
}

func (l gormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	logger.Error(ctx, msg, append([]interface{}{gorm_utils.FileWithLineNum()}, data...)...)
}

func (l gormLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	elapsed := time.Since(begin)
	switch {
	case err != nil && !errors.Is(err, gorm_logger.ErrRecordNotFound):
		sql, rows := fc()
		if rows == -1 {
			logger.Debug(ctx, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			logger.Debug(ctx, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	case elapsed > l.SlowThreshold && l.SlowThreshold != 0:
		sql, rows := fc()
		slowLog := fmt.Sprintf("SLOW SQL >= %v", l.SlowThreshold)
		if rows == -1 {
			logger.Warn(ctx, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			logger.Warn(ctx, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	default:
		sql, rows := fc()
		if rows == -1 {
			logger.Debug(ctx, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			logger.Debug(ctx, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	}
}
