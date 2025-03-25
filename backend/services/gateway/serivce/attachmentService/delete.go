package attachment_service

import (
	"backend/protos/gen/go/attachments"
	"context"
	"log/slog"
)

func (s *Service) Delete(ctx context.Context, req *attachments.DeleteAttachmentRequest) (*attachments.DeleteAttachmentResponse, error) {
	const op = "attachment_service.Delete"
	log := s.logger.With(slog.String("op", op))

	resp, err := s.client.Delete(ctx, req)
	if err != nil {
		log.Error("failed to delete attachment", "id", req.Id, "error", err)
		return nil, err
	}
	return resp, nil
}
