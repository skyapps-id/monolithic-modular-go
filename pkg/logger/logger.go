package logger

import (
	"context"
	"log/slog"
	"os"
)

func Init(level string) {
	var slogLevel slog.Level
	switch level {
	case "debug":
		slogLevel = slog.LevelDebug
	case "warn":
		slogLevel = slog.LevelWarn
	case "error":
		slogLevel = slog.LevelError
	default:
		slogLevel = slog.LevelInfo
	}

	baseHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slogLevel,
	})
	slog.SetDefault(slog.New(newCensorHandler(baseHandler)))
}

func DebugCtx(ctx context.Context, msg string, args ...any) {
	if id := TraceIDFromCtx(ctx); id != "" {
		args = append(args, "trace_id", id)
	}
	slog.Debug(msg, args...)
}

func InfoCtx(ctx context.Context, msg string, args ...any) {
	if id := TraceIDFromCtx(ctx); id != "" {
		args = append(args, "trace_id", id)
	}
	slog.Info(msg, args...)
}

func WarnCtx(ctx context.Context, msg string, args ...any) {
	if id := TraceIDFromCtx(ctx); id != "" {
		args = append(args, "trace_id", id)
	}
	slog.Warn(msg, args...)
}

func ErrorCtx(ctx context.Context, msg string, args ...any) {
	if id := TraceIDFromCtx(ctx); id != "" {
		args = append(args, "trace_id", id)
	}
	slog.Error(msg, args...)
}

func Debug(msg string, args ...any) {
	slog.Debug(msg, args...)
}

func Info(msg string, args ...any) {
	slog.Info(msg, args...)
}

func Warn(msg string, args ...any) {
	slog.Warn(msg, args...)
}

func Error(msg string, args ...any) {
	slog.Error(msg, args...)
}
