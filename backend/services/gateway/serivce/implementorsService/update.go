package implementors_service

import (
	"backend/protos/gen/go/implementors"
	"context"
	"log/slog"
)

func (s *Service) Update(ctx context.Context, req *implementors.UpdateImplementorRequest) (*implementors.ImplementorResponse, error) {
	const op = "implementors_service.Update"
	log := s.logger.With(slog.String("op", op))
	log.Debug("updating implementor", "id", req.Id, "name", req.Name, "email", req.Email)

	resp, err := s.client.Update(ctx, req)
	if err != nil {
		log.Error("failed to update implementor", "id", req.Id, "error", err)
		return nil, err
	}

	log.Info("implementor updated", "id", resp.Id, "name", resp.Name)
	return resp, nil
}
