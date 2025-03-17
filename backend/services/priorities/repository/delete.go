package repository

import (
	"context"
	"fmt"
	"log/slog"
)

func (r *Repository) Delete(ctx context.Context, id int) error {
	op := "repository.Delete"
	r.logger.Info("Deleting priority", slog.String("op", op), slog.Int("id", id))

	query := `DELETE FROM priorities WHERE id = $1`
	result, err := r.db.Exec(ctx, query, id)
	if err != nil {
		r.logger.Error("Failed to delete priority", slog.String("op", op), slog.String("error", err.Error()))
		return fmt.Errorf("%s: %w", op, err)
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		r.logger.Warn("Priority not found for deletion", slog.String("op", op), slog.Int("id", id))
		return ErrPrioritiesNotFound
	}

	r.logger.Info("Priority deleted", slog.String("op", op), slog.Int("id", id))
	return nil
}
