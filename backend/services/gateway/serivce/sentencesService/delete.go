package sentences_service

import (
	"backend/protos/gen/go/sentences"
	"context"
	"log/slog"
)

func (s *Service) Delete(ctx context.Context, req *sentences.DeleteSentenceRequest) (*sentences.DeleteSentenceResponse, error) {
	const op = "sentences_service.Delete"
	log := s.logger.With(slog.String("op", op), slog.Int64("id", req.GetId()))
	log.Debug("deleting sentence")

	resp, err := s.client.Delete(ctx, req)
	if err != nil {
		log.Error("failed to delete sentence", "error", err)
		return nil, err
	}

	if resp.GetSuccess() {
		log.Info("sentence deleted successfully")
	} else {
		log.Warn("sentence deletion reported failure")
	}

	return resp, nil
}
