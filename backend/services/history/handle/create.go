package handle

import (
	"backend/protos/gen/go/history"
	"context"
)

func (h *serverAPI) Create(ctx context.Context, req *history.CreateHistoryRequest) (*history.HistoryResponse, error) {
	return h.service.Create(ctx, req)
}
