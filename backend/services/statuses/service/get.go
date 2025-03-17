package service

import (
	"backend/protos/gen/go/statuses"
	"context"
	"log/slog"
)

func (s *Service) Get(ctx context.Context, req *statuses.GetStatusRequest) (*statuses.StatusResponse, error) {
	op := "service.Get"
	s.logger.Info("Getting status", slog.String("op", op), slog.Any("request", req))

	status, err := s.statusesProvider.Get(ctx, int(req.Id))
	if err != nil {
		s.logger.Error("Failed to get status", slog.String("op", op), slog.String("error", err.Error()))
		return nil, err
	}

	s.logger.Info("Status retrieved", slog.String("op", op), slog.Any("status", status))
	return status, nil
}
