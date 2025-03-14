package main

import (
	sl "backend/pkg/logger"
	"backend/pkg/server/grpc_client"
	"backend/pkg/server/http_server"
	"backend/services/gateway/clientFactory"
	"backend/services/gateway/config"
	"backend/services/gateway/handlers"
	"backend/services/gateway/handlers/attachment"
	"backend/services/gateway/handlers/attachment_types"
	"backend/services/gateway/handlers/departments"
	"backend/services/gateway/handlers/directions"
	"backend/services/gateway/handlers/history"
	"backend/services/gateway/handlers/implementors"
	"backend/services/gateway/handlers/priorities"
	"backend/services/gateway/handlers/sentence_attachment"
	"backend/services/gateway/handlers/sentences"
	"backend/services/gateway/handlers/statuses"
	"backend/services/gateway/serivce"
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

	// Инициализация фабрики клиентов
	clientFactory := client_factory.NewClientFactory(grpc_client.NewGRPCClientManager())

	// Создание клиентов для всех микросервисов
	clients, err := client_factory.CreateClients(clientFactory, cfg.Services.AttachmentTypesAddr, []client_factory.ServiceType{
		client_factory.ServiceTypeAttachmentTypes,
		client_factory.ServiceTypeDirections,
		client_factory.ServiceTypeHistory,
		client_factory.ServiceTypeAttachments,
		client_factory.ServiceTypeDepartments,
		client_factory.ServiceTypeImplementors,
		client_factory.ServiceTypePriorities,
		client_factory.ServiceTypeStatuses,
		client_factory.ServiceTypeSentences,
		client_factory.ServiceTypeSentencesAttachments,
	}, logger)
	if err != nil {
		logger.Error("Failed to create clients", sl.Err(err))
		return
	}

	// Создание сервисного слоя
	svc := serivce.New(
		clients[client_factory.ServiceTypeAttachmentTypes].(client_factory.AttachmentTypesClient),
		clients[client_factory.ServiceTypeDirections].(client_factory.DirectionsClient),
		clients[client_factory.ServiceTypeHistory].(client_factory.HistoryClient),
		clients[client_factory.ServiceTypeAttachments].(client_factory.AttachmentsClient),
		clients[client_factory.ServiceTypeDepartments].(client_factory.DepartmentsClient),
		clients[client_factory.ServiceTypeImplementors].(client_factory.ImplementorsClient),
		clients[client_factory.ServiceTypePriorities].(client_factory.PrioritiesClient),
		clients[client_factory.ServiceTypeStatuses].(client_factory.StatusesClient),
		clients[client_factory.ServiceTypeSentences].(client_factory.SentencesClient),
		clients[client_factory.ServiceTypeSentencesAttachments].(client_factory.SentencesAttachmentsClient),
	)

	// Создание хэндлеров
	attachmentTypesHandler := attachment_types_handle.New(svc.AttachmentTypesClient, logger)
	directionsHandler := directions_handle.New(svc.DirectionsClient, logger)
	historyHandler := history_handle.New(svc.HistoryClient, logger)
	attachmentsHandler := attachment_handle.New(svc.AttachmentsClient, logger)
	departmentsHandler := departments_handle.New(svc.DepartmentsClient, logger)
	implementorsHandler := implementors_handle.New(svc.ImplementorsClient, logger)
	prioritiesHandler := priorities_handle.New(svc.PrioritiesClient, logger)
	statusesHandler := statuses_handle.New(svc.StatusesClient, logger)
	sentencesHandler := sentences_handle.New(svc.SentencesClient, logger)
	sentencesAttachmentsHandler := sentence_attachment_handle.New(svc.SentencesAttachmentsClient, logger)

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
