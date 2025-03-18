package service

import (
	"backend/protos/gen/go/departments"
	"context"
	"log/slog"
)

type DepartmentsProvider interface {
	Create(ctx context.Context, department *departments.DepartmentResponse) error
	Get(ctx context.Context, id int) (*departments.DepartmentResponse, error)
	Update(ctx context.Context, department *departments.DepartmentResponse) error
	Delete(ctx context.Context, id int) error
	List(ctx context.Context) ([]*departments.DepartmentResponse, error)
}

type Service struct {
	logger   *slog.Logger
	provider DepartmentsProvider
}

func New(departmentsProvider DepartmentsProvider, logger *slog.Logger) *Service {
	op := "service.New"
	if departmentsProvider == nil {
		logger.Error("Departments provider is nil", slog.String("op", op))
		return nil
	}

	logger.Info("Service initialized", slog.String("op", op))
	return &Service{provider: departmentsProvider, logger: logger}
}
