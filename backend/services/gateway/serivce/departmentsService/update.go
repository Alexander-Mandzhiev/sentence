package departments_service

import (
	"backend/protos/gen/go/departments"
	"context"
	"log/slog"
)

func (s *Service) Update(ctx context.Context, req *departments.DepartmentResponse) (*departments.DepartmentResponse, error) {
	const op = "departments_service.Update"
	log := s.logger.With(slog.String("op", op))

	log.Debug("updating department", "id", req.Id, "name", req.Name)

	resp, err := s.client.Update(ctx, req)
	if err != nil {
		log.Error("failed to update department", "id", req.Id, "error", err)
		return nil, err
	}

	log.Info("department updated", "id", resp.Id, "name", resp.Name)
	return resp, nil
}
