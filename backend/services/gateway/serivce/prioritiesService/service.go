package priorities_service

import (
	"backend/protos/gen/go/priorities"
	"context"
	"log/slog"
)

type PrioritiesClient interface {
	Create(ctx context.Context, req *priorities.CreatePrioritiesRequest) (*priorities.PrioritiesResponse, error)
	Get(ctx context.Context, req *priorities.GetPrioritiesRequest) (*priorities.PrioritiesResponse, error)
	Update(ctx context.Context, req *priorities.UpdatePrioritiesRequest) (*priorities.PrioritiesResponse, error)
	Delete(ctx context.Context, req *priorities.DeletePrioritiesRequest) (*priorities.DeletePrioritiesResponse, error)
	List(ctx context.Context, req *priorities.ListPrioritiesRequest) (*priorities.PrioritiesListResponse, error)
}

type Service struct {
	client priorities.PrioritiesProviderClient
	logger *slog.Logger
}

func New(client priorities.PrioritiesProviderClient, logger *slog.Logger) *Service {
	return &Service{
		client: client,
		logger: logger,
	}
}
