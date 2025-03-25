package statuses_service

import (
	"backend/protos/gen/go/statuses"
	"context"
	"log/slog"
)

type StatusesClient interface {
	Create(ctx context.Context, req *statuses.CreateStatusRequest) (*statuses.StatusResponse, error)
	Get(ctx context.Context, req *statuses.GetStatusRequest) (*statuses.StatusResponse, error)
	Update(ctx context.Context, req *statuses.UpdateStatusRequest) (*statuses.StatusResponse, error)
	Delete(ctx context.Context, req *statuses.DeleteStatusRequest) (*statuses.DeleteStatusResponse, error)
	List(ctx context.Context, req *statuses.ListStatusesRequest) (*statuses.StatusListResponse, error)
}

type Service struct {
	client statuses.StatusProviderClient
	logger *slog.Logger
}

func New(client statuses.StatusProviderClient, logger *slog.Logger) *Service {
	return &Service{
		client: client,
		logger: logger,
	}
}
