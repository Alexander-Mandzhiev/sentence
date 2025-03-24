package service

import (
	"backend/protos/gen/go/sentences"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

func (s *Service) Get(ctx context.Context, req *sentences.GetSentenceRequest) (*sentences.SentenceResponse, error) {
	op := "service.Get"
	s.logger.Info("Getting sentence", slog.String("op", op), slog.Any("request", req))

	if req == nil || req.Id <= 0 {
		s.logger.Error("Invalid input data", slog.String("op", op), slog.Any("request", req))
		return nil, status.Error(codes.InvalidArgument, "invalid input data")
	}

	sentence, err := s.provider.Get(ctx, req.Id)
	if err != nil {
		s.logger.Error("Failed to get sentence", slog.String("op", op), slog.String("error", err.Error()))
		return nil, err
	}

	s.logger.Info("Sentence retrieved", slog.String("op", op), slog.Any("sentence", sentence))
	return sentence, nil
}
