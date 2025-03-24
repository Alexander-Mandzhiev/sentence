package service

import (
	"backend/protos/gen/go/sentences"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

func (s *Service) Delete(ctx context.Context, req *sentences.DeleteSentenceRequest) (*sentences.DeleteSentenceResponse, error) {
	op := "service.Delete"
	s.logger.Info("Deleting sentence", slog.String("op", op), slog.Any("request", req))

	if req == nil || req.Id <= 0 {
		s.logger.Error("Invalid input data", slog.String("op", op), slog.Any("request", req))
		return nil, status.Error(codes.InvalidArgument, "invalid input data")
	}

	if err := s.provider.Delete(ctx, req.Id); err != nil {
		s.logger.Error("Failed to delete sentence", slog.String("op", op), slog.String("error", err.Error()))
		return &sentences.DeleteSentenceResponse{Success: false}, err
	}

	s.logger.Info("Sentence deleted", slog.String("op", op), slog.Any("id", req.Id))
	return &sentences.DeleteSentenceResponse{Success: true}, nil
}
