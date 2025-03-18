package service

import (
	"backend/protos/gen/go/attachment_types"
	"context"
	"log/slog"
)

func (s *Service) Delete(ctx context.Context, req *attachment_types.DeleteAttachmentTypeRequest) (*attachment_types.DeleteAttachmentTypeResponse, error) {
	op := "service.Delete"
	s.logger.Info("Deleting attachment type", slog.String("op", op), slog.Int64("id", int64(req.Id)))

	if err := s.provider.Delete(ctx, req.Id); err != nil {
		s.logger.Error("Failed to delete attachment type", slog.String("op", op), slog.String("error", err.Error()))
		return &attachment_types.DeleteAttachmentTypeResponse{Success: false}, err
	}

	s.logger.Info("Attachment type deleted", slog.String("op", op), slog.Int64("id", int64(req.Id)))
	return &attachment_types.DeleteAttachmentTypeResponse{Success: true}, nil
}
