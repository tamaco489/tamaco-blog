package logger

import (
	"context"
	"log/slog"
)

// WithContext returns a logger with the given context
func WithContext(ctx context.Context) *slog.Logger {
	return GetLogger().With(slog.String("request_id", GetRequestID(ctx)))
}

// GetRequestID extracts request ID from context
func GetRequestID(ctx context.Context) string {
	if requestID, ok := ctx.Value("request_id").(string); ok {
		return requestID
	}
	return "unknown"
}

// Info logs an info message
func Info(msg string, args ...any) {
	GetLogger().Info(msg, args...)
}

// InfoContext logs an info message with context
func InfoContext(ctx context.Context, msg string, args ...any) {
	GetLogger().InfoContext(ctx, msg, args...)
}

// Warn logs a warning message
func Warn(msg string, args ...any) {
	GetLogger().Warn(msg, args...)
}

// WarnContext logs a warning message with context
func WarnContext(ctx context.Context, msg string, args ...any) {
	GetLogger().WarnContext(ctx, msg, args...)
}

// Error logs an error message
func Error(msg string, args ...any) {
	GetLogger().Error(msg, args...)
}

// ErrorContext logs an error message with context
func ErrorContext(ctx context.Context, msg string, args ...any) {
	GetLogger().ErrorContext(ctx, msg, args...)
}

// Debug logs a debug message
func Debug(msg string, args ...any) {
	GetLogger().Debug(msg, args...)
}

// DebugContext logs a debug message with context
func DebugContext(ctx context.Context, msg string, args ...any) {
	GetLogger().DebugContext(ctx, msg, args...)
}
