package service

import (
	"backend/protos/gen/go/attachments"
	"context"
	"log/slog"
)

func (s *Service) Update(ctx context.Context, req *attachments.AttachmentResponse) (*attachments.AttachmentResponse, error) {
	op := "service.Update"
	s.logger.Info("Updating attachment", slog.String("op", op), slog.Any("request", req))

	if err := s.provider.Update(ctx, req); err != nil {
		s.logger.Error("Failed to update attachment", slog.String("op", op), slog.String("error", err.Error()))
		return nil, err
	}

	s.logger.Info("Attachment updated", slog.String("op", op), slog.Any("attachment", req))
	return req, nil
}
