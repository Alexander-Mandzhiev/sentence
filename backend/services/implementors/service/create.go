package service

import (
	"backend/protos/gen/go/implementors"
	"context"
	"log/slog"
)

func (s *Service) Create(ctx context.Context, req *implementors.CreateImplementorRequest) (*implementors.ImplementorResponse, error) {
	op := "service.Create"
	s.logger.Info("Creating implementor", slog.String("op", op), slog.Any("request", req))

	implementor := &implementors.ImplementorResponse{
		Name:     req.Name,
		Email:    req.Email,
		IsActive: req.IsActive,
	}

	id, err := s.provider.Create(ctx, implementor)
	if err != nil {
		s.logger.Error("Failed to create implementor", slog.String("op", op), slog.String("error", err.Error()))
		return nil, err
	}

	implementor.Id = int32(id)
	s.logger.Info("Implementor created", slog.String("op", op), slog.Any("implementor", implementor))
	return implementor, nil
}
