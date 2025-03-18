package handle

import (
	"backend/protos/gen/go/departments"
	"context"
	"google.golang.org/grpc"
)

type DepartmentsService interface {
	Create(ctx context.Context, request *departments.CreateDepartmentRequest) (*departments.DepartmentResponse, error)
	Get(ctx context.Context, request *departments.GetDepartmentRequest) (*departments.DepartmentResponse, error)
	Update(ctx context.Context, department *departments.DepartmentResponse) (*departments.DepartmentResponse, error)
	Delete(ctx context.Context, request *departments.DeleteDepartmentRequest) (*departments.DeleteDepartmentResponse, error)
	List(ctx context.Context) (*departments.DepartmentsListResponse, error)
}

type serverAPI struct {
	departments.UnimplementedDepartmentsProviderServer
	service DepartmentsService
}

func Register(gRPCServer *grpc.Server, service DepartmentsService) {
	departments.RegisterDepartmentsProviderServer(gRPCServer, &serverAPI{service: service})
}
