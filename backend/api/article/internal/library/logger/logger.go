package logger

import (
	"io"
	"log/slog"
	"os"
)

// LogLevel represents the logging level
type LogLevel string

const (
	LevelDebug LogLevel = "debug"
	LevelInfo  LogLevel = "info"
	LevelWarn  LogLevel = "warn"
	LevelError LogLevel = "error"
)

// Config represents logger configuration
type Config struct {
	Level  LogLevel `json:"level"`
	Format string   `json:"format"` // "json" or "text"
}

var defaultLogger *slog.Logger

// Init initializes the global logger with the given configuration
func Init(config Config) {
	var level slog.Level
	switch config.Level {
	case LevelDebug:
		level = slog.LevelDebug
	case LevelInfo:
		level = slog.LevelInfo
	case LevelWarn:
		level = slog.LevelWarn
	case LevelError:
		level = slog.LevelError
	default:
		level = slog.LevelInfo
	}

	var handler slog.Handler
	opts := &slog.HandlerOptions{
		Level: level,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			// Remove timestamp prefix for cleaner logs
			if a.Key == slog.TimeKey {
				a.Key = "timestamp"
			}
			return a
		},
	}

	switch config.Format {
	case "json":
		handler = slog.NewJSONHandler(os.Stdout, opts)
	default:
		handler = slog.NewTextHandler(os.Stdout, opts)
	}

	defaultLogger = slog.New(handler)
	slog.SetDefault(defaultLogger)
}

// InitDefault initializes the logger with default configuration
func InitDefault() {
	Init(Config{
		Level:  LevelInfo,
		Format: "json",
	})
}

// GetLogger returns the global logger instance
func GetLogger() *slog.Logger {
	if defaultLogger == nil {
		InitDefault()
	}
	return defaultLogger
}

// SetOutput sets the output destination for the logger
func SetOutput(w io.Writer) {
	if defaultLogger == nil {
		InitDefault()
	}

	// Create new handler with the new output
	opts := &slog.HandlerOptions{
		Level: slog.LevelInfo,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				a.Key = "timestamp"
			}
			return a
		},
	}

	newHandler := slog.NewJSONHandler(w, opts)
	defaultLogger = slog.New(newHandler)
	slog.SetDefault(defaultLogger)
}
