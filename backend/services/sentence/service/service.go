package service

import (
	"backend/protos/gen/go/sentences"
	"context"
	"log/slog"
)

type SentencesProvider interface {
	Create(ctx context.Context, sentence *sentences.SentenceResponse) error
	Get(ctx context.Context, id int64) (*sentences.SentenceResponse, error)
	Update(ctx context.Context, sentence *sentences.SentenceResponse) error
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context, limit, offset int32) ([]*sentences.SentenceResponse, int32, error)
}
type Service struct {
	logger   *slog.Logger
	provider SentencesProvider
}

func New(provider SentencesProvider, logger *slog.Logger) *Service {
	op := "service.New"
	if provider == nil {
		logger.Error("Sentences provider is nil", slog.String("op", op))
		return nil
	}

	if logger == nil {
		logger = slog.Default()
	}

	logger.Info("Service initialized", slog.String("op", op))
	return &Service{provider: provider, logger: logger}
}
