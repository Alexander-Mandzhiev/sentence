package service

import (
	"backend/protos/gen/go/history"
	"context"
	"fmt"
	"log/slog"
)

func (s *Service) Get(ctx context.Context, req *history.GetHistoryRequest) (*history.HistoryResponse, error) {
	const op = "service.HistoryService.Get"

	if req == nil {
		s.logger.Error("request is nil", slog.String("op", op))
		return nil, fmt.Errorf("%s: request is nil", op)
	}

	if req.Id == 0 {
		s.logger.Error("id is required", slog.String("op", op))
		return nil, fmt.Errorf("%s: id is required", op)
	}

	return s.provider.Get(ctx, req)
}
