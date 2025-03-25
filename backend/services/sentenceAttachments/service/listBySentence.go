package service

import (
	"backend/protos/gen/go/sentences_attachments"
	"context"
	"log/slog"
)

func (s *Service) ListBySentence(ctx context.Context, req *sentences_attachments.ListBySentenceRequest) (*sentences_attachments.SentenceAttachmentsListResponse, error) {
	const op = "service.ListBySentence"
	log := s.logger.With(slog.String("op", op), slog.Int64("sentence_id", req.SentenceId))
	log.Info("listing attachments for sentence")

	links, err := s.provider.ListBySentence(ctx, req.SentenceId)
	if err != nil {
		log.Error("failed to list attachments", "error", err)
		return nil, err
	}

	log.Info("attachments listed successfully", "count", len(links))
	return &sentences_attachments.SentenceAttachmentsListResponse{Data: links}, nil
}
