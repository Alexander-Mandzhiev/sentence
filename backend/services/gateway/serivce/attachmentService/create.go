package attachment_service

import (
	"backend/protos/gen/go/attachments"
	"context"
	"log/slog"
)

func (s *Service) Create(ctx context.Context, req *attachments.CreateAttachmentRequest) (*attachments.AttachmentResponse, error) {
	const op = "attachment_service.Create"
	log := s.logger.With(slog.String("op", op))

	resp, err := s.client.Create(ctx, req)
	if err != nil {
		log.Error("failed to create attachment", "error", err)
		return nil, err
	}
	return resp, nil
}
