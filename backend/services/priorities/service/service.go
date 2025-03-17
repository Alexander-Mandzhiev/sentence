package service

import (
	"backend/protos/gen/go/priorities"
	"context"
	"log/slog"
)

type PrioritiesProvider interface {
	Create(ctx context.Context, priority *priorities.PrioritiesResponse) (int, error)
	Get(ctx context.Context, id int) (*priorities.PrioritiesResponse, error)
	Update(ctx context.Context, priority *priorities.PrioritiesResponse) error
	Delete(ctx context.Context, id int) error
	List(ctx context.Context) ([]*priorities.PrioritiesResponse, error)
}

type Service struct {
	logger   *slog.Logger
	provider PrioritiesProvider
}

func New(prioritiesProvider PrioritiesProvider, logger *slog.Logger) *Service {
	op := "service.New"
	if prioritiesProvider == nil {
		logger.Error("Priorities provider is nil", slog.String("op", op))
		return nil
	}

	logger.Info("Service initialized", slog.String("op", op))
	return &Service{provider: prioritiesProvider, logger: logger}
}
