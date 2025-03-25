package attachment_service

import (
	"backend/protos/gen/go/attachments"
	"context"
	"log/slog"
)

func (s *Service) Update(ctx context.Context, req *attachments.AttachmentResponse) (*attachments.AttachmentResponse, error) {
	const op = "attachment_service.Update"
	log := s.logger.With(slog.String("op", op))

	resp, err := s.client.Update(ctx, req)
	if err != nil {
		log.Error("failed to update attachment", "id", req.Id, "error", err)
		return nil, err
	}
	return resp, nil
}
