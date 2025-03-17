package service

import (
	"backend/protos/gen/go/directions"
	"context"
	"log/slog"
)

type DirectionsProvider interface {
	Create(ctx context.Context, direction *directions.DirectionResponse) (int, error)
	Get(ctx context.Context, id int) (*directions.DirectionResponse, error)
	Update(ctx context.Context, direction *directions.DirectionResponse) error
	Delete(ctx context.Context, id int) error
	List(ctx context.Context) ([]*directions.DirectionResponse, error)
}

type Service struct {
	logger   *slog.Logger
	provider DirectionsProvider
}

func New(directionsProvider DirectionsProvider, logger *slog.Logger) *Service {
	op := "service.New"
	if directionsProvider == nil {
		logger.Error("Directions provider is nil", slog.String("op", op))
		return nil
	}

	logger.Info("Service initialized", slog.String("op", op))
	return &Service{provider: directionsProvider, logger: logger}
}
