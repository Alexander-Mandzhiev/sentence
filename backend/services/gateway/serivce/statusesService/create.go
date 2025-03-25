package statuses_service

import (
	"backend/protos/gen/go/statuses"
	"context"
	"log/slog"
)

func (s *Service) Create(ctx context.Context, req *statuses.CreateStatusRequest) (*statuses.StatusResponse, error) {
	const op = "statuses_service.Create"
	log := s.logger.With(slog.String("op", op), slog.String("name", req.GetName()))
	log.Debug("creating new status")

	resp, err := s.client.Create(ctx, req)
	if err != nil {
		log.Error("failed to create status", "error", err, "name", req.GetName(), "description", req.GetDescription())
		return nil, err
	}

	log.Info("status created successfully", "id", resp.GetId(), "name", resp.GetName())
	return resp, nil
}
