package direction_service

import (
	"backend/protos/gen/go/directions"
	"context"
	"log/slog"
)

type DirectionsClient interface {
	Create(ctx context.Context, req *directions.CreateDirectionRequest) (*directions.DirectionResponse, error)
	Get(ctx context.Context, req *directions.GetDirectionRequest) (*directions.DirectionResponse, error)
	Update(ctx context.Context, req *directions.UpdateDirectionRequest) (*directions.DirectionResponse, error)
	Delete(ctx context.Context, req *directions.DeleteDirectionRequest) (*directions.DeleteDirectionResponse, error)
	List(ctx context.Context, req *directions.ListDirectionsRequest) (*directions.DirectionsListResponse, error)
}
type Service struct {
	client directions.DirectionsProviderClient
	logger *slog.Logger
}

func New(client directions.DirectionsProviderClient, logger *slog.Logger) *Service {
	return &Service{
		client: client,
		logger: logger,
	}
}
