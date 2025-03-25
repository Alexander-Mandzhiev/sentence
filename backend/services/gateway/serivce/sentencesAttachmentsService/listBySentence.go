package sentences_attachments_service

import (
	"backend/protos/gen/go/sentences_attachments"
	"context"
	"log/slog"
)

func (s *Service) ListBySentence(ctx context.Context, req *sentences_attachments.ListBySentenceRequest) (*sentences_attachments.SentenceAttachmentsListResponse, error) {
	const op = "sentences_attachments_service.ListBySentence"
	log := s.logger.With(slog.String("op", op), slog.Int64("sentence_id", req.GetSentenceId()))
	log.Debug("listing attachments for sentence")

	resp, err := s.client.ListBySentence(ctx, req)
	if err != nil {
		log.Error("failed to list attachments for sentence", "error", err)
		return nil, err
	}

	log.Debug("attachments listed successfully", "count", len(resp.GetData()))
	return resp, nil
}
