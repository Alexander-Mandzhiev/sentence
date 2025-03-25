package sentences_service

import (
	"backend/protos/gen/go/sentences"
	"context"
	"log/slog"
	"time"
)

func (s *Service) Create(ctx context.Context, req *sentences.CreateSentenceRequest) (*sentences.SentenceResponse, error) {
	const op = "sentences_service.Create"
	log := s.logger.With(slog.String("op", op), slog.String("name", req.GetName()), slog.Int("status_id", int(req.GetStatusId())), slog.Int("direction_id", int(req.GetDirectionId())))
	log.Debug("creating new sentence")

	resp, err := s.client.Create(ctx, req)
	if err != nil {
		log.Error("failed to create sentence", "error", err)
		return nil, err
	}

	log.Info("sentence created successfully", "id", resp.GetId(), "created_at", resp.GetCreatedAt().AsTime().Format(time.RFC3339))
	return resp, nil
}
