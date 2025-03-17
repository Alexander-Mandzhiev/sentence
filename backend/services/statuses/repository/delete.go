package repository

import (
	"context"
	"fmt"
	"log/slog"
)

func (r *Repository) Delete(ctx context.Context, id int) error {
	op := "repository.Delete"
	r.logger.Info("Deleting status", slog.String("op", op), slog.Int("id", id))

	query := `DELETE FROM statuses WHERE id = $1`
	result, err := r.db.Exec(ctx, query, id)
	if err != nil {
		r.logger.Error("Failed to delete status", slog.String("op", op), slog.String("error", err.Error()))
		return fmt.Errorf("%s: %w", op, err)
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		r.logger.Warn("Status not found for deletion", slog.String("op", op), slog.Int("id", id))
		return ErrStatusNotFound
	}

	r.logger.Info("Status deleted", slog.String("op", op), slog.Int("id", id))
	return nil
}
