package handle

import (
	"backend/protos/gen/go/departments"
	"context"
)

func (s *serverAPI) Get(ctx context.Context, req *departments.GetDepartmentRequest) (*departments.DepartmentResponse, error) {
	resp, err := s.service.Get(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
