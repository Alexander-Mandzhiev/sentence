package sl

import (
	"bytes"
	"io"
	"log/slog"
)

type LoggerWriter struct {
	logger *slog.Logger
	level  slog.Level
}

func (w *LoggerWriter) Write(p []byte) (n int, err error) {
	message := string(bytes.TrimSpace(p))
	if message == "" {
		return len(p), nil
	}

	switch w.level {
	case slog.LevelDebug:
		w.logger.Debug(message)
	case slog.LevelInfo:
		w.logger.Info(message)
	case slog.LevelWarn:
		w.logger.Warn(message)
	case slog.LevelError:
		w.logger.Error(message)
	default:
		w.logger.Info(message)
	}

	return len(p), nil
}

func NewLoggerWriter(logger *slog.Logger, level slog.Level) io.Writer {
	return &LoggerWriter{logger: logger, level: level}
}
