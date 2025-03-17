package service

import (
	"backend/protos/gen/go/directions"
	"context"
	"log/slog"
)

func (s *Service) Update(ctx context.Context, req *directions.UpdateDirectionRequest) (*directions.DirectionResponse, error) {
	op := "service.Update"
	s.logger.Info("Updating direction", slog.String("op", op), slog.Any("request", req))

	direction := &directions.DirectionResponse{
		Id:          req.Id,
		Name:        req.Name,
		Description: req.Description,
		IsActive:    req.IsActive,
	}

	err := s.provider.Update(ctx, direction)
	if err != nil {
		s.logger.Error("Failed to update direction", slog.String("op", op), slog.String("error", err.Error()))
		return nil, err
	}

	s.logger.Info("Direction updated", slog.String("op", op), slog.Any("direction", direction))
	return direction, nil
}
