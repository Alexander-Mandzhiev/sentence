package service

import (
	"backend/protos/gen/go/directions"
	"context"
	"log/slog"
)

func (s *Service) Create(ctx context.Context, req *directions.CreateDirectionRequest) (*directions.DirectionResponse, error) {
	op := "service.Create"
	s.logger.Info("Creating direction", slog.String("op", op), slog.Any("request", req))

	direction := &directions.DirectionResponse{
		Name:        req.Name,
		Description: req.Description,
		IsActive:    req.IsActive,
	}

	id, err := s.provider.Create(ctx, direction)
	if err != nil {
		s.logger.Error("Failed to create direction", slog.String("op", op), slog.String("error", err.Error()))
		return nil, err
	}

	direction.Id = int32(id)
	s.logger.Info("Direction created", slog.String("op", op), slog.Any("direction", direction))
	return direction, nil
}
