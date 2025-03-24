package attachment_types_service

import (
	"backend/protos/gen/go/attachment_types"
	"context"
	"log/slog"
)

func (s *Service) Get(ctx context.Context, req *attachment_types.GetAttachmentTypeRequest) (*attachment_types.AttachmentTypeResponse, error) {
	const op = "attachment_types_service.Get"
	log := s.logger.With(slog.String("op", op))

	log.Debug("getting attachment type", "id", req.Id)
	resp, err := s.client.Get(ctx, req)
	if err != nil {
		log.Error("failed to get attachment type", "id", req.Id, "error", err)
		return nil, err
	}

	return resp, nil
}
