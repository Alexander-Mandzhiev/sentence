package handle

import (
	"backend/protos/gen/go/statuses"
	"context"
	"google.golang.org/grpc"
)

type StatusesService interface {
	Create(ctx context.Context, request *statuses.CreateStatusRequest) (*statuses.StatusResponse, error)
	Get(ctx context.Context, request *statuses.GetStatusRequest) (*statuses.StatusResponse, error)
	Update(ctx context.Context, request *statuses.UpdateStatusRequest) (*statuses.StatusResponse, error)
	Delete(ctx context.Context, request *statuses.DeleteStatusRequest) (*statuses.DeleteStatusResponse, error)
	List(ctx context.Context) (*statuses.StatusListResponse, error)
}
type serverAPI struct {
	statuses.UnimplementedStatusProviderServer
	service StatusesService
}

func Register(gRPCServer *grpc.Server, service StatusesService) {
	statuses.RegisterStatusProviderServer(gRPCServer, &serverAPI{service: service})
}
