package service

import (
	"backend/protos/gen/go/sentences_attachments"
	"context"
	"log/slog"
)

func (s *Service) ListByAttachment(ctx context.Context, req *sentences_attachments.ListByAttachmentRequest) (*sentences_attachments.SentenceAttachmentsListResponse, error) {
	const op = "service.ListByAttachment"
	log := s.logger.With(slog.String("op", op), slog.Int64("attachment_id", int64(req.AttachmentId)))
	log.Info("listing sentences for attachment")

	links, err := s.provider.ListByAttachment(ctx, req.AttachmentId)
	if err != nil {
		log.Error("failed to list sentences", "error", err)
		return nil, err
	}

	log.Info("sentences listed successfully", "count", len(links))
	return &sentences_attachments.SentenceAttachmentsListResponse{Data: links}, nil
}
