package service

import (
	"backend/protos/gen/go/departments"
	"context"
	"log/slog"
)

func (s *Service) Delete(ctx context.Context, req *departments.DeleteDepartmentRequest) (*departments.DeleteDepartmentResponse, error) {
	op := "service.Delete"
	s.logger.Info("Deleting department", slog.String("op", op), slog.Any("request", req))

	err := s.provider.Delete(ctx, int(req.Id))
	if err != nil {
		s.logger.Error("Failed to delete department", slog.String("op", op), slog.String("error", err.Error()))
		return &departments.DeleteDepartmentResponse{Success: false}, err
	}

	s.logger.Info("Department deleted", slog.String("op", op), slog.Any("id", req.Id))
	return &departments.DeleteDepartmentResponse{Success: true}, nil
}
