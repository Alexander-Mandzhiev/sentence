package service

import (
	"backend/protos/gen/go/sentences"
	"context"
	"log/slog"
)

func (s *Service) Create(ctx context.Context, req *sentences.CreateSentenceRequest) (*sentences.SentenceResponse, error) {
	op := "service.Create"
	s.logger.Info("Creating sentence", slog.String("op", op), slog.Any("request", req))

	sentence := &sentences.SentenceResponse{
		StatusId:      req.StatusId,
		Name:          req.Name,
		Phone:         req.Phone,
		DepartmentId:  req.DepartmentId,
		Problem:       req.Problem,
		Solution:      req.Solution,
		DirectionId:   req.DirectionId,
		ImplementorId: req.ImplementorId,
		PriorityId:    req.PriorityId,
		Encouragement: req.Encouragement,
	}

	if err := s.provider.Create(ctx, sentence); err != nil {
		s.logger.Error("Failed to create sentence", slog.String("op", op), slog.String("error", err.Error()))
		return nil, err
	}

	s.logger.Info("Sentence created", slog.String("op", op), slog.Any("sentence", sentence))
	return sentence, nil
}
