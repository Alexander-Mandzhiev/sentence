package handle

import (
	"backend/protos/gen/go/statuses"
	"context"
)

func (s *serverAPI) Update(ctx context.Context, req *statuses.UpdateStatusRequest) (*statuses.StatusResponse, error) {
	return s.service.Update(ctx, req)
}
