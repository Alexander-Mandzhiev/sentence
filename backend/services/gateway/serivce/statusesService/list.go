package statuses_service

import (
	"backend/protos/gen/go/statuses"
	"context"
	"log/slog"
)

func (s *Service) List(ctx context.Context, req *statuses.ListStatusesRequest) (*statuses.StatusListResponse, error) {
	const op = "statuses_service.List"
	log := s.logger.With(slog.String("op", op))
	log.Debug("listing all statuses")

	resp, err := s.client.List(ctx, req)
	if err != nil {
		log.Error("failed to list statuses", "error", err)
		return nil, err
	}

	log.Debug("statuses listed successfully", "count", len(resp.GetData()))
	return resp, nil
}
