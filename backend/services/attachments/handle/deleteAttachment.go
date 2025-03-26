package handle

import (
	"backend/protos/gen/go/attachments"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *serverAPI) DeleteAttachment(ctx context.Context, req *attachments.DeleteAttachmentRequest) (*emptypb.Empty, error) {
	if req == nil || req.Id <= 0 {
		return nil, status.Error(codes.InvalidArgument, "id must be positive")
	}

	if err := s.service.DeleteAttachment(ctx, req.Id); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete attachment: %v", err)
	}

	return &emptypb.Empty{}, nil
}
