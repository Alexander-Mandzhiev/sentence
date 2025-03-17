package service

import (
	"backend/protos/gen/go/priorities"
	"context"
	"log/slog"
)

func (s *Service) Update(ctx context.Context, req *priorities.UpdatePrioritiesRequest) (*priorities.PrioritiesResponse, error) {
	op := "service.Update"
	s.logger.Info("Updating priority", slog.String("op", op), slog.Any("request", req))

	priority := &priorities.PrioritiesResponse{
		Id:          req.Id,
		Name:        req.Name,
		Description: req.Description,
	}

	err := s.provider.Update(ctx, priority)
	if err != nil {
		s.logger.Error("Failed to update priority", slog.String("op", op), slog.String("error", err.Error()))
		return nil, err
	}

	s.logger.Info("Priority updated", slog.String("op", op), slog.Any("priority", priority))
	return priority, nil
}
