package service

import (
	"backend/protos/gen/go/directions"
	"context"
	"log/slog"
)

func (s *Service) Get(ctx context.Context, req *directions.GetDirectionRequest) (*directions.DirectionResponse, error) {
	op := "service.Get"
	s.logger.Info("Getting direction", slog.String("op", op), slog.Any("request", req))

	direction, err := s.provider.Get(ctx, int(req.Id))
	if err != nil {
		s.logger.Error("Failed to get direction", slog.String("op", op), slog.String("error", err.Error()))
		return nil, err
	}

	s.logger.Info("Direction retrieved", slog.String("op", op), slog.Any("direction", direction))
	return direction, nil
}
