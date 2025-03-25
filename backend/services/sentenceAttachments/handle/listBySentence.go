package handle

import (
	"backend/protos/gen/go/sentences_attachments"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *serverAPI) ListBySentence(ctx context.Context, req *sentences_attachments.ListBySentenceRequest) (*sentences_attachments.SentenceAttachmentsListResponse, error) {
	if req.SentenceId == 0 {
		return nil, status.Error(codes.InvalidArgument, "sentence_id is required")
	}

	resp, err := s.service.ListBySentence(ctx, req)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}
