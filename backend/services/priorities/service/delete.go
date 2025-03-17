package service

import (
	"backend/protos/gen/go/priorities"
	"context"
	"log/slog"
)

func (s *Service) Delete(ctx context.Context, req *priorities.DeletePrioritiesRequest) (*priorities.DeletePrioritiesResponse, error) {
	op := "service.Delete"
	s.logger.Info("Deleting priority", slog.String("op", op), slog.Any("request", req))

	err := s.provider.Delete(ctx, int(req.Id))
	if err != nil {
		s.logger.Error("Failed to delete priority", slog.String("op", op), slog.String("error", err.Error()))
		return &priorities.DeletePrioritiesResponse{Success: false}, err
	}

	s.logger.Info("Priority deleted", slog.String("op", op), slog.Any("id", req.Id))
	return &priorities.DeletePrioritiesResponse{Success: true}, nil
}
