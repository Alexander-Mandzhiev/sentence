package service

import (
	"backend/protos/gen/go/departments"
	"context"
	"log/slog"
)

func (s *Service) List(ctx context.Context) (*departments.DepartmentsListResponse, error) {
	op := "service.List"
	s.logger.Info("Listing departments", slog.String("op", op))

	departmentsList, err := s.provider.List(ctx)
	if err != nil {
		s.logger.Error("Failed to list departments", slog.String("op", op), slog.String("error", err.Error()))
		return nil, err
	}

	s.logger.Info("Departments listed", slog.String("op", op), slog.Int("count", len(departmentsList)))
	return &departments.DepartmentsListResponse{Data: departmentsList}, nil
}
