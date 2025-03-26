package main

import (
	sl "backend/pkg/logger"
	"backend/pkg/server/grpc_client"
	"backend/pkg/server/http_server"
	"backend/services/gateway/clientFactory"
	"backend/services/gateway/config"
	"backend/services/gateway/handlers"
	"backend/services/gateway/handlers/attachment"
	"backend/services/gateway/handlers/attachmentTypes"
	"backend/services/gateway/handlers/departments"
	"backend/services/gateway/handlers/directions"
	"backend/services/gateway/handlers/history"
	"backend/services/gateway/handlers/implementors"
	"backend/services/gateway/handlers/priorities"
	"backend/services/gateway/handlers/sentence_attachment"
	"backend/services/gateway/handlers/sentences"
	"backend/services/gateway/handlers/statuses"
	"backend/services/gateway/serivce/attachmentService"
	"backend/services/gateway/serivce/attachmentTypesService"
	"backend/services/gateway/serivce/departmentsService"
	"backend/services/gateway/serivce/directionService"
	"backend/services/gateway/serivce/historyesService"
	"backend/services/gateway/serivce/implementorsService"
	"backend/services/gateway/serivce/prioritiesService"
	"backend/services/gateway/serivce/sentencesAttachmentsService"
	"backend/services/gateway/serivce/sentencesService"
	"backend/services/gateway/serivce/statusesService"
	"golang.org/x/net/context"
	"log/slog"
	"net/http"
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

	grpcManager := grpc_client.NewGRPCClientManager(logger)
	defer func() {
		if err := grpcManager.CloseAll(); err != nil {
			logger.Error("Failed to close GRPC connections", sl.Err(err))
		}
	}()

	clientFactory := client_factory.New(grpcManager, logger)
	defer func() {
		if err := clientFactory.Close(); err != nil {
			logger.Error("Failed to close clients", sl.Err(err))
		}
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Инициализация всех клиентов
	attachmentsClient, err := clientFactory.GetAttachmentsClient(ctx, cfg.Services.AttachmentsAddr)
	if err != nil {
		logger.Error("Failed to create attachments client", sl.Err(err))
		os.Exit(1)
	}

	attachmentTypesClient, err := clientFactory.GetAttachmentTypesClient(ctx, cfg.Services.AttachmentTypesAddr)
	if err != nil {
		logger.Error("Failed to create attachment types client", sl.Err(err))
		os.Exit(1)
	}

	departmentsClient, err := clientFactory.GetDepartmentsClient(ctx, cfg.Services.DepartmentsAddr)
	if err != nil {
		logger.Error("Failed to create departments client", sl.Err(err))
		os.Exit(1)
	}

	directionsClient, err := clientFactory.GetDirectionsClient(ctx, cfg.Services.DirectionsAddr)
	if err != nil {
		logger.Error("Failed to create directions client", sl.Err(err))
		os.Exit(1)
	}

	historyClient, err := clientFactory.GetHistoryClient(ctx, cfg.Services.HistoryAddr)
	if err != nil {
		logger.Error("Failed to create history client", sl.Err(err))
		os.Exit(1)
	}

	implementorsClient, err := clientFactory.GetImplementorsClient(ctx, cfg.Services.ImplementorsAddr)
	if err != nil {
		logger.Error("Failed to create implementors client", sl.Err(err))
		os.Exit(1)
	}

	prioritiesClient, err := clientFactory.GetPrioritiesClient(ctx, cfg.Services.PrioritiesAddr)
	if err != nil {
		logger.Error("Failed to create priorities client", sl.Err(err))
		os.Exit(1)
	}

	sentencesAttachmentsClient, err := clientFactory.GetSentencesAttachmentsClient(ctx, cfg.Services.SentencesAttachmentsAddr)
	if err != nil {
		logger.Error("Failed to create sentences attachments client", sl.Err(err))
		os.Exit(1)
	}

	sentencesClient, err := clientFactory.GetSentencesClient(ctx, cfg.Services.SentencesAddr)
	if err != nil {
		logger.Error("Failed to create sentences client", sl.Err(err))
		os.Exit(1)
	}

	statusesClient, err := clientFactory.GetStatusesClient(ctx, cfg.Services.StatusesAddr)
	if err != nil {
		logger.Error("Failed to create statuses client", sl.Err(err))
		os.Exit(1)
	}

	// Инициализация сервисов
	attachmentSvc := attachment_service.New(attachmentsClient, logger)
	attachmentTypesSvc := attachment_types_service.New(attachmentTypesClient, logger)
	departmentSvc := departments_service.New(departmentsClient, logger)
	directionSvc := direction_service.New(directionsClient, logger)
	historySvc := historyes_service.New(historyClient, logger)
	implementorSvc := implementors_service.New(implementorsClient, logger)
	prioritySvc := priorities_service.New(prioritiesClient, logger)
	sentencesAttachmentsSvc := sentences_attachments_service.New(sentencesAttachmentsClient, logger)
	sentencesSvc := sentences_service.New(sentencesClient, logger)
	statusesSvc := statuses_service.New(statusesClient, logger)

	// Инициализация HTTP сервера
	serverAPI := handlers.New(logger, "/app/media")
	serverAPI.RegisterHandlers(
		attachment_handle.New(attachmentSvc, logger),
		attachment_types_handle.New(attachmentTypesSvc, logger),
		departments_handle.New(departmentSvc, logger),
		directions_handle.New(directionSvc, logger),
		history_handle.New(historySvc, logger),
		implementors_handle.New(implementorSvc, logger),
		priorities_handle.New(prioritySvc, logger),
		sentence_attachment_handle.New(sentencesAttachmentsSvc, logger),
		sentences_handle.New(sentencesSvc, logger),
		statuses_handle.New(statusesSvc, logger),
	)

	server := http_server.New(serverAPI, logger)

	go func() {
		if err = server.Start(cfg.Frontend.Addr, cfg.HTTPServer); err != nil && err != http.ErrServerClosed {
			logger.Error("HTTP server error", sl.Err(err))
			os.Exit(1)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info("Shutting down server...")

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer shutdownCancel()

	if err = server.Shutdown(shutdownCtx); err != nil {
		logger.Error("Server shutdown error", sl.Err(err))
	}

	logger.Info("Server exited gracefully")
}
