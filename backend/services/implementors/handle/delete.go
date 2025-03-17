package handle

import (
	"backend/protos/gen/go/implementors"
	"context"
)

func (s *serverAPI) Delete(ctx context.Context, req *implementors.DeleteImplementorRequest) (*implementors.DeleteImplementorResponse, error) {
	resp, err := s.service.Delete(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
