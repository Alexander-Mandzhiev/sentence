package repository

import (
	"backend/protos/gen/go/directions"
	"context"
	"fmt"
	"log/slog"
)

func (r *Repository) Update(ctx context.Context, direction *directions.DirectionResponse) error {
	op := "repository.Update"
	r.logger.Info("Updating direction", slog.String("op", op), slog.Any("direction", direction))

	query := `UPDATE directions SET name = $1, description = $2, is_active = $3 WHERE id = $4`
	result, err := r.db.Exec(ctx, query, direction.Name, direction.Description, direction.IsActive, direction.Id)
	if err != nil {
		r.logger.Error("Failed to update direction", slog.String("op", op), slog.String("error", err.Error()))
		return fmt.Errorf("%s: %w", op, err)
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		r.logger.Warn("Direction not found for update", slog.String("op", op), slog.Int("id", int(direction.Id)))
		return ErrDirectionsNotFound
	}

	r.logger.Info("Direction updated", slog.String("op", op), slog.Any("direction", direction))
	return nil
}
