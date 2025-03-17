package service

import (
	"backend/protos/gen/go/implementors"
	"context"
	"log/slog"
)

type ImplementorsProvider interface {
	Create(ctx context.Context, implementor *implementors.ImplementorResponse) (int, error)
	Get(ctx context.Context, id int) (*implementors.ImplementorResponse, error)
	Update(ctx context.Context, implementor *implementors.ImplementorResponse) error
	Delete(ctx context.Context, id int) error
	List(ctx context.Context) ([]*implementors.ImplementorResponse, error)
}

type Service struct {
	logger   *slog.Logger
	provider ImplementorsProvider
}

func New(implementorsProvider ImplementorsProvider, logger *slog.Logger) *Service {
	op := "service.New"
	if implementorsProvider == nil {
		logger.Error("Implementors provider is nil", slog.String("op", op))
		return nil
	}

	logger.Info("Service initialized", slog.String("op", op))
	return &Service{provider: implementorsProvider, logger: logger}
}
