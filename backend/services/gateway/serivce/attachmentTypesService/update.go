package attachment_types_service

import (
	"backend/protos/gen/go/attachment_types"
	"context"
	"log/slog"
)

func (s *Service) Update(ctx context.Context, req *attachment_types.AttachmentTypeResponse) (*attachment_types.AttachmentTypeResponse, error) {
	const op = "attachment_types_service.Update"
	log := s.logger.With(slog.String("op", op))

	log.Debug("updating attachment type", "id", req.Id)
	resp, err := s.client.Update(ctx, req)
	if err != nil {
		log.Error("failed to update attachment type", "id", req.Id, "error", err)
		return nil, err
	}

	log.Info("attachment type updated", "id", resp.Id)
	return resp, nil
}
