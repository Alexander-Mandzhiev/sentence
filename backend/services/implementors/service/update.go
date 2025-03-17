package service

import (
	"backend/protos/gen/go/implementors"
	"context"
	"log/slog"
)

func (s *Service) Update(ctx context.Context, req *implementors.UpdateImplementorRequest) (*implementors.ImplementorResponse, error) {
	op := "service.Update"
	s.logger.Info("Updating implementor", slog.String("op", op), slog.Any("request", req))

	implementor := &implementors.ImplementorResponse{
		Id:       req.Id,
		Name:     req.Name,
		Email:    req.Email,
		IsActive: req.IsActive,
	}

	err := s.provider.Update(ctx, implementor)
	if err != nil {
		s.logger.Error("Failed to update implementor", slog.String("op", op), slog.String("error", err.Error()))
		return nil, err
	}

	s.logger.Info("Implementor updated", slog.String("op", op), slog.Any("implementor", implementor))
	return implementor, nil
}
