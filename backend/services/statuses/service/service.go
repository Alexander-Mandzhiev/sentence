package service

import (
	"backend/protos/gen/go/statuses"
	"context"
	"log/slog"
)

type StatusesProvider interface {
	Create(ctx context.Context, status *statuses.StatusResponse) (int, error)
	Get(ctx context.Context, id int) (*statuses.StatusResponse, error)
	Update(ctx context.Context, status *statuses.StatusResponse) error
	Delete(ctx context.Context, id int) error
	List(ctx context.Context) ([]*statuses.StatusResponse, error)
}

type Service struct {
	logger           *slog.Logger
	statusesProvider StatusesProvider
}

func New(statusesProvider StatusesProvider, logger *slog.Logger) *Service {
	op := "service.New"
	if statusesProvider == nil {
		logger.Error("Location provider is nil", slog.String("op", op))
		return nil
	}

	logger.Info("Service initialized", slog.String("op", op))
	return &Service{statusesProvider: statusesProvider, logger: logger}
}
