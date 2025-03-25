package priorities_service

import (
	"backend/protos/gen/go/priorities"
	"context"
	"log/slog"
)

func (s *Service) Create(ctx context.Context, req *priorities.CreatePrioritiesRequest) (*priorities.PrioritiesResponse, error) {
	const op = "priorities_service.Create"
	log := s.logger.With(slog.String("op", op), slog.String("name", req.GetName()))
	log.Debug("creating priority")

	resp, err := s.client.Create(ctx, req)
	if err != nil {
		log.Error("failed to create priority", "error", err)
		return nil, err
	}

	log.Info("priority created successfully", "id", resp.GetId())
	return resp, nil
}
