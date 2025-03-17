package handle

import (
	"backend/protos/gen/go/priorities"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *serverAPI) Get(ctx context.Context, req *priorities.GetPrioritiesRequest) (*priorities.PrioritiesResponse, error) {
	resp, err := s.service.Get(ctx, req)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}
	return resp, nil
}
