package repository

import (
	"backend/protos/gen/go/directions"
	"context"
	"fmt"
	"log/slog"
)

func (r *Repository) List(ctx context.Context) ([]*directions.DirectionResponse, error) {
	op := "repository.List"
	r.logger.Info("Listing directions", slog.String("op", op))

	query := `SELECT id, name, description, is_active FROM directions`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		r.logger.Error("Failed to list directions", slog.String("op", op), slog.String("error", err.Error()))
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	var directionsList []*directions.DirectionResponse
	for rows.Next() {
		var direction directions.DirectionResponse
		if err = rows.Scan(&direction.Id, &direction.Name, &direction.Description, &direction.IsActive); err != nil {
			r.logger.Error("Failed to scan direction", slog.String("op", op), slog.String("error", err.Error()))
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		directionsList = append(directionsList, &direction)
	}

	if err = rows.Err(); err != nil {
		r.logger.Error("Error iterating over rows", slog.String("op", op), slog.String("error", err.Error()))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	r.logger.Info("Directions listed", slog.String("op", op), slog.Int("count", len(directionsList)))
	return directionsList, nil
}
