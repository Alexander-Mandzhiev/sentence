package attachment_service

import (
	"backend/protos/gen/go/attachments"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

func (s *Service) DeleteAttachment(ctx context.Context, id int32) error {
	const op = "attachment_service.DeleteAttachment"
	log := s.logger.With(slog.String("op", op), slog.Int("id", int(id)))
	if id <= 0 {
		return status.Error(codes.InvalidArgument, "invalid attachment ID")
	}

	_, err := s.client.DeleteAttachment(ctx, &attachments.DeleteAttachmentRequest{Id: id})
	if err != nil {
		log.Error("failed to delete attachment", slog.String("error", err.Error()))
		return status.Errorf(codes.Internal, "failed to delete attachment")
	}

	return nil
}
