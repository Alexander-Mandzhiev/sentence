package handle

import (
	"backend/protos/gen/go/attachments"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *serverAPI) List(ctx context.Context, req *attachments.ListAttachmentsRequest) (*attachments.AttachmentsListResponse, error) {
	resp, err := s.service.List(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}
