package implementors_service

import (
	"backend/protos/gen/go/implementors"
	"context"
	"log/slog"
)

func (s *Service) List(ctx context.Context, req *implementors.ListImplementorsRequest) (*implementors.ImplementorsListResponse, error) {
	const op = "implementors_service.List"
	log := s.logger.With(slog.String("op", op))
	log.Debug("listing implementors")

	resp, err := s.client.List(ctx, req)
	if err != nil {
		log.Error("failed to list implementors", "error", err)
		return nil, err
	}

	log.Debug("implementors listed", "count", len(resp.Data))
	return resp, nil
}
