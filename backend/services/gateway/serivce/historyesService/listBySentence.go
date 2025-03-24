package historyes_service

import (
	"backend/protos/gen/go/history"
	"context"
	"log/slog"
)

func (s *Service) ListBySentence(ctx context.Context, req *history.ListBySentenceRequest) (*history.HistoryListResponse, error) {
	const op = "historyes_service.ListBySentence"
	log := s.logger.With(slog.String("op", op))
	log.Debug("listing history records by sentence", "sentence_id", req.SentenceId)

	resp, err := s.client.ListBySentence(ctx, req)
	if err != nil {
		log.Error("failed to list history records", "sentence_id", req.SentenceId, "error", err)
		return nil, err
	}

	log.Debug("history records listed", "sentence_id", req.SentenceId, "count", len(resp.Data))
	return resp, nil
}
