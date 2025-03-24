package direction_service

import (
	"backend/protos/gen/go/directions"
	"context"
	"log/slog"
)

func (s *Service) Create(ctx context.Context, req *directions.CreateDirectionRequest) (*directions.DirectionResponse, error) {
	const op = "direction_service.Create"
	log := s.logger.With(slog.String("op", op))

	log.Debug("creating direction", "name", req.Name, "description", req.Description, "is_active", req.IsActive)

	resp, err := s.client.Create(ctx, req)
	if err != nil {
		log.Error("failed to create direction", "error", err, "name", req.Name)
		return nil, err
	}

	log.Info("direction created", "id", resp.Id, "name", resp.Name)
	return resp, nil
}
