package service

import (
	"backend/protos/gen/go/attachments"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

func (s *Service) Delete(ctx context.Context, req *attachments.DeleteAttachmentRequest) (*attachments.DeleteAttachmentResponse, error) {
	op := "service.Delete"
	s.logger.Info("Deleting attachment", slog.String("op", op), slog.Int64("id", int64(req.Id)))

	if req.Id <= 0 {
		s.logger.Error("Invalid input data", slog.String("op", op), slog.Any("request", req))
		return nil, status.Error(codes.InvalidArgument, "invalid input data: request is nil or ID is invalid")
	}

	if err := s.provider.Delete(ctx, req.Id); err != nil {
		s.logger.Error("Failed to delete attachment", slog.String("op", op), slog.String("error", err.Error()))
		return nil, err
	}

	s.logger.Info("Attachment deleted", slog.String("op", op), slog.Int64("id", int64(req.Id)))
	return &attachments.DeleteAttachmentResponse{Success: true}, nil
}
