package service

import (
	"backend/protos/gen/go/sentences"
	"context"
	"log/slog"
)

func (s *Service) List(ctx context.Context) (*sentences.SentencesListResponse, error) {
	op := "service.List"
	s.logger.Info("Listing sentences", slog.String("op", op))

	sentencesList, err := s.provider.List(ctx)
	if err != nil {
		s.logger.Error("Failed to list sentences", slog.String("op", op), slog.String("error", err.Error()))
		return nil, err
	}

	s.logger.Info("Sentences listed", slog.String("op", op), slog.Int("count", len(sentencesList)))
	return &sentences.SentencesListResponse{Data: sentencesList}, nil
}
