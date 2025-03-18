package service

import (
	"backend/protos/gen/go/sentences"
	"context"
	"log/slog"
)

func (s *Service) Update(ctx context.Context, req *sentences.SentenceResponse) (*sentences.SentenceResponse, error) {
	op := "service.Update"
	s.logger.Info("Updating sentence", slog.String("op", op), slog.Any("request", req))

	if err := s.provider.Update(ctx, req); err != nil {
		s.logger.Error("Failed to update sentence", slog.String("op", op), slog.String("error", err.Error()))
		return nil, err
	}

	s.logger.Info("Sentence updated", slog.String("op", op), slog.Any("sentence", req))
	return req, nil
}
