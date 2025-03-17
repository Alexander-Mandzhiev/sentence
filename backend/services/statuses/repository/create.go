package repository

import (
	"backend/protos/gen/go/statuses"
	"context"
	"fmt"
	"log/slog"
)

func (r *Repository) Create(ctx context.Context, status *statuses.StatusResponse) (int, error) {
	op := "repository.Create"
	r.logger.Info("Creating status", slog.String("op", op), slog.Any("status", status))

	query := `INSERT INTO statuses (name, description) VALUES ($1, $2) RETURNING id`
	var id int
	err := r.db.QueryRow(ctx, query, status.Name, status.Description).Scan(&id)
	if err != nil {
		r.logger.Error("Failed to create status", slog.String("op", op), slog.String("error", err.Error()))
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	r.logger.Info("Status created", slog.String("op", op), slog.Int("id", id))
	return id, nil
}
