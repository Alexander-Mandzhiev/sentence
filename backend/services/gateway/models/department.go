package models

import (
	"backend/protos/gen/go/departments"
)

type Department struct {
	ID          int32  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IsActive    bool   `json:"is_active"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

func DepartmentFromProto(pb *departments.DepartmentResponse) *Department {
	return &Department{
		ID:          pb.GetId(),
		Name:        pb.GetName(),
		Description: pb.GetDescription(),
		IsActive:    pb.GetIsActive(),
		CreatedAt:   convertTimestampToRFC3339(pb.GetCreatedAt()),
		UpdatedAt:   convertTimestampToRFC3339(pb.GetUpdatedAt()),
	}
}
