package repository

import (
	"backend/protos/gen/go/directions"
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log/slog"
)

func (r *Repository) Get(ctx context.Context, id int) (*directions.DirectionResponse, error) {
	op := "repository.Get"
	r.logger.Info("Getting direction", slog.String("op", op), slog.Int("id", id))

	query := `SELECT id, name, description, is_active FROM directions WHERE id = $1`
	row := r.db.QueryRow(ctx, query, id)

	var direction directions.DirectionResponse
	err := row.Scan(&direction.Id, &direction.Name, &direction.Description, &direction.IsActive)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			r.logger.Warn("Direction not found", slog.String("op", op), slog.Int("id", id))
			return nil, ErrDirectionsNotFound
		}
		r.logger.Error("Failed to get direction", slog.String("op", op), slog.String("error", err.Error()))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	r.logger.Info("Direction retrieved", slog.String("op", op), slog.Any("direction", direction))
	return &direction, nil
}
