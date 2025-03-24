package sentences_handle

import (
	"backend/protos/gen/go/sentences"
	"context"
	"github.com/gin-gonic/gin"
	"log/slog"
)

type SentenceService interface {
	Create(ctx context.Context, req *sentences.CreateSentenceRequest) (*sentences.SentenceResponse, error)
	Get(ctx context.Context, req *sentences.GetSentenceRequest) (*sentences.SentenceResponse, error)
	Update(ctx context.Context, req *sentences.SentenceResponse) (*sentences.SentenceResponse, error)
	Delete(ctx context.Context, req *sentences.DeleteSentenceRequest) (*sentences.DeleteSentenceResponse, error)
	List(ctx context.Context, req *sentences.ListSentencesRequest) (*sentences.SentencesListResponse, error)
}
type Handler struct {
	logger  *slog.Logger
	service SentenceService
}

func New(service SentenceService, logger *slog.Logger) *Handler {
	return &Handler{service: service, logger: logger}
}

func (h *Handler) InitRoutes(router *gin.RouterGroup) {
	sentences := router.Group("/sentences")
	{
		sentences.POST("/", h.create)
		sentences.GET("/:id", h.get)
		sentences.PUT("/:id", h.update)
		sentences.DELETE("/:id", h.delete)
		sentences.GET("/", h.list)
	}
}
