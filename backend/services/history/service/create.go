package service

import (
	"backend/protos/gen/go/history"
	"context"
	"fmt"
	"log/slog"
)

func (s *Service) Create(ctx context.Context, req *history.CreateHistoryRequest) (*history.HistoryResponse, error) {
	const op = "service.HistoryService.Create"

	if req == nil {
		s.logger.Error("request is nil", slog.String("op", op))
		return nil, fmt.Errorf("%s: request is nil", op)
	}

	if req.SentenceId == 0 {
		s.logger.Error("sentence_id is required", slog.String("op", op))
		return nil, fmt.Errorf("%s: sentence_id is required", op)
	}

	if req.StatusId == 0 {
		s.logger.Error("status_id is required", slog.String("op", op))
		return nil, fmt.Errorf("%s: status_id is required", op)
	}

	return s.provider.Create(ctx, req)
}
