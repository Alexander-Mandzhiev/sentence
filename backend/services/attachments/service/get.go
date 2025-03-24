package service

import (
	"backend/protos/gen/go/attachments"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

func (s *Service) Get(ctx context.Context, req *attachments.GetAttachmentRequest) (*attachments.AttachmentResponse, error) {
	op := "service.Get"
	s.logger.Info("Getting attachment", slog.String("op", op), slog.Int64("id", int64(req.Id)))

	if req == nil || req.Id <= 0 {
		s.logger.Error("Invalid input data", slog.String("op", op), slog.Any("request", req))
		return nil, status.Error(codes.InvalidArgument, "invalid input data: request is nil or ID is invalid")
	}

	attachment, err := s.provider.Get(ctx, req.Id)
	if err != nil {
		s.logger.Error("Failed to get attachment", slog.String("op", op), slog.String("error", err.Error()))
		return nil, err
	}

	s.logger.Info("Attachment fetched", slog.String("op", op), slog.Any("attachment", attachment))
	return attachment, nil
}
