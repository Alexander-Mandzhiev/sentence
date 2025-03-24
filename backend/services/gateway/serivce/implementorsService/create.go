package implementors_service

import (
	"backend/protos/gen/go/implementors"
	"context"
	"log/slog"
)

func (s *Service) Create(ctx context.Context, req *implementors.CreateImplementorRequest) (*implementors.ImplementorResponse, error) {
	const op = "implementors_service.Create"
	log := s.logger.With(slog.String("op", op))
	log.Debug("creating implementor", "name", req.Name, "email", req.Email)

	resp, err := s.client.Create(ctx, req)
	if err != nil {
		log.Error("failed to create implementor", "error", err, "name", req.Name)
		return nil, err
	}

	log.Info("implementor created", "id", resp.Id, "name", resp.Name)
	return resp, nil
}
