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

func (s *Service) List(ctx context.Context, limit, offset int32) ([]*models.Attachment, error) {
	const op = "attachment_service.ListAttachments"
	log := s.logger.With(slog.String("op", op))

	if limit <= 0 || limit > 100 {
		limit = 50
	}
	if offset < 0 {
		offset = 0
	}

	// Получаем список вложений
	resp, err := s.client.ListAttachments(ctx, &attachments.ListAttachmentsRequest{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		log.Error("failed to list attachments", slog.String("error", err.Error()))
		return nil, status.Errorf(codes.Internal, "failed to list attachments")
	}

	// Собираем ID всех типов вложений
	typeIDs := make([]int32, 0, len(resp.Data))
	for _, att := range resp.Data {
		typeIDs = append(typeIDs, att.GetAttachmentTypeId())
	}

	// Получаем информацию о всех типах
	typeInfos := make(map[int32]*attachment_types.AttachmentTypeResponse)
	if len(typeIDs) > 0 {
		listResp, err := s.typesClient.List(ctx, &attachment_types.ListAttachmentTypesRequest{})
		if err == nil {
			for _, t := range listResp.Data {
				typeInfos[t.GetId()] = t
			}
		} else {
			log.Error("failed to get attachment types", slog.String("error", err.Error()))
		}
	}

	// Формируем результат
	result := make([]*models.Attachment, 0, len(resp.Data))
	for _, att := range resp.Data {
		typeInfo, ok := typeInfos[att.GetAttachmentTypeId()]
		if !ok {
			typeInfo = &attachment_types.AttachmentTypeResponse{}
		}
		result = append(result, models.AttachmentFromProto(att, typeInfo))
	}

	return result, nil
}
