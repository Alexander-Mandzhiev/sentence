package handle

import (
	"backend/protos/gen/go/attachments"
	"context"
	"google.golang.org/grpc"
)

type AttachmentsService interface {
	Create(ctx context.Context, req *attachments.CreateAttachmentRequest) (*attachments.AttachmentResponse, error)
	Update(ctx context.Context, req *attachments.AttachmentResponse) (*attachments.AttachmentResponse, error)
	Get(ctx context.Context, req *attachments.GetAttachmentRequest) (*attachments.AttachmentResponse, error)
	Delete(ctx context.Context, req *attachments.DeleteAttachmentRequest) (*attachments.DeleteAttachmentResponse, error)
	List(ctx context.Context) (*attachments.AttachmentsListResponse, error)
}

type serverAPI struct {
	attachments.UnimplementedAttachmentsProviderServer
	service AttachmentsService
}

func Register(gRPCServer *grpc.Server, service AttachmentsService) {
	attachments.RegisterAttachmentsProviderServer(gRPCServer, &serverAPI{service: service})
}
