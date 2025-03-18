package service

import (
	"backend/protos/gen/go/departments"
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log/slog"
)

func (s *Service) Update(ctx context.Context, req *departments.DepartmentResponse) (*departments.DepartmentResponse, error) {
	op := "service.Update"
	s.logger.Info("Updating department", slog.String("op", op), slog.Any("request", req))
	req.UpdatedAt = timestamppb.Now()

	err := s.provider.Update(ctx, req)
	if err != nil {
		s.logger.Error("Failed to update department", slog.String("op", op), slog.String("error", err.Error()))
		return nil, err
	}

	s.logger.Info("Department updated", slog.String("op", op), slog.Any("department", req))
	return req, nil
}
