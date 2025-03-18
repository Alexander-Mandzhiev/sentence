package service

import (
	"backend/protos/gen/go/attachment_types"
	"context"
	"log/slog"
)

func (s *Service) List(ctx context.Context) (*attachment_types.AttachmentTypesListResponse, error) {
	op := "service.List"
	s.logger.Info("Listing attachment types", slog.String("op", op))

	attachmentTypes, err := s.provider.List(ctx)
	if err != nil {
		s.logger.Error("Failed to list attachment types", slog.String("op", op), slog.String("error", err.Error()))
		return nil, err
	}

	s.logger.Info("Attachment types listed", slog.String("op", op), slog.Int("count", len(attachmentTypes)))
	return &attachment_types.AttachmentTypesListResponse{Data: attachmentTypes}, nil
}
