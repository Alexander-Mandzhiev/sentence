package handle

import (
	"backend/protos/gen/go/directions"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *serverAPI) Get(ctx context.Context, req *directions.GetDirectionRequest) (*directions.DirectionResponse, error) {
	resp, err := s.service.Get(ctx, req)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}
	return resp, nil
}
