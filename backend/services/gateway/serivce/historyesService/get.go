package historyes_service

import (
	"backend/protos/gen/go/history"
	"context"
	"log/slog"
)

func (s *Service) Get(ctx context.Context, req *history.GetHistoryRequest) (*history.HistoryResponse, error) {
	const op = "historyes_service.Get"
	log := s.logger.With(slog.String("op", op))
	log.Debug("getting history record", "id", req.Id)

	resp, err := s.client.Get(ctx, req)
	if err != nil {
		log.Error("failed to get history record", "id", req.Id, "error", err)
		return nil, err
	}

	log.Debug("history record retrieved", "id", resp.Id, "sentence_id", resp.SentenceId)
	return resp, nil
}
