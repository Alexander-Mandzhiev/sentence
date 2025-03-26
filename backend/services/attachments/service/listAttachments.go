package service

import (
	"backend/protos/gen/go/attachments"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

func (s *Service) ListAttachments(ctx context.Context, limit, offset int32) ([]*attachments.AttachmentResponse, error) {
	const op = "Service.ListAttachments"
	logger := s.logger.With(slog.String("op", op))

	att, err := s.provider.List(ctx, limit, offset)
	if err != nil {
		logger.Error("failed to list attachments", slog.String("error", err.Error()))
		return nil, status.Errorf(codes.Internal, "failed to list attachments")
	}

	return att, nil
}
