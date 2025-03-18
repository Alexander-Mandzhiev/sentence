package service

import (
	"backend/protos/gen/go/attachment_types"
	"context"
	"log/slog"
)

func (s *Service) Update(ctx context.Context, req *attachment_types.AttachmentTypeResponse) (*attachment_types.AttachmentTypeResponse, error) {
	op := "service.Update"
	s.logger.Info("Updating attachment type", slog.String("op", op), slog.Any("request", req))

	if err := s.provider.Update(ctx, req); err != nil {
		s.logger.Error("Failed to update attachment type", slog.String("op", op), slog.String("error", err.Error()))
		return nil, err
	}

	s.logger.Info("Attachment type updated", slog.String("op", op), slog.Any("attachmentType", req))
	return req, nil
}
