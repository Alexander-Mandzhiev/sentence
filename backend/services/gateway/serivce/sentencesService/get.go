package sentences_service

import (
	"backend/protos/gen/go/sentences"
	"context"
	"log/slog"
)

func (s *Service) Get(ctx context.Context, req *sentences.GetSentenceRequest) (*sentences.SentenceResponse, error) {
	const op = "sentences_service.Get"
	log := s.logger.With(slog.String("op", op), slog.Int64("id", req.GetId()))
	log.Debug("getting sentence")

	resp, err := s.client.Get(ctx, req)
	if err != nil {
		log.Error("failed to get sentence", "error", err)
		return nil, err
	}

	log.Debug("sentence retrieved successfully", "status_id", resp.GetStatusId(), "name", resp.GetName())
	return resp, nil
}
