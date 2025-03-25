package priorities_service

import (
	"backend/protos/gen/go/priorities"
	"context"
	"log/slog"
)

func (s *Service) Get(ctx context.Context, req *priorities.GetPrioritiesRequest) (*priorities.PrioritiesResponse, error) {
	const op = "priorities_service.Get"
	log := s.logger.With(slog.String("op", op), slog.Int("id", int(req.GetId())))
	log.Debug("getting priority")

	resp, err := s.client.Get(ctx, req)
	if err != nil {
		log.Error("failed to get priority", "error", err)
		return nil, err
	}

	log.Debug("priority retrieved successfully")
	return resp, nil
}
