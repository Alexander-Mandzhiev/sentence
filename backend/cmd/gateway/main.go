package main

import (
	sl "backend/pkg/logger"
	"backend/pkg/server/grpc_client"
	"backend/pkg/server/http_server"
	"backend/services/gateway/clientFactory"
	"backend/services/gateway/config"
	"backend/services/gateway/handlers"
	service "backend/services/gateway/serivce"
	"golang.org/x/net/context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Инициализация конфигурации
	cfg := config.MustLoad()

	// Инициализация логгера
	logger := sl.SetupLogger(cfg.Env)
	logger.Info("Starting service direction", slog.String("address", cfg.HTTPServer.Address), slog.Int("port", cfg.HTTPServer.Port))
	logger.Debug("Debug messages are enabled")

	// Создание менеджера GRPC соединений
	grpcManager := grpc_client.NewGRPCClientManager(logger)

	// Инициализация фабрики клиентов с кэшированием
	clientFactory := client_factory.NewClientFactory(grpcManager, logger)

	// Создание сервисного слоя
	svc, err := service.New(clientFactory, cfg.Services, logger)
	if err != nil {
		logger.Error("failed to initialize services", sl.Err(err))
		os.Exit(1)
	}

	// Создание ServerAPI
	serverAPI := handlers.New(
		logger,
		attachmentTypesHandler,
		directionsHandler,
		historyHandler,
		attachmentsHandler,
		departmentsHandler,
		implementorsHandler,
		prioritiesHandler,
		statusesHandler,
		sentencesHandler,
		sentencesAttachmentsHandler,
	)

	// Создание HTTP-сервера
	apiServer := http_server.New(serverAPI, logger)

	// Запуск HTTP-сервера в отдельной горутине
	go func() {
		if err = apiServer.Start(cfg.Frontend.Addr, cfg.HTTPServer); err != nil {
			logger.Error("HTTP server error", sl.Err(err))
		}
	}()

	// Ожидание сигналов для завершения работы
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// Корректное завершение работы
	logger.Info("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = apiServer.Shutdown(ctx); err != nil {
		logger.Error("Server shutdown error", sl.Err(err))
	}

	logger.Info("Server stopped")
}
