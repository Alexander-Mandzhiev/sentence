package handle

import (
	"backend/protos/gen/go/sentences_attachments"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *serverAPI) ListByAttachment(ctx context.Context, req *sentences_attachments.ListByAttachmentRequest) (*sentences_attachments.SentenceAttachmentsListResponse, error) {
	if req.AttachmentId == 0 {
		return nil, status.Error(codes.InvalidArgument, "attachment_id is required")
	}

	resp, err := s.service.ListByAttachment(ctx, req)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}
