package logger

import (
	"context"
	"time"

	gormlogger "gorm.io/gorm/logger"
)

func NewGormLogger(level string) gormlogger.Interface {
	var logLevel gormlogger.LogLevel
	switch level {
	case "silent":
		logLevel = gormlogger.Silent
	case "error":
		logLevel = gormlogger.Error
	case "warn":
		logLevel = gormlogger.Warn
	case "info":
		logLevel = gormlogger.Info
	default:
		logLevel = gormlogger.Warn
	}

	return &gormSlogLogger{
		level:         logLevel,
		slowThreshold: time.Second,
	}
}

type gormSlogLogger struct {
	level         gormlogger.LogLevel
	slowThreshold time.Duration
}

func (l *gormSlogLogger) LogMode(level gormlogger.LogLevel) gormlogger.Interface {
	newLogger := *l
	newLogger.level = level
	return &newLogger
}

func (l *gormSlogLogger) Info(ctx context.Context, msg string, args ...interface{}) {
	if l.level >= gormlogger.Info {
		InfoCtx(ctx, msg, attrsFromArgs(args)...)
	}
}

func (l *gormSlogLogger) Warn(ctx context.Context, msg string, args ...interface{}) {
	if l.level >= gormlogger.Warn {
		WarnCtx(ctx, msg, attrsFromArgs(args)...)
	}
}

func (l *gormSlogLogger) Error(ctx context.Context, msg string, args ...interface{}) {
	if l.level >= gormlogger.Error {
		ErrorCtx(ctx, msg, attrsFromArgs(args)...)
	}
}

func (l *gormSlogLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	if l.level <= gormlogger.Silent {
		return
	}

	elapsed := time.Since(begin)
	sql, rows := fc()

	attrs := []any{
		"sql", sql,
		"rows", rows,
		"elapsed_ms", elapsed.Milliseconds(),
	}

	if err != nil {
		attrs = append(attrs, "error", err.Error())
		ErrorCtx(ctx, "gorm query error", attrs...)
		return
	}

	if l.slowThreshold > 0 && elapsed > l.slowThreshold {
		attrs = append(attrs, "slow_threshold_ms", l.slowThreshold.Milliseconds())
		WarnCtx(ctx, "gorm slow query", attrs...)
		return
	}

	if l.level >= gormlogger.Info {
		InfoCtx(ctx, "gorm query", attrs...)
	}
}

func attrsFromArgs(args []interface{}) []any {
	attrs := make([]any, 0, len(args)*2)
	for i := 0; i < len(args)-1; i += 2 {
		attrs = append(attrs, args[i].(string), args[i+1])
	}
	return attrs
}
