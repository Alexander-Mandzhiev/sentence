package handle

import (
	"backend/protos/gen/go/statuses"
	"context"
)

func (s *serverAPI) Get(ctx context.Context, req *statuses.GetStatusRequest) (*statuses.StatusResponse, error) {
	return s.service.Get(ctx, req)
}
