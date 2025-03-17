package repository

import (
	"backend/protos/gen/go/priorities"
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log/slog"
)

func (r *Repository) Get(ctx context.Context, id int) (*priorities.PrioritiesResponse, error) {
	op := "repository.Get"
	r.logger.Info("Getting priority", slog.String("op", op), slog.Int("id", id))

	query := `SELECT id, name, description FROM priorities WHERE id = $1`
	row := r.db.QueryRow(ctx, query, id)

	var priority priorities.PrioritiesResponse
	err := row.Scan(&priority.Id, &priority.Name, &priority.Description)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			r.logger.Warn("Priority not found", slog.String("op", op), slog.Int("id", id))
			return nil, ErrPrioritiesNotFound
		}
		r.logger.Error("Failed to get priority", slog.String("op", op), slog.String("error", err.Error()))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	r.logger.Info("Priority retrieved", slog.String("op", op), slog.Any("priority", priority))
	return &priority, nil
}
