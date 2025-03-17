package service

import (
	"backend/protos/gen/go/statuses"
	"context"
	"log/slog"
)

func (s *Service) Create(ctx context.Context, req *statuses.CreateStatusRequest) (*statuses.StatusResponse, error) {
	op := "service.Create"
	s.logger.Info("Creating status", slog.String("op", op), slog.Any("request", req))

	status := &statuses.StatusResponse{
		Name:        req.Name,
		Description: req.Description,
	}

	id, err := s.statusesProvider.Create(ctx, status)
	if err != nil {
		s.logger.Error("Failed to create status", slog.String("op", op), slog.String("error", err.Error()))
		return nil, err
	}

	status.Id = int32(id)
	s.logger.Info("Status created", slog.String("op", op), slog.Any("status", status))
	return status, nil
}
