package historyes_service

import (
	"backend/protos/gen/go/history"
	"context"
	"log/slog"
)

type HistoryClient interface {
	Create(ctx context.Context, req *history.CreateHistoryRequest) (*history.HistoryResponse, error)
	Get(ctx context.Context, req *history.GetHistoryRequest) (*history.HistoryResponse, error)
	ListBySentence(ctx context.Context, req *history.ListBySentenceRequest) (*history.HistoryListResponse, error)
}
type Service struct {
	client history.HistoryProviderClient
	logger *slog.Logger
}

func New(client history.HistoryProviderClient, logger *slog.Logger) *Service {
	return &Service{
		client: client,
		logger: logger,
	}
}
