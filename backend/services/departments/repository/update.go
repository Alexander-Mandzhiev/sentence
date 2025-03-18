package repository

import (
	"backend/protos/gen/go/departments"
	"context"
	"fmt"
	"log/slog"
)

func (r *Repository) Update(ctx context.Context, dept *departments.DepartmentResponse) error {
	op := "repository.Update"
	r.logger.Info("Updating department", slog.String("op", op), slog.Any("department", dept))

	query := `UPDATE departments SET name = $1, description = $2, is_active = $3, updated_at = $4 WHERE id = $5`

	_, err := r.db.Exec(ctx, query, dept.Name, dept.Description, dept.IsActive, dept.UpdatedAt.AsTime(), dept.Id)
	if err != nil {
		r.logger.Error("Failed to update department", slog.String("op", op), slog.String("error", err.Error()))
		return fmt.Errorf("%s: %w", op, err)
	}

	r.logger.Info("Department updated", slog.String("op", op))
	return nil
}
