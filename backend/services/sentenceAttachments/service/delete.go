package service

import (
	"backend/protos/gen/go/sentences_attachments"
	"context"
	"log/slog"
)

func (s *Service) Delete(ctx context.Context, req *sentences_attachments.DeleteSentenceAttachmentRequest) (*sentences_attachments.DeleteSentenceAttachmentResponse, error) {
	const op = "service.Delete"
	log := s.logger.With(slog.String("op", op), slog.Int64("sentence_id", req.SentenceId), slog.Int64("attachment_id", int64(req.AttachmentId)))
	log.Info("deleting sentence-attachment link")

	if err := s.provider.Delete(ctx, req.SentenceId, req.AttachmentId); err != nil {
		log.Error("failed to delete link", "error", err)
		return &sentences_attachments.DeleteSentenceAttachmentResponse{Success: false, Message: err.Error()}, err
	}

	log.Info("sentence-attachment link deleted successfully")
	return &sentences_attachments.DeleteSentenceAttachmentResponse{Success: true, Message: "link deleted successfully"}, nil
}
