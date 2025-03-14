package departments_handle

import (
	"backend/protos/gen/go/departments"
	"context"
	"github.com/gin-gonic/gin"
	"log/slog"
)

type DepartmentService interface {
	Create(ctx context.Context, req *departments.CreateDepartmentRequest) (*departments.DepartmentResponse, error)
	Get(ctx context.Context, req *departments.GetDepartmentRequest) (*departments.DepartmentResponse, error)
	Update(ctx context.Context, req *departments.UpdateDepartmentRequest) (*departments.DepartmentResponse, error)
	Delete(ctx context.Context, req *departments.DeleteDepartmentRequest) (*departments.DeleteDepartmentResponse, error)
	List(ctx context.Context, req *departments.ListDepartmentsRequest) (*departments.DepartmentsListResponse, error)
}
type Handler struct {
	logger  *slog.Logger
	service DepartmentService
}

func New(service DepartmentService, logger *slog.Logger) *Handler {
	return &Handler{service: service, logger: logger}
}

func (h *Handler) InitRoutes(router *gin.RouterGroup) {
	deps := router.Group("/departments")
	{
		deps.POST("/", h.create)
		deps.GET("/:id", h.get)
		deps.PUT("/:id", h.update)
		deps.DELETE("/:id", h.delete)
		deps.GET("/", h.list)
	}
}
