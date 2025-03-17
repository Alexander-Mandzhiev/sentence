package handle

import (
	"backend/protos/gen/go/priorities"
	"context"
	"google.golang.org/grpc"
)

type PrioritiesService interface {
	Create(ctx context.Context, request *priorities.CreatePrioritiesRequest) (*priorities.PrioritiesResponse, error)
	Get(ctx context.Context, request *priorities.GetPrioritiesRequest) (*priorities.PrioritiesResponse, error)
	Update(ctx context.Context, request *priorities.UpdatePrioritiesRequest) (*priorities.PrioritiesResponse, error)
	Delete(ctx context.Context, request *priorities.DeletePrioritiesRequest) (*priorities.DeletePrioritiesResponse, error)
	List(ctx context.Context) (*priorities.PrioritiesListResponse, error)
}

type serverAPI struct {
	priorities.UnimplementedPrioritiesProviderServer
	service PrioritiesService
}

func Register(gRPCServer *grpc.Server, service PrioritiesService) {
	priorities.RegisterPrioritiesProviderServer(gRPCServer, &serverAPI{service: service})
}
