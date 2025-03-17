package handle

import (
	"backend/protos/gen/go/directions"
	"context"
	"google.golang.org/grpc"
)

type DirectionsService interface {
	Create(ctx context.Context, request *directions.CreateDirectionRequest) (*directions.DirectionResponse, error)
	Get(ctx context.Context, request *directions.GetDirectionRequest) (*directions.DirectionResponse, error)
	Update(ctx context.Context, request *directions.UpdateDirectionRequest) (*directions.DirectionResponse, error)
	Delete(ctx context.Context, request *directions.DeleteDirectionRequest) (*directions.DeleteDirectionResponse, error)
	List(ctx context.Context) (*directions.DirectionsListResponse, error)
}

type serverAPI struct {
	directions.UnimplementedDirectionsProviderServer
	service DirectionsService
}

func Register(gRPCServer *grpc.Server, service DirectionsService) {
	directions.RegisterDirectionsProviderServer(gRPCServer, &serverAPI{service: service})
}
