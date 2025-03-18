package handle

import (
	"backend/protos/gen/go/attachment_types"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *serverAPI) Create(ctx context.Context, req *attachment_types.CreateAttachmentTypeRequest) (*attachment_types.AttachmentTypeResponse, error) {
	resp, err := s.service.Create(ctx, req)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}
