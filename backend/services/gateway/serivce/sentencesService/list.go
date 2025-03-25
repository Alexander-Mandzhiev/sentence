package sentences_service

import (
	"backend/protos/gen/go/sentences"
	"context"
	"log/slog"
)

func (s *Service) List(ctx context.Context, req *sentences.ListSentencesRequest) (*sentences.SentencesListResponse, error) {
	const op = "sentences_service.List"
	log := s.logger.With(slog.String("op", op))
	log.Debug("listing sentences")

	resp, err := s.client.List(ctx, req)
	if err != nil {
		log.Error("failed to list sentences", "error", err)
		return nil, err
	}

	log.Debug("sentences listed successfully", "count", len(resp.GetData()))
	return resp, nil
}
