package attachment_types_service

import (
	"backend/protos/gen/go/attachment_types"
	"context"
	"log/slog"
)

func (s *Service) Create(ctx context.Context, req *attachment_types.CreateAttachmentTypeRequest) (*attachment_types.AttachmentTypeResponse, error) {
	const op = "attachment_types_service.Create"
	log := s.logger.With(slog.String("op", op))

	log.Debug("creating attachment type", "name", req.Name)
	resp, err := s.client.Create(ctx, req)
	if err != nil {
		log.Error("failed to create attachment type", "error", err)
		return nil, err
	}

	log.Info("attachment type created", "id", resp.Id)
	return resp, nil
}
