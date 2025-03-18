package handle

import (
	"backend/protos/gen/go/departments"
	"context"
)

func (s *serverAPI) Create(ctx context.Context, req *departments.CreateDepartmentRequest) (*departments.DepartmentResponse, error) {
	resp, err := s.service.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
