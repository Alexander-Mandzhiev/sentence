package main

import (
	"backend/pkg/config"
	"backend/pkg/dbManager"
	sl "backend/pkg/logger"
	"log/slog"
)

func main() {
	// Инициализация конфигурации
	cfg := config.Initialize()

	// Инициализация логгера
	logger := sl.SetupLogger(cfg.Env)
	logger.Info("Starting service direction", slog.String("address", cfg.GRPCServer.Address), slog.Int("port", cfg.GRPCServer.Port))
	logger.Debug("Debug messages are enabled")

	// Открытие подключения к PostgreSQL
	pool, err := dbManager.OpenPostgreSQLConnection(cfg.DBConfig.PostgreSQL, logger)
	if err != nil {
		logger.Error("Failed to open PostgreSQL connection", sl.Err(err))
		return
	}
	defer func() {
		if err = dbManager.ClosePostgreSQLConnection(pool, logger); err != nil {
			logger.Error("Failed to close PostgreSQL connection", sl.Err(err))
		}
	}()
}
