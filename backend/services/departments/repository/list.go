package repository

import (
	"backend/protos/gen/go/departments"
	"context"
	"fmt"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log/slog"
	"time"
)

func (r *Repository) List(ctx context.Context) ([]*departments.DepartmentResponse, error) {
	op := "repository.List"
	r.logger.Info("Listing departments", slog.String("op", op))

	query := `SELECT id, name, description, is_active, created_at, updated_at FROM departments`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		r.logger.Error("Failed to list departments", slog.String("op", op), slog.String("error", err.Error()))
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	var departmentsList []*departments.DepartmentResponse
	for rows.Next() {
		var deeps departments.DepartmentResponse
		var createdAt, updatedAt time.Time

		if err = rows.Scan(&deeps.Id, &deeps.Name, &deeps.Description, &deeps.IsActive, &createdAt, &updatedAt); err != nil {
			r.logger.Error("Failed to scan department", slog.String("op", op), slog.String("error", err.Error()))
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		deeps.CreatedAt = timestamppb.New(createdAt)
		deeps.UpdatedAt = timestamppb.New(updatedAt)

		departmentsList = append(departmentsList, &deeps)
	}

	if err = rows.Err(); err != nil {
		r.logger.Error("Error iterating over rows", slog.String("op", op), slog.String("error", err.Error()))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	r.logger.Info("Departments listed", slog.String("op", op), slog.Int("count", len(departmentsList)))
	return departmentsList, nil
}
