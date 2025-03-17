package handle

import (
	"backend/protos/gen/go/implementors"
	"context"
)

func (s *serverAPI) Create(ctx context.Context, req *implementors.CreateImplementorRequest) (*implementors.ImplementorResponse, error) {
	resp, err := s.service.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
