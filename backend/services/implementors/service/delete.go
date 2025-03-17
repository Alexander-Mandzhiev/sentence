package service

import (
	"backend/protos/gen/go/implementors"
	"context"
	"golang.org/x/exp/slog"
)

func (s *Service) Delete(ctx context.Context, req *implementors.DeleteImplementorRequest) (*implementors.DeleteImplementorResponse, error) {
	op := "service.Delete"
	s.logger.Info("Deleting implementor", slog.String("op", op), slog.Any("request", req))

	err := s.provider.Delete(ctx, int(req.Id))
	if err != nil {
		s.logger.Error("Failed to delete implementor", slog.String("op", op), slog.String("error", err.Error()))
		return &implementors.DeleteImplementorResponse{Success: false}, err
	}

	s.logger.Info("Implementor deleted", slog.String("op", op), slog.Any("id", req.Id))
	return &implementors.DeleteImplementorResponse{Success: true}, nil
}
