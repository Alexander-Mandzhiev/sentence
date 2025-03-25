package sentences_attachments_service

import (
	"backend/protos/gen/go/sentences_attachments"
	"context"
	"log/slog"
)

func (s *Service) Create(ctx context.Context, req *sentences_attachments.CreateSentenceAttachmentRequest) (*sentences_attachments.SentenceAttachmentResponse, error) {
	const op = "sentences_attachments_service.Create"
	log := s.logger.With(slog.String("op", op), slog.Int64("sentence_id", req.GetSentenceId()), slog.Int("attachment_id", int(req.GetAttachmentId())))
	log.Debug("creating sentence-attachment link")

	resp, err := s.client.Create(ctx, req)
	if err != nil {
		log.Error("failed to create sentence-attachment link", "error", err)
		return nil, err
	}

	log.Info("sentence-attachment link created successfully", "sentence_id", resp.GetSentenceId(), "attachment_id", resp.GetAttachmentId())
	return resp, nil
}
