package departments_service

import (
	"backend/protos/gen/go/departments"
	"context"
	"log/slog"
)

func (s *Service) List(ctx context.Context, req *departments.ListDepartmentsRequest) (*departments.DepartmentsListResponse, error) {
	const op = "departments_service.List"
	log := s.logger.With(slog.String("op", op))

	log.Debug("listing departments")
	resp, err := s.client.List(ctx, req)
	if err != nil {
		log.Error("failed to list departments", "error", err)
		return nil, err
	}

	log.Debug("departments listed", "count", len(resp.Data))
	return resp, nil
}
