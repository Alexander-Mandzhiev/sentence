package repository

import (
	"context"
	"fmt"
	"log/slog"
)

func (r *Repository) Delete(ctx context.Context, id int64) error {
	op := "repository.Delete"
	r.logger.Info("Deleting sentence", slog.String("op", op), slog.Int64("id", id))

	if id <= 0 {
		r.logger.Error("Invalid input data", slog.String("op", op), slog.Int64("id", id))
		return fmt.Errorf("%s: invalid input data", op)
	}

	query := `DELETE FROM sentences WHERE id = $1`

	result, err := r.db.Exec(ctx, query, id)
	if err != nil {
		r.logger.Error("Failed to delete sentence", slog.String("op", op), slog.String("error", err.Error()))
		return fmt.Errorf("%s: %w", op, err)
	}

	if result.RowsAffected() == 0 {
		r.logger.Error("Sentence not found for deletion", slog.String("op", op), slog.Int64("id", id))
		return ErrSentenceNotFound
	}

	r.logger.Info("Sentence deleted", slog.String("op", op), slog.Int64("id", id))
	return nil
}
