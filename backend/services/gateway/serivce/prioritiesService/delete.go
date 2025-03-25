package priorities_service

import (
	"backend/protos/gen/go/priorities"
	"context"
	"log/slog"
)

func (s *Service) Delete(ctx context.Context, req *priorities.DeletePrioritiesRequest) (*priorities.DeletePrioritiesResponse, error) {
	const op = "priorities_service.Delete"
	log := s.logger.With(slog.String("op", op), slog.Int("id", int(req.GetId())))
	log.Debug("deleting priority")

	resp, err := s.client.Delete(ctx, req)
	if err != nil {
		log.Error("failed to delete priority", "error", err)
		return nil, err
	}

	log.Info("priority deleted successfully")
	return resp, nil
}
