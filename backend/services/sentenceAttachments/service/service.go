package service

import (
	"backend/protos/gen/go/sentences_attachments"
	"context"
	"log/slog"
)

type SentencesAttachmentsProvider interface {
	Create(ctx context.Context, link *sentences_attachments.SentenceAttachmentResponse) error
	Delete(ctx context.Context, sentenceID int64, attachmentID int32) error
	ListBySentence(ctx context.Context, sentenceID int64) ([]*sentences_attachments.SentenceAttachmentResponse, error)
	ListByAttachment(ctx context.Context, attachmentID int32) ([]*sentences_attachments.SentenceAttachmentResponse, error)
}

type Service struct {
	logger   *slog.Logger
	provider SentencesAttachmentsProvider
}

func New(provider SentencesAttachmentsProvider, logger *slog.Logger) *Service {
	const op = "service.New"
	if provider == nil {
		logger.Error("SentencesAttachments provider is nil", slog.String("op", op))
		return nil
	}

	logger.Info("SentencesAttachments service initialized", slog.String("op", op))
	return &Service{
		provider: provider,
		logger:   logger,
	}
}
