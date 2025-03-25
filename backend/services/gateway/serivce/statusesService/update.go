package statuses_service

import (
	"backend/protos/gen/go/statuses"
	"context"
	"log/slog"
)

func (s *Service) Update(ctx context.Context, req *statuses.UpdateStatusRequest) (*statuses.StatusResponse, error) {
	const op = "statuses_service.Update"
	log := s.logger.With(slog.String("op", op), slog.Int("id", int(req.GetId())))
	log.Debug("updating status", "new_name", req.GetName(), "new_description", req.GetDescription())

	resp, err := s.client.Update(ctx, req)
	if err != nil {
		log.Error("failed to update status", "error", err, "status_id", req.GetId())
		return nil, err
	}

	log.Info("status updated successfully", "id", resp.GetId(), "name", resp.GetName())
	return resp, nil
}
