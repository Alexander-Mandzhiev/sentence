package handle

import (
	"backend/protos/gen/go/sentences"
	"context"
	"google.golang.org/grpc"
)

type SentencesService interface {
	Create(ctx context.Context, request *sentences.CreateSentenceRequest) (*sentences.SentenceResponse, error)
	Get(ctx context.Context, request *sentences.GetSentenceRequest) (*sentences.SentenceResponse, error)
	Update(ctx context.Context, attachment *sentences.SentenceResponse) (*sentences.SentenceResponse, error)
	Delete(ctx context.Context, request *sentences.DeleteSentenceRequest) (*sentences.DeleteSentenceResponse, error)
	List(ctx context.Context) (*sentences.SentencesListResponse, error)
}

type serverAPI struct {
	sentences.UnimplementedSentencesProviderServer
	service SentencesService
}

func Register(gRPCServer *grpc.Server, service SentencesService) {
	sentences.RegisterSentencesProviderServer(gRPCServer, &serverAPI{service: service})
}
