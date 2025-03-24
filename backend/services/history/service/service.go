package service

import (
	"backend/protos/gen/go/history"
	"context"
	"log/slog"
)

type HistoriesProvider interface {
	Create(ctx context.Context, req *history.CreateHistoryRequest) (*history.HistoryResponse, error)
	Get(ctx context.Context, req *history.GetHistoryRequest) (*history.HistoryResponse, error)
	ListBySentence(ctx context.Context, priority *history.ListBySentenceRequest) (*history.HistoryListResponse, error)
}

type Service struct {
	logger   *slog.Logger
	provider HistoriesProvider
}

func New(historiesProvider HistoriesProvider, logger *slog.Logger) *Service {
	op := "service.New"
	if historiesProvider == nil {
		logger.Error("Priorities provider is nil", slog.String("op", op))
		return nil
	}

	logger.Info("Service initialized", slog.String("op", op))
	return &Service{provider: historiesProvider, logger: logger}
}
