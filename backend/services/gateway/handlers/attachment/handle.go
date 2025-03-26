package attachment_handle

import (
	"backend/protos/gen/go/attachments"
	"context"
	"github.com/gin-gonic/gin"
	"io"
	"log/slog"
)

type AttachmentsService interface {
	CreateAttachment(ctx context.Context, metadata *attachments.AttachmentMetadata, file io.Reader) (*attachments.AttachmentResponse, error)
	GetAttachment(ctx context.Context, id int32) (*attachments.AttachmentResponse, error)
	DeleteAttachment(ctx context.Context, id int32) error
	ListAttachments(ctx context.Context, limit, offset int32) (*attachments.AttachmentsListResponse, error)
	DownloadFile(ctx context.Context, id int32) (io.ReadCloser, *attachments.FileMetadata, error)
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
		attachment.DELETE("/:id", h.delete)
		attachment.GET("/", h.list)
		attachment.GET("/:id/download", h.download)
	}
}
