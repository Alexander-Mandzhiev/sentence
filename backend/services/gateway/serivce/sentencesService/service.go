package sentences_service

import (
	"backend/protos/gen/go/sentences"
	"context"
)

type SentencesClient interface {
	Create(ctx context.Context, req *sentences.CreateSentenceRequest) (*sentences.SentenceResponse, error)
	Get(ctx context.Context, req *sentences.GetSentenceRequest) (*sentences.SentenceResponse, error)
	Update(ctx context.Context, req *sentences.SentenceResponse) (*sentences.SentenceResponse, error)
	Delete(ctx context.Context, req *sentences.DeleteSentenceRequest) (*sentences.DeleteSentenceResponse, error)
	List(ctx context.Context, req *sentences.ListSentencesRequest) (*sentences.SentencesListResponse, error)
}
