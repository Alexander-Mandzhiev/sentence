package historyes_service

import (
	"backend/protos/gen/go/history"
	"context"
	"log/slog"
)

func (s *Service) Create(ctx context.Context, req *history.CreateHistoryRequest) (*history.HistoryResponse, error) {
	const op = "historyes_service.Create"
	log := s.logger.With(slog.String("op", op))
	log.Debug("creating history record", "sentence_id", req.SentenceId)

	resp, err := s.client.Create(ctx, req)
	if err != nil {
		log.Error("failed to create history record", "error", err, "sentence_id", req.SentenceId)
		return nil, err
	}

	log.Info("history record created", "id", resp.Id, "sentence_id", resp.SentenceId)
	return resp, nil
}
