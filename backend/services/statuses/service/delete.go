package service

import (
	"backend/protos/gen/go/statuses"
	"context"
	"log/slog"
)

func (s *Service) Delete(ctx context.Context, req *statuses.DeleteStatusRequest) (*statuses.DeleteStatusResponse, error) {
	op := "service.Delete"
	s.logger.Info("Deleting status", slog.String("op", op), slog.Any("request", req))

	err := s.statusesProvider.Delete(ctx, int(req.Id))
	if err != nil {
		s.logger.Error("Failed to delete status", slog.String("op", op), slog.String("error", err.Error()))
		return &statuses.DeleteStatusResponse{Success: false}, err
	}

	s.logger.Info("Status deleted", slog.String("op", op), slog.Any("id", req.Id))
	return &statuses.DeleteStatusResponse{Success: true}, nil
}
