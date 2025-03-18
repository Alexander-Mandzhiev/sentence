package service

import (
	"backend/protos/gen/go/attachments"
	"context"
	"log/slog"
)

func (s *Service) Get(ctx context.Context, req *attachments.GetAttachmentRequest) (*attachments.AttachmentResponse, error) {
	op := "service.Get"
	s.logger.Info("Getting attachment", slog.String("op", op), slog.Int64("id", int64(req.Id)))

	attachment, err := s.provider.Get(ctx, req.Id)
	if err != nil {
		s.logger.Error("Failed to get attachment", slog.String("op", op), slog.String("error", err.Error()))
		return nil, err
	}

	s.logger.Info("Attachment fetched", slog.String("op", op), slog.Any("attachment", attachment))
	return attachment, nil
}
