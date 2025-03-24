package implementors_service

import (
	"backend/protos/gen/go/implementors"
	"context"
	"log/slog"
)

func (s *Service) Get(ctx context.Context, req *implementors.GetImplementorRequest) (*implementors.ImplementorResponse, error) {
	const op = "implementors_service.Get"
	log := s.logger.With(slog.String("op", op))
	log.Debug("getting implementor", "id", req.Id)

	resp, err := s.client.Get(ctx, req)
	if err != nil {
		log.Error("failed to get implementor", "id", req.Id, "error", err)
		return nil, err
	}

	log.Debug("implementor retrieved", "id", resp.Id, "name", resp.Name)
	return resp, nil
}
