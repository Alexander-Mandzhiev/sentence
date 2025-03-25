package priorities_service

import (
	"backend/protos/gen/go/priorities"
	"context"
	"log/slog"
)

func (s *Service) List(ctx context.Context, req *priorities.ListPrioritiesRequest) (*priorities.PrioritiesListResponse, error) {
	const op = "priorities_service.List"
	log := s.logger.With(slog.String("op", op))
	log.Debug("listing priorities")

	resp, err := s.client.List(ctx, req)
	if err != nil {
		log.Error("failed to list priorities", "error", err)
		return nil, err
	}

	log.Debug("priorities listed successfully", "count", len(resp.Data))
	return resp, nil
}
