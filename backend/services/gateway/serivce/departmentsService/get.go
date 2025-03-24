package departments_service

import (
	"backend/protos/gen/go/departments"
	"context"
	"log/slog"
)

func (s *Service) Get(ctx context.Context, req *departments.GetDepartmentRequest) (*departments.DepartmentResponse, error) {
	const op = "departments_service.Get"
	log := s.logger.With(slog.String("op", op))

	log.Debug("getting department", "id", req.Id)
	resp, err := s.client.Get(ctx, req)
	if err != nil {
		log.Error("failed to get department", "id", req.Id, "error", err)
		return nil, err
	}

	log.Debug("department retrieved", "id", resp.Id, "name", resp.Name)
	return resp, nil
}
