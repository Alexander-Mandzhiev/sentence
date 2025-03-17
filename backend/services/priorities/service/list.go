package service

import (
	"backend/protos/gen/go/priorities"
	"context"
	"log/slog"
)

func (s *Service) List(ctx context.Context) (*priorities.PrioritiesListResponse, error) {
	op := "service.List"
	s.logger.Info("Listing priorities", slog.String("op", op))

	prioritiesList, err := s.provider.List(ctx)
	if err != nil {
		s.logger.Error("Failed to list priorities", slog.String("op", op), slog.String("error", err.Error()))
		return nil, err
	}

	s.logger.Info("Priorities listed", slog.String("op", op), slog.Int("count", len(prioritiesList)))
	return &priorities.PrioritiesListResponse{Data: prioritiesList}, nil
}
