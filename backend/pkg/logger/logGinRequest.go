package sl

import (
	"context"
	"github.com/gin-gonic/gin"
	"log/slog"
)

func LogGinRequest(logger *slog.Logger, params gin.LogFormatterParams, level slog.Level) {
	if logger == nil {
		return
	}

	attrs := []slog.Attr{
		slog.String("method", params.Method),
		slog.String("path", params.Path),
		slog.Int("status", params.StatusCode),
		slog.String("latency", params.Latency.String()),
		slog.String("client_ip", params.ClientIP),
	}

	logger.LogAttrs(
		context.Background(),
		level,
		"HTTP request",
		attrs...,
	)
}
