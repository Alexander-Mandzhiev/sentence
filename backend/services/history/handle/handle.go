package handle

import (
	"backend/protos/gen/go/history"
	"context"
	"google.golang.org/grpc"
)

type HistoryService interface {
	Create(ctx context.Context, request *history.CreateHistoryRequest) (*history.HistoryResponse, error)
	Get(ctx context.Context, request *history.GetHistoryRequest) (*history.HistoryResponse, error)
	ListBySentence(ctx context.Context, request *history.ListBySentenceRequest) (*history.HistoryListResponse, error)
}

type serverAPI struct {
	history.UnimplementedHistoryProviderServer
	service HistoryService
}

func Register(gRPCServer *grpc.Server, service HistoryService) {
	history.RegisterHistoryProviderServer(gRPCServer, &serverAPI{service: service})
}
