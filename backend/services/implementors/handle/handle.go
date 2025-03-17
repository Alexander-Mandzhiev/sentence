package handle

import (
	"backend/protos/gen/go/implementors"
	"context"
	"google.golang.org/grpc"
)

type ImplementorsService interface {
	Create(ctx context.Context, request *implementors.CreateImplementorRequest) (*implementors.ImplementorResponse, error)
	Get(ctx context.Context, request *implementors.GetImplementorRequest) (*implementors.ImplementorResponse, error)
	Update(ctx context.Context, request *implementors.UpdateImplementorRequest) (*implementors.ImplementorResponse, error)
	Delete(ctx context.Context, request *implementors.DeleteImplementorRequest) (*implementors.DeleteImplementorResponse, error)
	List(ctx context.Context) (*implementors.ImplementorsListResponse, error)
}

type serverAPI struct {
	implementors.UnimplementedImplementorsProviderServer
	service ImplementorsService
}

func Register(gRPCServer *grpc.Server, service ImplementorsService) {
	implementors.RegisterImplementorsProviderServer(gRPCServer, &serverAPI{service: service})
}
