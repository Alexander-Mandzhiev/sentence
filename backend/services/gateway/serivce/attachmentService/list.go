package attachment_service

import (
	"backend/protos/gen/go/attachments"
	"context"
	"log/slog"
)

func (s *Service) List(ctx context.Context, req *attachments.ListAttachmentsRequest) (*attachments.AttachmentsListResponse, error) {
	const op = "attachment_service.List"
	log := s.logger.With(slog.String("op", op))

	resp, err := s.client.List(ctx, req)
	if err != nil {
		log.Error("failed to list attachments", "error", err)
		return nil, err
	}
	return resp, nil
}
