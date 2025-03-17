package service

import (
	"backend/protos/gen/go/statuses"
	"context"
	"log/slog"
)

func (s *Service) Update(ctx context.Context, req *statuses.UpdateStatusRequest) (*statuses.StatusResponse, error) {
	op := "service.Update"
	s.logger.Info("Updating status", slog.String("op", op), slog.Any("request", req))

	status := &statuses.StatusResponse{
		Id:          req.Id,
		Name:        req.Name,
		Description: req.Description,
	}

	err := s.statusesProvider.Update(ctx, status)
	if err != nil {
		s.logger.Error("Failed to update status", slog.String("op", op), slog.String("error", err.Error()))
		return nil, err
	}

	s.logger.Info("Status updated", slog.String("op", op), slog.Any("status", status))
	return status, nil
}
