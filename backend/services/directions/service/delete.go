package service

import (
	"backend/protos/gen/go/directions"
	"context"
	"log/slog"
)

func (s *Service) Delete(ctx context.Context, req *directions.DeleteDirectionRequest) (*directions.DeleteDirectionResponse, error) {
	op := "service.Delete"
	s.logger.Info("Deleting direction", slog.String("op", op), slog.Any("request", req))

	err := s.provider.Delete(ctx, int(req.Id))
	if err != nil {
		s.logger.Error("Failed to delete direction", slog.String("op", op), slog.String("error", err.Error()))
		return &directions.DeleteDirectionResponse{Success: false}, err
	}

	s.logger.Info("Direction deleted", slog.String("op", op), slog.Any("id", req.Id))
	return &directions.DeleteDirectionResponse{Success: true}, nil
}
