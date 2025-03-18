package handle

import (
	"backend/protos/gen/go/sentences"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *serverAPI) Get(ctx context.Context, request *sentences.GetSentenceRequest) (*sentences.SentenceResponse, error) {
	resp, err := s.service.Get(ctx, request)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}
