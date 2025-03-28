package handle

import (
	"backend/protos/gen/go/sentences"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *serverAPI) List(ctx context.Context, request *sentences.ListSentencesRequest) (*sentences.SentencesListResponse, error) {
	resp, err := s.service.List(ctx, request.GetLimit(), request.GetOffset())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}
