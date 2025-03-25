package handle

import (
	"backend/protos/gen/go/sentences_attachments"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *serverAPI) Create(ctx context.Context, req *sentences_attachments.CreateSentenceAttachmentRequest) (*sentences_attachments.SentenceAttachmentResponse, error) {
	if req.SentenceId == 0 || req.AttachmentId == 0 {
		return nil, status.Error(codes.InvalidArgument, "sentence_id and attachment_id are required")
	}

	resp, err := s.service.Create(ctx, req)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}
