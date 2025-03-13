package sl

import (
	"log/slog"
	"os"
)

const (
	envDevelopment = "development"
	envProduction  = "production"
)

func SetupLogger(env string) *slog.Logger {
	var level slog.Level
	switch env {
	case envDevelopment:
		level = slog.LevelDebug
	case envProduction:
		level = slog.LevelInfo
	default:
		level = slog.LevelInfo
	}

	var handler slog.Handler
	if env == envDevelopment {
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: level})
	} else {
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: level})
	}

	return slog.New(handler)
}
