package departments_service

import (
	"backend/protos/gen/go/departments"
	"context"
	"log/slog"
)

func (s *Service) Create(ctx context.Context, req *departments.CreateDepartmentRequest) (*departments.DepartmentResponse, error) {
	const op = "departments_service.Create"
	log := s.logger.With(slog.String("op", op))

	log.Debug("creating department", "name", req.Name, "description", req.Description, "is_active", req.IsActive)

	resp, err := s.client.Create(ctx, req)
	if err != nil {
		log.Error("failed to create department", "error", err, "name", req.Name)
		return nil, err
	}

	log.Info("department created", "id", resp.Id, "name", resp.Name)
	return resp, nil
}
