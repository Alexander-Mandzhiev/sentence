package attachment_types_service

import (
	"backend/protos/gen/go/attachment_types"
	"context"
	"log/slog"
)

type AttachmentTypesClient interface {
	Create(ctx context.Context, req *attachment_types.CreateAttachmentTypeRequest) (*attachment_types.AttachmentTypeResponse, error)
	Get(ctx context.Context, req *attachment_types.GetAttachmentTypeRequest) (*attachment_types.AttachmentTypeResponse, error)
	Update(ctx context.Context, req *attachment_types.AttachmentTypeResponse) (*attachment_types.AttachmentTypeResponse, error)
	Delete(ctx context.Context, req *attachment_types.DeleteAttachmentTypeRequest) (*attachment_types.DeleteAttachmentTypeResponse, error)
	List(ctx context.Context, req *attachment_types.ListAttachmentTypesRequest) (*attachment_types.AttachmentTypesListResponse, error)
}

type Service struct {
	client attachment_types.AttachmentTypesProviderClient
	logger *slog.Logger
}

func New(client attachment_types.AttachmentTypesProviderClient, logger *slog.Logger) *Service {
	return &Service{
		client: client,
		logger: logger,
	}
}
