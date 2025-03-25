package statuses_service

import (
	"backend/protos/gen/go/statuses"
	"context"
	"log/slog"
)

func (s *Service) Get(ctx context.Context, req *statuses.GetStatusRequest) (*statuses.StatusResponse, error) {
	const op = "statuses_service.Get"
	log := s.logger.With(slog.String("op", op), slog.Int("id", int(req.GetId())))
	log.Debug("getting status")

	resp, err := s.client.Get(ctx, req)
	if err != nil {
		log.Error("failed to get status", "error", err, "status_id", req.GetId())
		return nil, err
	}

	log.Debug("status retrieved successfully", "name", resp.GetName())
	return resp, nil
}
