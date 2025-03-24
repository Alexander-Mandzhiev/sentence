package implementors_service

import (
	"backend/protos/gen/go/implementors"
	"context"
	"log/slog"
)

type ImplementorsClient interface {
	Create(ctx context.Context, req *implementors.CreateImplementorRequest) (*implementors.ImplementorResponse, error)
	Get(ctx context.Context, req *implementors.GetImplementorRequest) (*implementors.ImplementorResponse, error)
	Update(ctx context.Context, req *implementors.UpdateImplementorRequest) (*implementors.ImplementorResponse, error)
	Delete(ctx context.Context, req *implementors.DeleteImplementorRequest) (*implementors.DeleteImplementorResponse, error)
	List(ctx context.Context, req *implementors.ListImplementorsRequest) (*implementors.ImplementorsListResponse, error)
}

type Service struct {
	client implementors.ImplementorsProviderClient
	logger *slog.Logger
}

func New(client implementors.ImplementorsProviderClient, logger *slog.Logger) *Service {
	return &Service{
		client: client,
		logger: logger,
	}
}
