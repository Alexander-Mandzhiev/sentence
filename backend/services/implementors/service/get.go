package service

import (
	"backend/protos/gen/go/implementors"
	"context"
	"golang.org/x/exp/slog"
)

func (s *Service) Get(ctx context.Context, req *implementors.GetImplementorRequest) (*implementors.ImplementorResponse, error) {
	op := "service.Get"
	s.logger.Info("Getting implementor", slog.String("op", op), slog.Any("request", req))

	implementor, err := s.provider.Get(ctx, int(req.Id))
	if err != nil {
		s.logger.Error("Failed to get implementor", slog.String("op", op), slog.String("error", err.Error()))
		return nil, err
	}

	s.logger.Info("Implementor retrieved", slog.String("op", op), slog.Any("implementor", implementor))
	return implementor, nil
}
