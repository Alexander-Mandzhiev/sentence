package service

import (
	"backend/protos/gen/go/implementors"
	"context"
	"golang.org/x/exp/slog"
)

func (s *Service) List(ctx context.Context) (*implementors.ImplementorsListResponse, error) {
	op := "service.List"
	s.logger.Info("Listing implementors", slog.String("op", op))

	implementorsList, err := s.provider.List(ctx)
	if err != nil {
		s.logger.Error("Failed to list implementors", slog.String("op", op), slog.String("error", err.Error()))
		return nil, err
	}

	s.logger.Info("Implementors listed", slog.String("op", op), slog.Int("count", len(implementorsList)))
	return &implementors.ImplementorsListResponse{Data: implementorsList}, nil
}
