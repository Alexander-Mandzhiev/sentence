package attachment_types_service

import (
	"backend/protos/gen/go/attachment_types"
	"context"
	"log/slog"
)

func (s *Service) List(ctx context.Context, req *attachment_types.ListAttachmentTypesRequest) (*attachment_types.AttachmentTypesListResponse, error) {
	const op = "attachment_types_service.List"
	log := s.logger.With(slog.String("op", op))

	log.Debug("listing attachment types")
	resp, err := s.client.List(ctx, req)
	if err != nil {
		log.Error("failed to list attachment types", "error", err)
		return nil, err
	}

	log.Debug("attachment types listed", "count", len(resp.Data))
	return resp, nil
}
