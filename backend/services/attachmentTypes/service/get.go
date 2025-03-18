package service

import (
	"backend/protos/gen/go/attachment_types"
	"context"
	"log/slog"
)

func (s *Service) Get(ctx context.Context, req *attachment_types.GetAttachmentTypeRequest) (*attachment_types.AttachmentTypeResponse, error) {
	op := "service.Get"
	s.logger.Info("Getting attachment type", slog.String("op", op), slog.Int64("id", int64(req.Id)))

	attachmentType, err := s.provider.Get(ctx, req.Id)
	if err != nil {
		s.logger.Error("Failed to get attachment type", slog.String("op", op), slog.String("error", err.Error()))
		return nil, err
	}

	s.logger.Info("Attachment type fetched", slog.String("op", op), slog.Any("attachmentType", attachmentType))
	return attachmentType, nil
}
