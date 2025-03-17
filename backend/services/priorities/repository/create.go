package repository

import (
	"backend/protos/gen/go/priorities"
	"context"
	"fmt"
	"golang.org/x/exp/slog"
)

func (r *Repository) Create(ctx context.Context, priority *priorities.PrioritiesResponse) (int, error) {
	op := "repository.Create"
	r.logger.Info("Creating priority", slog.String("op", op), slog.Any("priority", priority))

	query := `INSERT INTO priorities (name, description) VALUES ($1, $2) RETURNING id`
	var id int
	err := r.db.QueryRow(ctx, query, priority.Name, priority.Description).Scan(&id)
	if err != nil {
		r.logger.Error("Failed to create priority", slog.String("op", op), slog.String("error", err.Error()))
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	r.logger.Info("Priority created", slog.String("op", op), slog.Int("id", id))
	return id, nil
}
