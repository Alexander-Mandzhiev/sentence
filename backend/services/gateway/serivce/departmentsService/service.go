package departments_service

import (
	"backend/protos/gen/go/departments"
	"context"
	"log/slog"
)

type DepartmentsClient interface {
	Create(ctx context.Context, req *departments.CreateDepartmentRequest) (*departments.DepartmentResponse, error)
	Get(ctx context.Context, req *departments.GetDepartmentRequest) (*departments.DepartmentResponse, error)
	Update(ctx context.Context, req *departments.DepartmentResponse) (*departments.DepartmentResponse, error)
	Delete(ctx context.Context, req *departments.DeleteDepartmentRequest) (*departments.DeleteDepartmentResponse, error)
	List(ctx context.Context, req *departments.ListDepartmentsRequest) (*departments.DepartmentsListResponse, error)
}
type Service struct {
	client departments.DepartmentsProviderClient
	logger *slog.Logger
}

func New(client departments.DepartmentsProviderClient, logger *slog.Logger) *Service {
	return &Service{
		client: client,
		logger: logger,
	}
}
