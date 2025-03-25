package sentences_attachments_service

import (
	"backend/protos/gen/go/sentences_attachments"
	"context"
	"log/slog"
)

func (s *Service) Delete(ctx context.Context, req *sentences_attachments.DeleteSentenceAttachmentRequest) (*sentences_attachments.DeleteSentenceAttachmentResponse, error) {
	const op = "sentences_attachments_service.Delete"
	log := s.logger.With(slog.String("op", op), slog.Int64("sentence_id", req.GetSentenceId()), slog.Int("attachment_id", int(req.GetAttachmentId())))
	log.Debug("deleting sentence-attachment link")

	resp, err := s.client.Delete(ctx, req)
	if err != nil {
		log.Error("failed to delete sentence-attachment link", "error", err)
		return nil, err
	}

	if resp.GetSuccess() {
		log.Info("sentence-attachment link deleted successfully")
	} else {
		log.Warn("sentence-attachment link deletion reported failure", "message", resp.GetMessage())
	}

	return resp, nil
}
