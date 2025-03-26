package attachment_service

import (
	"backend/protos/gen/go/attachments"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

func (s *Service) ListAttachments(ctx context.Context, limit, offset int32) (*attachments.AttachmentsListResponse, error) {
	const op = "attachment_service.ListAttachments"
	log := s.logger.With(slog.String("op", op))

	if limit <= 0 || limit > 100 {
		limit = 50
	}
	if offset < 0 {
		offset = 0
	}

	resp, err := s.client.ListAttachments(ctx, &attachments.ListAttachmentsRequest{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		log.Error("failed to list attachments", slog.String("error", err.Error()))
		return nil, status.Errorf(codes.Internal, "failed to list attachments")
	}

	return resp, nil
}
