package attachment_service

import (
	"backend/protos/gen/go/attachments"
	"context"
	"log/slog"
)

type AttachmentsClient interface {
	Create(ctx context.Context, req *attachments.CreateAttachmentRequest) (*attachments.AttachmentResponse, error)
	Get(ctx context.Context, req *attachments.GetAttachmentRequest) (*attachments.AttachmentResponse, error)
	Update(ctx context.Context, req *attachments.AttachmentResponse) (*attachments.AttachmentResponse, error)
	Delete(ctx context.Context, req *attachments.DeleteAttachmentRequest) (*attachments.DeleteAttachmentResponse, error)
	List(ctx context.Context, req *attachments.ListAttachmentsRequest) (*attachments.AttachmentsListResponse, error)
}
type Service struct {
	client attachments.AttachmentsProviderClient
	logger *slog.Logger
}

func New(client attachments.AttachmentsProviderClient, logger *slog.Logger) *Service {
	return &Service{
		client: client,
		logger: logger,
	}
}
