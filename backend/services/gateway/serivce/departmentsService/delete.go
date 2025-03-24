package departments_service

import (
	"backend/protos/gen/go/departments"
	"context"
	"log/slog"
)

func (s *Service) Delete(ctx context.Context, req *departments.DeleteDepartmentRequest) (*departments.DeleteDepartmentResponse, error) {
	const op = "departments_service.Delete"
	log := s.logger.With(slog.String("op", op))

	log.Debug("deleting department", "id", req.Id)
	resp, err := s.client.Delete(ctx, req)
	if err != nil {
		log.Error("failed to delete department", "id", req.Id, "error", err)
		return nil, err
	}

	log.Info("department deleted", "id", req.Id, "success", resp.Success)
	return resp, nil
}
