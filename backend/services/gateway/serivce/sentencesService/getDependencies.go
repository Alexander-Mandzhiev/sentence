package sentences_service

import (
	"backend/protos/gen/go/departments"
	"backend/protos/gen/go/directions"
	"backend/protos/gen/go/implementors"
	"backend/protos/gen/go/priorities"
	"backend/protos/gen/go/sentences"
	"backend/protos/gen/go/statuses"
	"backend/services/gateway/models"
	"context"
	"sync"
)

func (s *Service) getDependencies(ctx context.Context, sentence *sentences.SentenceResponse) (*models.SentenceDependencies, error) {
	var deps models.SentenceDependencies
	var err error
	var wg sync.WaitGroup

	wg.Add(3)
	go func() {
		defer wg.Done()
		deps.Status, err = s.statusService.Get(ctx, &statuses.GetStatusRequest{Id: sentence.GetStatusId()})
	}()

	go func() {
		defer wg.Done()
		deps.Direction, err = s.directionService.Get(ctx, &directions.GetDirectionRequest{Id: sentence.GetDirectionId()})
	}()

	go func() {
		defer wg.Done()
		deps.Priority, err = s.priorityService.Get(ctx, &priorities.GetPrioritiesRequest{Id: sentence.GetPriorityId()})
	}()

	if sentence.GetDepartmentId() != 0 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			deps.Department, err = s.departmentService.Get(ctx, &departments.GetDepartmentRequest{Id: sentence.GetDepartmentId()})
		}()
	}

	if sentence.GetImplementorId() != 0 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			deps.Implementor, err = s.implementorService.Get(ctx, &implementors.GetImplementorRequest{Id: sentence.GetImplementorId()})
		}()
	}

	wg.Wait()
	if err != nil {
		return nil, err
	}

	return &deps, nil
}
