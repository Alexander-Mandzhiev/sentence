package direction_service

import (
	"backend/protos/gen/go/directions"
	"context"
	"log/slog"
)

func (s *Service) Update(ctx context.Context, req *directions.UpdateDirectionRequest) (*directions.DirectionResponse, error) {
	const op = "direction_service.Update"
	log := s.logger.With(slog.String("op", op))
	log.Debug("updating direction", "id", req.Id, "name", req.Name)

	resp, err := s.client.Update(ctx, req)
	if err != nil {
		log.Error("failed to update direction", "id", req.Id, "error", err)
		return nil, err
	}

	log.Info("direction updated", "id", resp.Id, "name", resp.Name)
	return resp, nil
}
