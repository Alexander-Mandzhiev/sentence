package attachment_service

import (
	"backend/protos/gen/go/attachments"
	attachment_types_service "backend/services/gateway/serivce/attachmentTypesService"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"log/slog"
)

type AttachmentsClient interface {
	CreateAttachment(ctx context.Context, opts ...grpc.CallOption) (attachments.AttachmentsProvider_CreateAttachmentClient, error)
	GetAttachment(ctx context.Context, in *attachments.GetAttachmentRequest, opts ...grpc.CallOption) (*attachments.AttachmentResponse, error)
	DeleteAttachment(ctx context.Context, in *attachments.DeleteAttachmentRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ListAttachments(ctx context.Context, in *attachments.ListAttachmentsRequest, opts ...grpc.CallOption) (*attachments.AttachmentsListResponse, error)
	DownloadFile(ctx context.Context, in *attachments.GetAttachmentRequest, opts ...grpc.CallOption) (attachments.AttachmentsProvider_DownloadFileClient, error)
}

type Service struct {
	client      attachments.AttachmentsProviderClient
	typesClient *attachment_types_service.Service
	logger      *slog.Logger
}

func New(client attachments.AttachmentsProviderClient, typesClient *attachment_types_service.Service, logger *slog.Logger) *Service {
	return &Service{
		client:      client,
		typesClient: typesClient,
		logger:      logger,
	}
}
