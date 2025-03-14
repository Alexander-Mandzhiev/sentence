package sentence_attachment_handle

import (
	"backend/protos/gen/go/sentences_attachments"
	"context"
	"github.com/gin-gonic/gin"
	"log/slog"
)

type SentencesAttachmentsService interface {
	Create(ctx context.Context, req *sentences_attachments.CreateSentenceAttachmentRequest) (*sentences_attachments.SentenceAttachmentResponse, error)
	Delete(ctx context.Context, req *sentences_attachments.DeleteSentenceAttachmentRequest) (*sentences_attachments.DeleteSentenceAttachmentResponse, error)
	ListBySentence(ctx context.Context, req *sentences_attachments.ListBySentenceRequest) (*sentences_attachments.SentenceAttachmentsListResponse, error)
	ListByAttachment(ctx context.Context, req *sentences_attachments.ListByAttachmentRequest) (*sentences_attachments.SentenceAttachmentsListResponse, error)
}

type Handler struct {
	logger  *slog.Logger
	service SentencesAttachmentsService
}

func New(service SentencesAttachmentsService, logger *slog.Logger) *Handler {
	return &Handler{service: service, logger: logger}
}

func (h *Handler) InitRoutes(router *gin.RouterGroup) {
	attachments := router.Group("/sentences-attachments")
	{
		attachments.POST("/", h.create)
		attachments.DELETE("/", h.delete)
		attachments.GET("/sentence/:sentence_id", h.listBySentence)
		attachments.GET("/attachment/:attachment_id", h.listByAttachment)
	}
}
