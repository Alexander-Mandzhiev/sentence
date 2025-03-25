package handle

import (
	"backend/protos/gen/go/sentences_attachments"
	"context"
	"google.golang.org/grpc"
)

type SentencesAttachmentsService interface {
	Create(ctx context.Context, request *sentences_attachments.CreateSentenceAttachmentRequest) (*sentences_attachments.SentenceAttachmentResponse, error)
	Delete(ctx context.Context, request *sentences_attachments.DeleteSentenceAttachmentRequest) (*sentences_attachments.DeleteSentenceAttachmentResponse, error)
	ListBySentence(ctx context.Context, request *sentences_attachments.ListBySentenceRequest) (*sentences_attachments.SentenceAttachmentsListResponse, error)
	ListByAttachment(ctx context.Context, request *sentences_attachments.ListByAttachmentRequest) (*sentences_attachments.SentenceAttachmentsListResponse, error)
}

type serverAPI struct {
	sentences_attachments.UnimplementedSentencesAttachmentsProviderServer
	service SentencesAttachmentsService
}

func Register(gRPCServer *grpc.Server, service SentencesAttachmentsService) {
	sentences_attachments.RegisterSentencesAttachmentsProviderServer(gRPCServer, &serverAPI{service: service})
}
