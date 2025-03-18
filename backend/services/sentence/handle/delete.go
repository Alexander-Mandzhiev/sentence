package handle

import (
	"backend/protos/gen/go/sentences"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *serverAPI) Delete(ctx context.Context, request *sentences.DeleteSentenceRequest) (*sentences.DeleteSentenceResponse, error) {
	resp, err := s.service.Delete(ctx, request)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}
