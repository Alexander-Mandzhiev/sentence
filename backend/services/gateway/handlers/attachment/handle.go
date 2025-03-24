package attachment_handle

import (
	"backend/protos/gen/go/attachments"
	"context"
	"github.com/gin-gonic/gin"
	"log/slog"
)

type AttachmentsService interface {
	Create(ctx context.Context, req *attachments.CreateAttachmentRequest) (*attachments.AttachmentResponse, error)
	Get(ctx context.Context, req *attachments.GetAttachmentRequest) (*attachments.AttachmentResponse, error)
	Update(ctx context.Context, req *attachments.AttachmentResponse) (*attachments.AttachmentResponse, error)
	Delete(ctx context.Context, req *attachments.DeleteAttachmentRequest) (*attachments.DeleteAttachmentResponse, error)
	List(ctx context.Context, req *attachments.ListAttachmentsRequest) (*attachments.AttachmentsListResponse, error)
}

type Handler struct {
	logger  *slog.Logger
	service AttachmentsService
}

func New(service AttachmentsService, logger *slog.Logger) *Handler {
	return &Handler{service: service, logger: logger}
}

func (h *Handler) InitRoutes(router *gin.RouterGroup) {
	attachment := router.Group("/attachments")
	{
		attachment.POST("/", h.create)
		attachment.GET("/:id", h.get)
		attachment.PUT("/:id", h.update)
		attachment.DELETE("/:id", h.delete)
		attachment.GET("/", h.list)
	}
}
