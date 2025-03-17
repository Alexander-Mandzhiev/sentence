package main

import (
	"backend/pkg/config"
	"backend/pkg/dbManager"
	sl "backend/pkg/logger"
	"backend/pkg/server/grpc_server"
	"backend/services/directions/handle"
	"backend/services/directions/repository"
	"backend/services/directions/service"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
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

	// Инициализация репозитория
	repo, err := repository.New(pool, logger)
	if err != nil {
		logger.Error("Failed to create repository", sl.Err(err))
		return
	}

	// Инициализация сервиса
	srv := service.New(repo, logger)

	// Инициализация gRPC сервера
	app := grpc_server.New()

	// Регистрация сервиса в gRPC сервере
	handle.Register(app.GRPCServer, srv)

	// Запуск gRPC сервера
	go func() {
		app.MustRun(logger, cfg.GRPCServer.Port)
	}()

	// Ожидание сигнала для graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	// Graceful shutdown
	logger.Info("Shutting down gracefully...")
	app.Shutdown()
	logger.Info("Service stopped")
}
