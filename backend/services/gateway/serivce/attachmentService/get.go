package attachment_service

import (
	"backend/protos/gen/go/attachments"
	"context"
	"log/slog"
)

func (s *Service) Get(ctx context.Context, req *attachments.GetAttachmentRequest) (*attachments.AttachmentResponse, error) {
	const op = "attachment_service.Get"
	log := s.logger.With(slog.String("op", op))

	resp, err := s.client.Get(ctx, req)
	if err != nil {
		log.Error("failed to get attachment", "id", req.Id, "error", err)
		return nil, err
	}
	return resp, nil
}
