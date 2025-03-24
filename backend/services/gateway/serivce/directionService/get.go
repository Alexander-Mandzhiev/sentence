package direction_service

import (
	"backend/protos/gen/go/directions"
	"context"
	"log/slog"
)

func (s *Service) Get(ctx context.Context, req *directions.GetDirectionRequest) (*directions.DirectionResponse, error) {
	const op = "direction_service.Get"
	log := s.logger.With(slog.String("op", op))

	log.Debug("getting direction", "id", req.Id)
	resp, err := s.client.Get(ctx, req)
	if err != nil {
		log.Error("failed to get direction", "id", req.Id, "error", err)
		return nil, err
	}

	log.Debug("direction retrieved", "id", resp.Id, "name", resp.Name)
	return resp, nil
}
