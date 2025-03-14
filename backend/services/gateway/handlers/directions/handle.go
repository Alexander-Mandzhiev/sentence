package directions_handle

import (
	"backend/protos/gen/go/directions"
	"context"
	"github.com/gin-gonic/gin"
	"log/slog"
)

type DirectionService interface {
	Create(ctx context.Context, req *directions.CreateDirectionRequest) (*directions.DirectionResponse, error)
	Get(ctx context.Context, req *directions.GetDirectionRequest) (*directions.DirectionResponse, error)
	Update(ctx context.Context, req *directions.UpdateDirectionRequest) (*directions.DirectionResponse, error)
	Delete(ctx context.Context, req *directions.DeleteDirectionRequest) (*directions.DeleteDirectionResponse, error)
	List(ctx context.Context, req *directions.ListDirectionsRequest) (*directions.DirectionsListResponse, error)
}

type Handler struct {
	logger *slog.Logger
	client DirectionService
}

func New(client DirectionService, logger *slog.Logger) *Handler {
	return &Handler{client: client, logger: logger}
}

func (h *Handler) InitRoutes(router *gin.RouterGroup) {
	direction := router.Group("/directions")
	{
		direction.POST("/", h.create)
		direction.GET("/:id", h.get)
		direction.PUT("/:id", h.update)
		direction.DELETE("/:id", h.delete)
		direction.GET("/", h.list)
	}
}
