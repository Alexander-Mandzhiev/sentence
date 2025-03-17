package service

import (
	"backend/protos/gen/go/directions"
	"context"
	"log/slog"
)

func (s *Service) List(ctx context.Context) (*directions.DirectionsListResponse, error) {
	op := "service.List"
	s.logger.Info("Listing directions", slog.String("op", op))

	directionsList, err := s.provider.List(ctx)
	if err != nil {
		s.logger.Error("Failed to list directions", slog.String("op", op), slog.String("error", err.Error()))
		return nil, err
	}

	s.logger.Info("Directions listed", slog.String("op", op), slog.Int("count", len(directionsList)))
	return &directions.DirectionsListResponse{Data: directionsList}, nil
}
