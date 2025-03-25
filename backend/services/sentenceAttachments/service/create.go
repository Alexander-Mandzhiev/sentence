package service

import (
	"backend/protos/gen/go/sentences_attachments"
	"context"
	"log/slog"
)

func (s *Service) Create(ctx context.Context, req *sentences_attachments.CreateSentenceAttachmentRequest) (*sentences_attachments.SentenceAttachmentResponse, error) {
	const op = "service.Create"
	log := s.logger.With(slog.String("op", op), slog.Int64("sentence_id", req.SentenceId), slog.Int64("attachment_id", int64(req.AttachmentId)))
	log.Info("creating sentence-attachment link")

	link := &sentences_attachments.SentenceAttachmentResponse{
		SentenceId:   req.SentenceId,
		AttachmentId: req.AttachmentId,
	}

	if err := s.provider.Create(ctx, link); err != nil {
		log.Error("failed to create link", "error", err)
		return nil, err
	}

	log.Info("sentence-attachment link created successfully")
	return link, nil
}
