package handle

import (
	"backend/protos/gen/go/priorities"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *serverAPI) Update(ctx context.Context, req *priorities.UpdatePrioritiesRequest) (*priorities.PrioritiesResponse, error) {
	resp, err := s.service.Update(ctx, req)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}
