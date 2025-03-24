package direction_service

import (
	"backend/protos/gen/go/directions"
	"context"
	"log/slog"
)

func (s *Service) List(ctx context.Context, req *directions.ListDirectionsRequest) (*directions.DirectionsListResponse, error) {
	const op = "direction_service.List"
	log := s.logger.With(slog.String("op", op))
	log.Debug("listing directions")

	resp, err := s.client.List(ctx, req)
	if err != nil {
		log.Error("failed to list directions", "error", err)
		return nil, err
	}

	log.Debug("directions listed", "count", len(resp.Data))
	return resp, nil
}
