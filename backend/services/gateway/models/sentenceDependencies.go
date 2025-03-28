package models

import (
	"backend/protos/gen/go/departments"
	"backend/protos/gen/go/directions"
	"backend/protos/gen/go/implementors"
	"backend/protos/gen/go/priorities"
	"backend/protos/gen/go/statuses"
)

type SentenceDependencies struct {
	Status      *statuses.StatusResponse
	Direction   *directions.DirectionResponse
	Priority    *priorities.PrioritiesResponse
	Department  *departments.DepartmentResponse
	Implementor *implementors.ImplementorResponse
}
