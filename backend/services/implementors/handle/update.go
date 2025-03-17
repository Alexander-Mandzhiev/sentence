package handle

import (
	"backend/protos/gen/go/implementors"
	"context"
)

func (s *serverAPI) Update(ctx context.Context, req *implementors.UpdateImplementorRequest) (*implementors.ImplementorResponse, error) {
	resp, err := s.service.Update(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
