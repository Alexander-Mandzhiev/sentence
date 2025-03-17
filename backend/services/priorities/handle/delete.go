package handle

import (
	"backend/protos/gen/go/priorities"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *serverAPI) Delete(ctx context.Context, req *priorities.DeletePrioritiesRequest) (*priorities.DeletePrioritiesResponse, error) {
	resp, err := s.service.Delete(ctx, req)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}
