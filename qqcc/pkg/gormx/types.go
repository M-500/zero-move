package gormx

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

type DBX struct {
	*gorm.DB
}

type Config struct {
	DSN          string
	MaxOpenConns int
	MaxIdleConns int
	MaxLifetime  int
}

type ormLog struct {
	LogLevel logger.LogLevel
}

func (l *ormLog) LogMode(level logger.LogLevel) logger.Interface {
	l.LogLevel = level
	return l
}

func (l *ormLog) Info(ctx context.Context, format string, v ...interface{}) {
	if l.LogLevel < logger.Info {
		return
	}
	logx.WithContext(ctx).Infof(format, v...)
}

func (l *ormLog) Warn(ctx context.Context, format string, v ...interface{}) {
	if l.LogLevel < logger.Warn {
		return
	}
	logx.WithContext(ctx).Infof(format, v...)
}

func (l *ormLog) Error(ctx context.Context, format string, v ...interface{}) {
	if l.LogLevel < logger.Error {
		return
	}
	logx.WithContext(ctx).Infof(format, v...)
}

// Trace 链路追踪相关的
func (l *ormLog) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	elapsed := time.Since(begin)
	sql, rows := fc()
	logx.WithContext(ctx).WithDuration(elapsed).Infof("[%.3fms] [rows:%v] %s", float64(elapsed.Nanoseconds())/1e6, rows, sql)
}
