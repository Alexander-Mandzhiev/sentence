package handle

import (
	"backend/protos/gen/go/departments"
	"context"
)

func (s *serverAPI) Update(ctx context.Context, req *departments.DepartmentResponse) (*departments.DepartmentResponse, error) {
	resp, err := s.service.Update(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
