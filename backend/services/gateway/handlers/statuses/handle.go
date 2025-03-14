package statuses_handle

import (
	"backend/protos/gen/go/statuses"
	"context"
	"github.com/gin-gonic/gin"
	"log/slog"
)

type StatusesService interface {
	Create(ctx context.Context, req *statuses.CreateStatusRequest) (*statuses.StatusResponse, error)
	Get(ctx context.Context, req *statuses.GetStatusRequest) (*statuses.StatusResponse, error)
	Update(ctx context.Context, req *statuses.UpdateStatusRequest) (*statuses.StatusResponse, error)
	Delete(ctx context.Context, req *statuses.DeleteStatusRequest) (*statuses.DeleteStatusResponse, error)
	List(ctx context.Context, req *statuses.ListStatusesRequest) (*statuses.StatusListResponse, error)
}
type Handler struct {
	logger  *slog.Logger
	service StatusesService
}

func New(service StatusesService, logger *slog.Logger) *Handler {
	return &Handler{service: service, logger: logger}
}

func (h *Handler) InitRoutes(router *gin.RouterGroup) {
	status := router.Group("/statuses")
	{
		status.POST("/", h.create)
		status.GET("/:id", h.get)
		status.PUT("/:id", h.update)
		status.DELETE("/:id", h.delete)
		status.GET("/", h.list)
	}
}
