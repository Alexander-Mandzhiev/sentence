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
	"fmt"
	"log/slog"
	"sync"
)

func (s *Service) validateAndGetDependencies(ctx context.Context, req *sentences.CreateSentenceRequest) (*models.SentenceDependencies, error) {
	const op = "SentenceService.validateAndGetDependencies"
	log := s.logger.With(slog.String("op", op))

	var deps models.SentenceDependencies
	var err error
	var wg sync.WaitGroup

	wg.Add(3)
	go func() {
		defer wg.Done()
		if deps.Status, err = s.statusService.Get(ctx, &statuses.GetStatusRequest{Id: req.GetStatusId()}); err != nil {
			err = fmt.Errorf("%w: %v", ErrInvalidStatus, err)
		}
	}()

	go func() {
		defer wg.Done()
		if deps.Direction, err = s.directionService.Get(ctx, &directions.GetDirectionRequest{Id: req.GetDirectionId()}); err != nil {
			err = fmt.Errorf("%w: %v", ErrInvalidDirection, err)
		}
	}()

	go func() {
		defer wg.Done()
		if deps.Priority, err = s.priorityService.Get(ctx, &priorities.GetPrioritiesRequest{Id: req.GetPriorityId()}); err != nil {
			err = fmt.Errorf("%w: %v", ErrInvalidPriority, err)
		}
	}()

	// Проверяем опциональные зависимости
	if req.GetDepartmentId() != 0 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if deps.Department, err = s.departmentService.Get(ctx, &departments.GetDepartmentRequest{Id: req.GetDepartmentId()}); err != nil {
				err = fmt.Errorf("%w: %v", ErrInvalidDepartment, err)
			}
		}()
	}

	if req.GetImplementorId() != 0 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if deps.Implementor, err = s.implementorService.Get(ctx, &implementors.GetImplementorRequest{Id: req.GetImplementorId()}); err != nil {
				err = fmt.Errorf("%w: %v", ErrInvalidImplementor, err)
			}
		}()
	}

	wg.Wait()
	if err != nil {
		log.Warn("validation failed", "error", err)
		return nil, err
	}

	return &deps, nil
}
