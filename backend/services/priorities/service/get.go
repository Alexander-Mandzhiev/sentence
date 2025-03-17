package service

import (
	"backend/protos/gen/go/priorities"
	"context"
	"log/slog"
)

func (s *Service) Get(ctx context.Context, req *priorities.GetPrioritiesRequest) (*priorities.PrioritiesResponse, error) {
	op := "service.Get"
	s.logger.Info("Getting priority", slog.String("op", op), slog.Any("request", req))

	priority, err := s.provider.Get(ctx, int(req.Id))
	if err != nil {
		s.logger.Error("Failed to get priority", slog.String("op", op), slog.String("error", err.Error()))
		return nil, err
	}

	s.logger.Info("Priority retrieved", slog.String("op", op), slog.Any("priority", priority))
	return priority, nil
}
