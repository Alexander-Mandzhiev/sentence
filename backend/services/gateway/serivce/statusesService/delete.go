package statuses_service

import (
	"backend/protos/gen/go/statuses"
	"context"
	"log/slog"
)

func (s *Service) Delete(ctx context.Context, req *statuses.DeleteStatusRequest) (*statuses.DeleteStatusResponse, error) {
	const op = "statuses_service.Delete"
	log := s.logger.With(slog.String("op", op), slog.Int("id", int(req.GetId())))
	log.Debug("deleting status")

	resp, err := s.client.Delete(ctx, req)
	if err != nil {
		log.Error("failed to delete status", "error", err, "status_id", req.GetId())
		return nil, err
	}

	if resp.GetSuccess() {
		log.Info("status deleted successfully")
	} else {
		log.Warn("status deletion reported failure", "status_id", req.GetId())
	}

	return resp, nil
}
