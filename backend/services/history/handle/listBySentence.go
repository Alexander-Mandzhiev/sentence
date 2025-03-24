package handle

import (
	"backend/protos/gen/go/history"
	"context"
)

func (h *serverAPI) ListBySentence(ctx context.Context, req *history.ListBySentenceRequest) (*history.HistoryListResponse, error) {
	return h.service.ListBySentence(ctx, req)
}
