package handle

import (
	"backend/protos/gen/go/attachment_types"
	"context"
	"google.golang.org/grpc"
)

type AttachmentTypesService interface {
	Create(ctx context.Context, request *attachment_types.CreateAttachmentTypeRequest) (*attachment_types.AttachmentTypeResponse, error)
	Get(ctx context.Context, request *attachment_types.GetAttachmentTypeRequest) (*attachment_types.AttachmentTypeResponse, error)
	Update(ctx context.Context, attachmentType *attachment_types.AttachmentTypeResponse) (*attachment_types.AttachmentTypeResponse, error)
	Delete(ctx context.Context, request *attachment_types.DeleteAttachmentTypeRequest) (*attachment_types.DeleteAttachmentTypeResponse, error)
	List(ctx context.Context) (*attachment_types.AttachmentTypesListResponse, error)
}

type serverAPI struct {
	attachment_types.UnimplementedAttachmentTypesProviderServer
	service AttachmentTypesService
}

func Register(gRPCServer *grpc.Server, service AttachmentTypesService) {
	attachment_types.RegisterAttachmentTypesProviderServer(gRPCServer, &serverAPI{service: service})
}
