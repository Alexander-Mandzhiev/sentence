package service

import (
	"backend/protos/gen/go/attachments"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

func (s *Service) Update(ctx context.Context, req *attachments.AttachmentResponse) (*attachments.AttachmentResponse, error) {
	op := "service.Update"
	s.logger.Info("Updating attachment", slog.String("op", op), slog.Any("request", req))

	if req == nil || req.Id <= 0 || req.FileName == "" || req.FilePath == "" {
		s.logger.Error("Invalid input data", slog.String("op", op), slog.Any("request", req))
		return nil, status.Error(codes.InvalidArgument, "invalid input data: request is nil or required fields are empty/invalid")
	}

	if err := s.provider.Update(ctx, req); err != nil {
		s.logger.Error("Failed to update attachment", slog.String("op", op), slog.String("error", err.Error()))
		return nil, err
	}

	s.logger.Info("Attachment updated", slog.String("op", op), slog.Any("attachment", req))
	return req, nil
}
