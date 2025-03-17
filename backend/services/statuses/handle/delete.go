package handle

import (
	"backend/protos/gen/go/statuses"
	"context"
)

func (s *serverAPI) Delete(ctx context.Context, req *statuses.DeleteStatusRequest) (*statuses.DeleteStatusResponse, error) {
	return s.service.Delete(ctx, req)
}
