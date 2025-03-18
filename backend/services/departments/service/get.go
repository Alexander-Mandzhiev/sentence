package service

import (
	"backend/protos/gen/go/departments"
	"context"
	"log/slog"
)

func (s *Service) Get(ctx context.Context, req *departments.GetDepartmentRequest) (*departments.DepartmentResponse, error) {
	op := "service.Get"
	s.logger.Info("Getting department", slog.String("op", op), slog.Any("request", req))

	department, err := s.provider.Get(ctx, int(req.Id))
	if err != nil {
		s.logger.Error("Failed to get department", slog.String("op", op), slog.String("error", err.Error()))
		return nil, err
	}

	s.logger.Info("Department retrieved", slog.String("op", op), slog.Any("department", department))
	return department, nil
}
