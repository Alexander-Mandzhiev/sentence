package attachment_types_service

import (
	"backend/protos/gen/go/attachment_types"
	"context"
	"log/slog"
)

func (s *Service) Delete(ctx context.Context, req *attachment_types.DeleteAttachmentTypeRequest) (*attachment_types.DeleteAttachmentTypeResponse, error) {
	const op = "attachment_types_service.Delete"
	log := s.logger.With(slog.String("op", op))

	log.Debug("deleting attachment type", "id", req.Id)
	resp, err := s.client.Delete(ctx, req)
	if err != nil {
		log.Error("failed to delete attachment type", "id", req.Id, "error", err)
		return nil, err
	}

	log.Info("attachment type deleted", "id", req.Id, "success", resp.Success)
	return resp, nil
}
