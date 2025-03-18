package repository

import (
	"backend/protos/gen/go/departments"
	"context"
	"fmt"
	"golang.org/x/exp/slog"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

func (r *Repository) Create(ctx context.Context, department *departments.DepartmentResponse) error {
	op := "repository.Create"
	var createdAt, updatedAt time.Time
	r.logger.Info("Creating department", slog.String("op", op), slog.Any("department", department))

	query := `INSERT INTO departments (name, description, is_active) VALUES ($1, $2, $3) RETURNING id, created_at, updated_at`

	err := r.db.QueryRow(ctx, query, department.Name, department.Description, department.IsActive).Scan(&department.Id, &createdAt, &updatedAt)
	if err != nil {
		r.logger.Error("Failed to create department", slog.String("op", op), slog.String("error", err.Error()))
		return fmt.Errorf("%s: %w", op, err)
	}

	department.CreatedAt = timestamppb.New(createdAt)
	department.UpdatedAt = timestamppb.New(updatedAt)

	r.logger.Info("Department created", slog.String("op", op), slog.Any("department", department))
	return nil
}
