package implementors_service

import (
	"backend/protos/gen/go/implementors"
	"context"
	"log/slog"
)

func (s *Service) Delete(ctx context.Context, req *implementors.DeleteImplementorRequest) (*implementors.DeleteImplementorResponse, error) {
	const op = "implementors_service.Delete"
	log := s.logger.With(slog.String("op", op))
	log.Debug("deleting implementor", "id", req.Id)

	resp, err := s.client.Delete(ctx, req)
	if err != nil {
		log.Error("failed to delete implementor", "id", req.Id, "error", err)
		return nil, err
	}

	log.Info("implementor deleted", "id", req.Id, "success", resp.Success)
	return resp, nil
}
