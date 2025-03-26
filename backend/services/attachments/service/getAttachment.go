package service

import (
	"backend/protos/gen/go/attachments"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

func (s *Service) GetAttachment(ctx context.Context, id int32) (*attachments.AttachmentResponse, error) {
	const op = "Service.GetAttachment"
	logger := s.logger.With(slog.String("op", op), slog.Int("id", int(id)))

	if id <= 0 {
		return nil, status.Error(codes.InvalidArgument, "ID must be positive")
	}

	attachment, err := s.provider.Get(ctx, id)
	if err != nil {
		logger.Error("failed to get attachment", slog.String("error", err.Error()))
		return nil, status.Errorf(codes.Internal, "failed to get attachment")
	}

	return attachment, nil
}
