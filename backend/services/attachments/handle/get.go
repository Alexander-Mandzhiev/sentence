package handle

import (
	"backend/protos/gen/go/attachments"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *serverAPI) Get(ctx context.Context, req *attachments.GetAttachmentRequest) (*attachments.AttachmentResponse, error) {
	resp, err := s.service.Get(ctx, req)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}
