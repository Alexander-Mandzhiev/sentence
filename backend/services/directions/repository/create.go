package repository

import (
	"backend/protos/gen/go/directions"
	"context"
	"fmt"
	"log/slog"
)

func (r *Repository) Create(ctx context.Context, direction *directions.DirectionResponse) (int, error) {
	op := "repository.Create"
	r.logger.Info("Creating direction", slog.String("op", op), slog.Any("direction", direction))

	query := `INSERT INTO directions (name, description, is_active) VALUES ($1, $2, $3) RETURNING id`
	var id int
	err := r.db.QueryRow(ctx, query, direction.Name, direction.Description, direction.IsActive).Scan(&id)
	if err != nil {
		r.logger.Error("Failed to create direction", slog.String("op", op), slog.String("error", err.Error()))
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	r.logger.Info("Direction created", slog.String("op", op), slog.Int("id", id))
	return id, nil
}
