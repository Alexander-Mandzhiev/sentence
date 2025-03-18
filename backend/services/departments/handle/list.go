package handle

import (
	"backend/protos/gen/go/departments"
	"context"
)

func (s *serverAPI) List(ctx context.Context, req *departments.ListDepartmentsRequest) (*departments.DepartmentsListResponse, error) {
	resp, err := s.service.List(ctx)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
