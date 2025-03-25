package sentences_service

import (
	"backend/protos/gen/go/sentences"
	"context"
	"log/slog"
)

func (s *Service) Update(ctx context.Context, req *sentences.SentenceResponse) (*sentences.SentenceResponse, error) {
	const op = "sentences_service.Update"
	log := s.logger.With(slog.String("op", op), slog.Int64("id", req.GetId()))
	log.Debug("updating sentence")

	resp, err := s.client.Update(ctx, req)
	if err != nil {
		log.Error("failed to update sentence", "error", err)
		return nil, err
	}

	log.Info("sentence updated successfully", "status_id", resp.GetStatusId(), "implementor_id", resp.GetImplementorId())
	return resp, nil
}
