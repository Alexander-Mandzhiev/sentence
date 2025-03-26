package handle

import (
	"backend/protos/gen/go/attachments"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *serverAPI) GetAttachment(ctx context.Context, req *attachments.GetAttachmentRequest) (*attachments.AttachmentResponse, error) {
	if req == nil || req.Id <= 0 {
		return nil, status.Error(codes.InvalidArgument, "id must be positive")
	}

	attachment, err := s.service.GetAttachment(ctx, req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get attachment: %v", err)
	}

	return attachment, nil
}
