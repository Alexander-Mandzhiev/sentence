package repository

import (
	"backend/protos/gen/go/departments"
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log/slog"
	"time"
)

func (r *Repository) Get(ctx context.Context, id int) (*departments.DepartmentResponse, error) {
	op := "repository.Get"
	r.logger.Info("Getting department", slog.String("op", op), slog.Int("id", id))

	query := `SELECT id, name, description, is_active, created_at, updated_at FROM departments WHERE id = $1`
	row := r.db.QueryRow(ctx, query, id)

	var deeps departments.DepartmentResponse
	var (
		createdAt time.Time
		updatedAt time.Time
	)

	err := row.Scan(&deeps.Id, &deeps.Name, &deeps.Description, &deeps.IsActive, &createdAt, &updatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			r.logger.Warn("Department not found", slog.String("op", op), slog.Int("id", id))
			return nil, ErrDepartmentsNotFound
		}
		r.logger.Error("Failed to get department", slog.String("op", op), slog.String("error", err.Error()))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if createdAt.IsZero() || updatedAt.IsZero() {
		r.logger.Warn("Invalid timestamp values", slog.String("op", op))
		return nil, fmt.Errorf("%s: invalid timestamp", op)
	}

	deeps.CreatedAt = timestamppb.New(createdAt)
	deeps.UpdatedAt = timestamppb.New(updatedAt)

	r.logger.Info("Department retrieved", slog.String("op", op), slog.Int("id department", int(deeps.Id)))

	return &deeps, nil
}
