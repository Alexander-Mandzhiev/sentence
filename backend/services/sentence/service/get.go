package service

import (
	"backend/protos/gen/go/sentences"
	"context"
	"log/slog"
)

func (s *Service) Get(ctx context.Context, req *sentences.GetSentenceRequest) (*sentences.SentenceResponse, error) {
	op := "service.Get"
	s.logger.Info("Getting sentence", slog.String("op", op), slog.Any("request", req))

	sentence, err := s.provider.Get(ctx, req.Id)
	if err != nil {
		s.logger.Error("Failed to get sentence", slog.String("op", op), slog.String("error", err.Error()))
		return nil, err
	}

	s.logger.Info("Sentence retrieved", slog.String("op", op), slog.Any("sentence", sentence))
	return sentence, nil
}
