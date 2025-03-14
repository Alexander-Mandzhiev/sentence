package implementors_handle

import (
	"backend/protos/gen/go/implementors"
	"context"
	"github.com/gin-gonic/gin"
	"log/slog"
)

type ImplementorService interface {
	Create(ctx context.Context, req *implementors.CreateImplementorRequest) (*implementors.ImplementorResponse, error)
	Get(ctx context.Context, req *implementors.GetImplementorRequest) (*implementors.ImplementorResponse, error)
	Update(ctx context.Context, req *implementors.UpdateImplementorRequest) (*implementors.ImplementorResponse, error)
	Delete(ctx context.Context, req *implementors.DeleteImplementorRequest) (*implementors.DeleteImplementorResponse, error)
	List(ctx context.Context, req *implementors.ListImplementorsRequest) (*implementors.ImplementorsListResponse, error)
}
type Handler struct {
	logger  *slog.Logger
	service ImplementorService
}

func New(service ImplementorService, logger *slog.Logger) *Handler {
	return &Handler{service: service, logger: logger}
}

func (h *Handler) InitRoutes(router *gin.RouterGroup) {
	impl := router.Group("/implementors")
	{
		impl.POST("/", h.create)
		impl.GET("/:id", h.get)
		impl.PUT("/:id", h.update)
		impl.DELETE("/:id", h.delete)
		impl.GET("/", h.list)
	}
}
