package priorities_service

import (
	"backend/protos/gen/go/priorities"
	"context"
	"log/slog"
)

func (s *Service) Update(ctx context.Context, req *priorities.UpdatePrioritiesRequest) (*priorities.PrioritiesResponse, error) {
	const op = "priorities_service.Update"
	log := s.logger.With(slog.String("op", op), slog.Int("id", int(req.GetId())))
	log.Debug("updating priority")

	resp, err := s.client.Update(ctx, req)
	if err != nil {
		log.Error("failed to update priority", "error", err)
		return nil, err
	}

	log.Info("priority updated successfully")
	return resp, nil
}
