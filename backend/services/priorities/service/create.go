package service

import (
	"backend/protos/gen/go/priorities"
	"context"
	"log/slog"
)

func (s *Service) Create(ctx context.Context, req *priorities.CreatePrioritiesRequest) (*priorities.PrioritiesResponse, error) {
	op := "service.Create"
	s.logger.Info("Creating priority", slog.String("op", op), slog.Any("request", req))

	priority := &priorities.PrioritiesResponse{
		Name:        req.Name,
		Description: req.Description,
	}

	id, err := s.provider.Create(ctx, priority)
	if err != nil {
		s.logger.Error("Failed to create priority", slog.String("op", op), slog.String("error", err.Error()))
		return nil, err
	}

	priority.Id = int32(id)
	s.logger.Info("Priority created", slog.String("op", op), slog.Any("priority", priority))
	return priority, nil
}
