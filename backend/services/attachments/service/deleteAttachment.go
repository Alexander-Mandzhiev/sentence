package service

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
	"os"
)

func (s *Service) DeleteAttachment(ctx context.Context, id int32) error {
	const op = "Service.DeleteAttachment"
	logger := s.logger.With(slog.String("op", op), slog.Int("id", int(id)))

	if id <= 0 {
		return status.Error(codes.InvalidArgument, "ID must be positive")
	}

	filePath, err := s.provider.Delete(ctx, id)
	if err != nil {
		logger.Error("failed to delete attachment from DB", slog.String("error", err.Error()))
		return fmt.Errorf("%s: %w", op, err)
	}

	if err = os.Remove(filePath); err != nil && !os.IsNotExist(err) {
		logger.Error("failed to delete file",
			slog.String("filePath", filePath),
			slog.String("error", err.Error()))
		return fmt.Errorf("%s: failed to delete file: %w", op, err)
	}

	logger.Info("attachment deleted successfully")
	return nil
}
