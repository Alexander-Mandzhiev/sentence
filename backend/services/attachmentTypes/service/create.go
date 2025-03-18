package service

import (
	"backend/protos/gen/go/attachment_types"
	"context"
	"log/slog"
)

func (s *Service) Create(ctx context.Context, req *attachment_types.CreateAttachmentTypeRequest) (*attachment_types.AttachmentTypeResponse, error) {
	op := "service.Create"
	s.logger.Info("Creating attachment type", slog.String("op", op), slog.Any("request", req))

	attachmentType := &attachment_types.AttachmentTypeResponse{
		Name:        req.Name,
		Description: req.Description,
	}

	id, err := s.provider.Create(ctx, attachmentType)
	if err != nil {
		s.logger.Error("Failed to create attachment type", slog.String("op", op), slog.String("error", err.Error()))
		return nil, err
	}

	attachmentType.Id = id

	s.logger.Info("Attachment type created", slog.String("op", op), slog.Any("attachmentType", attachmentType))
	return attachmentType, nil
}
