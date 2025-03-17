package repository

import (
	"backend/protos/gen/go/statuses"
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log/slog"
)

func (r *Repository) Get(ctx context.Context, id int) (*statuses.StatusResponse, error) {
	op := "repository.Get"
	r.logger.Info("Getting status", slog.String("op", op), slog.Int("id", id))

	query := `SELECT id, name, description FROM statuses WHERE id = $1`
	row := r.db.QueryRow(ctx, query, id)

	var status statuses.StatusResponse
	err := row.Scan(&status.Id, &status.Name, &status.Description)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			r.logger.Warn("Status not found", slog.String("op", op), slog.Int("id", id))
			return nil, ErrStatusNotFound
		}
		r.logger.Error("Failed to get status", slog.String("op", op), slog.String("error", err.Error()))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	r.logger.Info("Status retrieved", slog.String("op", op), slog.Any("status", status))
	return &status, nil
}
