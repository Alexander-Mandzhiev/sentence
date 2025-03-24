package direction_service

import (
	"backend/protos/gen/go/directions"
	"context"
	"log/slog"
)

func (s *Service) Delete(ctx context.Context, req *directions.DeleteDirectionRequest) (*directions.DeleteDirectionResponse, error) {
	const op = "direction_service.Delete"
	log := s.logger.With(slog.String("op", op))
	log.Debug("deleting direction", "id", req.Id)

	resp, err := s.client.Delete(ctx, req)
	if err != nil {
		log.Error("failed to delete direction", "id", req.Id, "error", err)
		return nil, err
	}

	log.Info("direction deleted", "id", req.Id, "success", resp.Success)
	return resp, nil
}
