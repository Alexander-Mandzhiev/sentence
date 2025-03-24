package attachment_types

import (
	"backend/protos/gen/go/attachment_types"
	"context"
	"github.com/gin-gonic/gin"
	"log/slog"
)

type AttachmentTypesService interface {
	Create(ctx context.Context, req *attachment_types.CreateAttachmentTypeRequest) (*attachment_types.AttachmentTypeResponse, error)
	Get(ctx context.Context, req *attachment_types.GetAttachmentTypeRequest) (*attachment_types.AttachmentTypeResponse, error)
	Update(ctx context.Context, req *attachment_types.AttachmentTypeResponse) (*attachment_types.AttachmentTypeResponse, error)
	Delete(ctx context.Context, req *attachment_types.DeleteAttachmentTypeRequest) (*attachment_types.DeleteAttachmentTypeResponse, error)
	List(ctx context.Context, req *attachment_types.ListAttachmentTypesRequest) (*attachment_types.AttachmentTypesListResponse, error)
}

type Handler struct {
	logger  *slog.Logger
	service AttachmentTypesService
}

func New(service AttachmentTypesService, logger *slog.Logger) *Handler {
	return &Handler{service: service, logger: logger}
}

func (h *Handler) InitRoutes(router *gin.RouterGroup) {
	attachmentTypes := router.Group("/attachment-types")
	{
		attachmentTypes.POST("/", h.create)
		attachmentTypes.GET("/:id", h.get)
		attachmentTypes.PUT("/:id", h.update)
		attachmentTypes.DELETE("/:id", h.delete)
		attachmentTypes.GET("/", h.list)
	}
}
