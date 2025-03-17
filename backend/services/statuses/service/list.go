package service

import (
	"backend/protos/gen/go/statuses"
	"context"
	"log/slog"
)

func (s *Service) List(ctx context.Context) (*statuses.StatusListResponse, error) {
	op := "service.List"
	s.logger.Info("Listing statuses", slog.String("op", op))

	statusesList, err := s.statusesProvider.List(ctx)
	if err != nil {
		s.logger.Error("Failed to list statuses", slog.String("op", op), slog.String("error", err.Error()))
		return nil, err
	}

	s.logger.Info("Statuses listed", slog.String("op", op), slog.Int("count", len(statusesList)))
	return &statuses.StatusListResponse{Data: statusesList}, nil
}
