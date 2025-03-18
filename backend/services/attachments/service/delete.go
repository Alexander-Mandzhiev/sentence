package service

import (
	"backend/protos/gen/go/attachments"
	"context"
	"log/slog"
)

func (s *Service) Delete(ctx context.Context, req *attachments.DeleteAttachmentRequest) (*attachments.DeleteAttachmentResponse, error) {
	op := "service.Delete"
	s.logger.Info("Deleting attachment", slog.String("op", op), slog.Int64("id", int64(req.Id)))

	if err := s.provider.Delete(ctx, req.Id); err != nil {
		s.logger.Error("Failed to delete attachment", slog.String("op", op), slog.String("error", err.Error()))
		return nil, err
	}

	s.logger.Info("Attachment deleted", slog.String("op", op), slog.Int64("id", int64(req.Id)))
	return &attachments.DeleteAttachmentResponse{Success: true}, nil
}
