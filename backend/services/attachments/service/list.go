package service

import (
	"backend/protos/gen/go/attachments"
	"context"
	"log/slog"
)

func (s *Service) List(ctx context.Context) (*attachments.AttachmentsListResponse, error) {
	op := "service.List"
	s.logger.Info("Listing attachments", slog.String("op", op))

	attachmentsList, err := s.provider.List(ctx)
	if err != nil {
		s.logger.Error("Failed to list attachments", slog.String("op", op), slog.String("error", err.Error()))
		return nil, err
	}

	s.logger.Info("Attachments listed", slog.String("op", op), slog.Int("count", len(attachmentsList)))
	return &attachments.AttachmentsListResponse{Data: attachmentsList}, nil
}
