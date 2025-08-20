// Package logger provides logging utilities for the application.
package logger

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GinLogFormatter returns a log formatter function for Gin middleware
func GinLogFormatter(param gin.LogFormatterParams) string {
	logger := GetLogger()

	switch {
	case param.StatusCode >= http.StatusInternalServerError:
		logger.ErrorContext(param.Request.Context(), param.ErrorMessage,
			"status", param.StatusCode,
			"method", param.Method,
			"path", param.Path,
			"ip", param.ClientIP,
			"latency_ms", param.Latency.Milliseconds(),
			"user_agent", param.Request.UserAgent(),
			"host", param.Request.Host,
		)

	case param.StatusCode >= http.StatusBadRequest:
		logger.WarnContext(param.Request.Context(), param.ErrorMessage,
			"status", param.StatusCode,
			"method", param.Method,
			"path", param.Path,
			"ip", param.ClientIP,
			"latency_ms", param.Latency.Milliseconds(),
			"user_agent", param.Request.UserAgent(),
			"host", param.Request.Host,
		)

	default:
		logger.InfoContext(param.Request.Context(), "access",
			"status", param.StatusCode,
			"method", param.Method,
			"path", param.Path,
			"ip", param.ClientIP,
			"latency_ms", param.Latency.Milliseconds(),
			"user_agent", param.Request.UserAgent(),
			"host", param.Request.Host,
		)
	}
	return ""
}
