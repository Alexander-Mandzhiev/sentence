package service

import (
	"backend/protos/gen/go/attachment_types"
	"context"
	"log/slog"
)

type AttachmentTypesProvider interface {
	Create(ctx context.Context, attachmentType *attachment_types.AttachmentTypeResponse) (int32, error)
	Get(ctx context.Context, id int32) (*attachment_types.AttachmentTypeResponse, error)
	Update(ctx context.Context, attachmentType *attachment_types.AttachmentTypeResponse) error
	Delete(ctx context.Context, id int32) error
	List(ctx context.Context) ([]*attachment_types.AttachmentTypeResponse, error)
}
type Service struct {
	logger   *slog.Logger
	provider AttachmentTypesProvider
}

func New(provider AttachmentTypesProvider, logger *slog.Logger) *Service {
	op := "service.New"
	if provider == nil {
		logger.Error("Attachment types provider is nil", slog.String("op", op))
		return nil
	}

	logger.Info("Service initialized", slog.String("op", op))
	return &Service{provider: provider, logger: logger}
}
