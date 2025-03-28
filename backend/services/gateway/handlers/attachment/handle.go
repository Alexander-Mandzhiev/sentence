package attachment_handle

import (
	"backend/protos/gen/go/attachments"
	"backend/services/gateway/models"
	"context"
	"github.com/gin-gonic/gin"
	"io"
	"log/slog"
)

type AttachmentsService interface {
	Create(ctx context.Context, metadata *attachments.AttachmentMetadata, file io.Reader) (*models.Attachment, error)
	Get(ctx context.Context, id int32) (*models.Attachment, error)
	Delete(ctx context.Context, id int32) error
	List(ctx context.Context, limit, offset int32) ([]*models.Attachment, error)
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
