package handle

import (
	"backend/protos/gen/go/attachments"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *serverAPI) ListAttachments(ctx context.Context, req *attachments.ListAttachmentsRequest) (*attachments.AttachmentsListResponse, error) {
	limit := req.GetLimit()
	offset := req.GetOffset()

	if limit <= 0 || limit > 100 {
		limit = 50
	}
	if offset < 0 {
		offset = 0
	}

	attachmentsList, err := s.service.ListAttachments(ctx, limit, offset)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list attachments: %v", err)
	}

	return &attachments.AttachmentsListResponse{Data: attachmentsList}, nil
}
