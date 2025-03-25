package sentences_attachments_service

import (
	"backend/protos/gen/go/sentences_attachments"
	"context"
	"log/slog"
)

func (s *Service) ListByAttachment(ctx context.Context, req *sentences_attachments.ListByAttachmentRequest) (*sentences_attachments.SentenceAttachmentsListResponse, error) {
	const op = "sentences_attachments_service.ListByAttachment"
	log := s.logger.With(slog.String("op", op), slog.Int("attachment_id", int(req.GetAttachmentId())))
	log.Debug("listing sentences for attachment")

	resp, err := s.client.ListByAttachment(ctx, req)
	if err != nil {
		log.Error("failed to list sentences for attachment", "error", err)
		return nil, err
	}

	log.Debug("sentences listed successfully", "count", len(resp.GetData()))
	return resp, nil
}
