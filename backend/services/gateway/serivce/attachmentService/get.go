package attachment_service

import (
	"backend/protos/gen/go/attachment_types"
	"backend/protos/gen/go/attachments"
	"backend/services/gateway/models"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

func (s *Service) Get(ctx context.Context, id int32) (*models.Attachment, error) {
	const op = "attachment_service.GetAttachment"
	log := s.logger.With(slog.String("op", op), slog.Int("id", int(id)))

	if id <= 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid attachment ID")
	}

	// Получаем само вложение
	attachment, err := s.client.GetAttachment(ctx, &attachments.GetAttachmentRequest{Id: id})
	if err != nil {
		log.Error("failed to get attachment", slog.String("error", err.Error()))
		return nil, status.Errorf(codes.Internal, "failed to get attachment")
	}

	// Получаем информацию о типе
	typeInfo, err := s.typesClient.Get(ctx, &attachment_types.GetAttachmentTypeRequest{
		Id: attachment.GetAttachmentTypeId(),
	})
	if err != nil {
		log.Error("failed to get attachment type",
			slog.Int("type_id", int(attachment.GetAttachmentTypeId())),
			slog.String("error", err.Error()))
		typeInfo = &attachment_types.AttachmentTypeResponse{}
	}

	return models.AttachmentFromProto(attachment, typeInfo), nil
}
