package repository

import (
	"backend/protos/gen/go/priorities"
	"context"
	"fmt"
	"log/slog"
)

func (r *Repository) List(ctx context.Context) ([]*priorities.PrioritiesResponse, error) {
	op := "repository.List"
	r.logger.Info("Listing priorities", slog.String("op", op))

	query := `SELECT id, name, description FROM priorities`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		r.logger.Error("Failed to list priorities", slog.String("op", op), slog.String("error", err.Error()))
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	var prioritiesList []*priorities.PrioritiesResponse
	for rows.Next() {
		var priority priorities.PrioritiesResponse
		if err = rows.Scan(&priority.Id, &priority.Name, &priority.Description); err != nil {
			r.logger.Error("Failed to scan priority", slog.String("op", op), slog.String("error", err.Error()))
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		prioritiesList = append(prioritiesList, &priority)
	}

	if err = rows.Err(); err != nil {
		r.logger.Error("Error iterating over rows", slog.String("op", op), slog.String("error", err.Error()))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	r.logger.Info("Priorities listed", slog.String("op", op), slog.Int("count", len(prioritiesList)))
	return prioritiesList, nil
}
