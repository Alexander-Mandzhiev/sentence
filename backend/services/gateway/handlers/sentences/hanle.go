package sentences_handle

import (
	"backend/services/gateway/models"
	"context"
	"github.com/gin-gonic/gin"
	"log/slog"
)

type SentenceService interface {
	Create(ctx context.Context, req *models.CreateSentenceRequest) (*models.EnrichedSentence, error)
	Get(ctx context.Context, id int64) (*models.EnrichedSentence, error)
	Update(ctx context.Context, req *models.UpdateSentenceRequest) (*models.EnrichedSentence, error)
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context, limit, offset int32) ([]*models.EnrichedSentence, int32, error)
}

type Handler struct {
	logger  *slog.Logger
	service SentenceService
}

func New(service SentenceService, logger *slog.Logger) *Handler {
	return &Handler{service: service, logger: logger}
}

func (h *Handler) InitRoutes(router *gin.RouterGroup) {
	sent := router.Group("/sentences")
	{
		sent.POST("/", h.create)
		sent.GET("/:id", h.get)
		sent.PUT("/:id", h.update)
		sent.DELETE("/:id", h.delete)
		sent.GET("/", h.list)
	}
}
