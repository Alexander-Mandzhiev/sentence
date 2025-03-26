package handle

import (
	"backend/protos/gen/go/attachments"
	"context"
	"google.golang.org/grpc"
	"io"
)

type AttachmentsService interface {
	CreateAttachment(ctx context.Context, metadata *attachments.AttachmentMetadata, file io.Reader) (*attachments.AttachmentResponse, error)
	GetAttachment(ctx context.Context, id int32) (*attachments.AttachmentResponse, error)
	DeleteAttachment(ctx context.Context, id int32) error
	ListAttachments(ctx context.Context, limit, offset int32) ([]*attachments.AttachmentResponse, error)
	DownloadFile(ctx context.Context, id int32) (io.ReadCloser, *attachments.FileMetadata, error)
}

type serverAPI struct {
	attachments.UnimplementedAttachmentsProviderServer
	service AttachmentsService
}

func Register(gRPCServer *grpc.Server, service AttachmentsService) {
	attachments.RegisterAttachmentsProviderServer(gRPCServer, &serverAPI{service: service})
}
