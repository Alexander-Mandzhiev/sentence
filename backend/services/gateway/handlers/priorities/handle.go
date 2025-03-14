package priorities_handle

import (
	"backend/protos/gen/go/priorities"
	"context"
	"github.com/gin-gonic/gin"
	"log/slog"
)

type PrioritiesService interface {
	Create(ctx context.Context, req *priorities.CreatePrioritiesRequest) (*priorities.PrioritiesResponse, error)
	Get(ctx context.Context, req *priorities.GetPrioritiesRequest) (*priorities.PrioritiesResponse, error)
	Update(ctx context.Context, req *priorities.UpdatePrioritiesRequest) (*priorities.PrioritiesResponse, error)
	Delete(ctx context.Context, req *priorities.DeletePrioritiesRequest) (*priorities.DeletePrioritiesResponse, error)
	List(ctx context.Context, req *priorities.ListPrioritiesRequest) (*priorities.PrioritiesListResponse, error)
}
type Handler struct {
	logger  *slog.Logger
	service PrioritiesService
}

func New(service PrioritiesService, logger *slog.Logger) *Handler {
	return &Handler{service: service, logger: logger}
}

func (h *Handler) InitRoutes(router *gin.RouterGroup) {
	prio := router.Group("/priorities")
	{
		prio.POST("/", h.create)
		prio.GET("/:id", h.get)
		prio.PUT("/:id", h.update)
		prio.DELETE("/:id", h.delete)
		prio.GET("/", h.list)
	}
}
