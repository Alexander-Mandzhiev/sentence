package sentences_attachments_service

import (
	"backend/protos/gen/go/sentences_attachments"
	"context"
	"log/slog"
)

type SentencesAttachmentsClient interface {
	Create(ctx context.Context, req *sentences_attachments.CreateSentenceAttachmentRequest) (*sentences_attachments.SentenceAttachmentResponse, error)
	Delete(ctx context.Context, req *sentences_attachments.DeleteSentenceAttachmentRequest) (*sentences_attachments.DeleteSentenceAttachmentResponse, error)
	ListBySentence(ctx context.Context, req *sentences_attachments.ListBySentenceRequest) (*sentences_attachments.SentenceAttachmentsListResponse, error)
	ListByAttachment(ctx context.Context, req *sentences_attachments.ListByAttachmentRequest) (*sentences_attachments.SentenceAttachmentsListResponse, error)
}

type Service struct {
	client sentences_attachments.SentencesAttachmentsProviderClient
	logger *slog.Logger
}

func New(client sentences_attachments.SentencesAttachmentsProviderClient, logger *slog.Logger) *Service {
	return &Service{
		client: client,
		logger: logger,
	}
}
