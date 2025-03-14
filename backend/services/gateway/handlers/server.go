package handlers

import (
	sl "backend/pkg/logger"
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
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

type ServerAPI struct {
	logger                      *slog.Logger
	attachmentTypesHandler      *attachment_types_handle.Handler
	directionsHandler           *directions_handle.Handler
	historyHandler              *history_handle.Handler
	attachmentsHandler          *attachment_handle.Handler
	departmentsHandler          *departments_handle.Handler
	implementorsHandler         *implementors_handle.Handler
	prioritiesHandler           *priorities_handle.Handler
	statusesHandler             *statuses_handle.Handler
	sentencesHandler            *sentences_handle.Handler
	sentencesAttachmentsHandler *sentence_attachment_handle.Handler
}

func New(
	logger *slog.Logger,
	attachmentTypesHandler *attachment_types_handle.Handler,
	directionsHandler *directions_handle.Handler,
	historyHandler *history_handle.Handler,
	attachmentsHandler *attachment_handle.Handler,
	departmentsHandler *departments_handle.Handler,
	implementorsHandler *implementors_handle.Handler,
	prioritiesHandler *priorities_handle.Handler,
	statusesHandler *statuses_handle.Handler,
	sentencesHandler *sentences_handle.Handler,
	sentencesAttachmentsHandler *sentence_attachment_handle.Handler,
) *ServerAPI {
	return &ServerAPI{
		logger:                      logger,
		attachmentTypesHandler:      attachmentTypesHandler,
		directionsHandler:           directionsHandler,
		historyHandler:              historyHandler,
		attachmentsHandler:          attachmentsHandler,
		departmentsHandler:          departmentsHandler,
		implementorsHandler:         implementorsHandler,
		prioritiesHandler:           prioritiesHandler,
		statusesHandler:             statusesHandler,
		sentencesHandler:            sentencesHandler,
		sentencesAttachmentsHandler: sentencesAttachmentsHandler,
	}
}

func (h *ServerAPI) InitRouters() http.Handler {
	router := gin.Default()
	router.Use(gin.LoggerWithConfig(gin.LoggerConfig{
		Output: sl.NewLoggerWriter(h.logger, slog.LevelInfo),
	}))
	router.Use(gin.Recovery())

	api := router.Group("/api/v1")
	{
		h.attachmentTypesHandler.InitRoutes(api)
		h.directionsHandler.InitRoutes(api)
		h.historyHandler.InitRoutes(api)
		h.attachmentsHandler.InitRoutes(api)
		h.departmentsHandler.InitRoutes(api)
		h.implementorsHandler.InitRoutes(api)
		h.prioritiesHandler.InitRoutes(api)
		h.statusesHandler.InitRoutes(api)
		h.sentencesHandler.InitRoutes(api)
		h.sentencesAttachmentsHandler.InitRoutes(api)
	}

	router.GET("/healthcheck", h.healthcheck)

	return router
}

func (h *ServerAPI) healthcheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
