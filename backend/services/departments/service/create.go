package service

import (
	"backend/protos/gen/go/departments"
	"context"
	"log/slog"
)

func (s *Service) Create(ctx context.Context, req *departments.CreateDepartmentRequest) (*departments.DepartmentResponse, error) {
	op := "service.Create"
	s.logger.Info("Creating department", slog.String("op", op), slog.Any("request", req))

	department := &departments.DepartmentResponse{
		Name:        req.Name,
		Description: req.Description,
		IsActive:    req.IsActive,
	}

	if err := s.provider.Create(ctx, department); err != nil {
		s.logger.Error("Failed to create department", slog.String("op", op), slog.String("error", err.Error()))
		return nil, err
	}

	s.logger.Info("Department created", slog.String("op", op), slog.Any("department", department))
	return department, nil
}
