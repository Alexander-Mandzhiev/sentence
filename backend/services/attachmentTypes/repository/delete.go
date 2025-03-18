package repository

import (
	"context"
	"fmt"
	"log/slog"
)

func (r *Repository) Delete(ctx context.Context, id int32) error {
	op := "repository.Delete"
	r.logger.Info("Deleting direction", slog.String("op", op), slog.Int64("id", int64(id)))

	query := `DELETE FROM attachment_types WHERE id = $1`
	result, err := r.db.Exec(ctx, query, id)
	if err != nil {
		r.logger.Error("Failed to delete direction", slog.String("op", op), slog.String("error", err.Error()))
		return fmt.Errorf("%s: %w", op, err)
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		r.logger.Warn("Direction not found for deletion", slog.String("op", op), slog.Int64("id", int64(id)))
		return ErrAttachmentTypesNotFound
	}

	r.logger.Info("Direction deleted", slog.String("op", op), slog.Int64("id", int64(id)))
	return nil
}
