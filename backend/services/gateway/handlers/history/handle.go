package history_handle

import (
	"backend/protos/gen/go/history"
	"context"
	"github.com/gin-gonic/gin"
	"log/slog"
)

type HistoryService interface {
	Create(ctx context.Context, req *history.CreateHistoryRequest) (*history.HistoryResponse, error)
	Get(ctx context.Context, req *history.GetHistoryRequest) (*history.HistoryResponse, error)
	ListBySentence(ctx context.Context, req *history.ListBySentenceRequest) (*history.HistoryListResponse, error)
}
type Handler struct {
	logger  *slog.Logger
	service HistoryService
}

func New(service HistoryService, logger *slog.Logger) *Handler {
	return &Handler{service: service, logger: logger}
}

func (h *Handler) InitRoutes(router *gin.RouterGroup) {
	ht := router.Group("/history")
	{
		ht.POST("/", h.create)
		ht.GET("/:id", h.get)
		ht.GET("/sentence/:sentence_id", h.listBySentence)
	}
}
