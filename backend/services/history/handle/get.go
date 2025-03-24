package handle

import (
	"backend/protos/gen/go/history"
	"context"
)

func (h *serverAPI) Get(ctx context.Context, req *history.GetHistoryRequest) (*history.HistoryResponse, error) {
	return h.service.Get(ctx, req)
}
