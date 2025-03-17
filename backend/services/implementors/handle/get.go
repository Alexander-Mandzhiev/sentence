package handle

import (
	"backend/protos/gen/go/implementors"
	"context"
)

func (s *serverAPI) Get(ctx context.Context, req *implementors.GetImplementorRequest) (*implementors.ImplementorResponse, error) {
	resp, err := s.service.Get(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
