package handle

import (
	"backend/protos/gen/go/implementors"
	"context"
)

func (s *serverAPI) List(ctx context.Context, req *implementors.ListImplementorsRequest) (*implementors.ImplementorsListResponse, error) {
	resp, err := s.service.List(ctx)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
