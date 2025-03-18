package handle

import (
	"backend/protos/gen/go/departments"
	"context"
)

func (s *serverAPI) Delete(ctx context.Context, req *departments.DeleteDepartmentRequest) (*departments.DeleteDepartmentResponse, error) {
	resp, err := s.service.Delete(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
