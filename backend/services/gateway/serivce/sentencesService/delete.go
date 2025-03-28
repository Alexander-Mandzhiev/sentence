package sentences_service

import (
	"backend/protos/gen/go/sentences"
	"backend/protos/gen/go/sentences_attachments"
	"context"
	"log/slog"
)

func (s *Service) Delete(ctx context.Context, id int64) error {
	const op = "SentenceService.Delete"
	log := s.logger.With(slog.String("op", op), slog.Int64("id", id))

	if _, err := s.sentenceAttachService.Delete(ctx, &sentences_attachments.DeleteSentenceAttachmentRequest{
		SentenceId: id,
	}); err != nil {
		log.Warn("failed to delete sentence attachments", "error", err)
	}

	if _, err := s.sentenceClient.Delete(ctx, &sentences.DeleteSentenceRequest{Id: id}); err != nil {
		log.Error("failed to delete sentence", "error", err)
		return err
	}

	return nil
}
