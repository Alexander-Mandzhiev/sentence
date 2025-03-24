package repository

import (
	"context"
	"fmt"
	"log/slog"
)

func (r *Repository) Delete(ctx context.Context, id int32) error {
	op := "repository.Delete"
	r.logger.Info("Deleting attachment", slog.String("op", op), slog.Int64("id", int64(id)))

	if id <= 0 {
		r.logger.Error("Invalid attachment ID", slog.String("op", op), slog.Int64("id", int64(id)))
		return fmt.Errorf("%s: invalid attachment ID", op)
	}

	query := `DELETE FROM attachments WHERE id = $1`

	result, err := r.db.Exec(ctx, query, id)
	if err != nil {
		r.logger.Error("Failed to delete attachment", slog.String("op", op), slog.String("error", err.Error()))
		return fmt.Errorf("%s: %w", op, err)
	}

	if result.RowsAffected() == 0 {
		r.logger.Error("Attachment not found for deletion", slog.String("op", op), slog.Int64("id", int64(id)))
		return ErrAttachmentNotFound
	}

	r.logger.Info("Attachment deleted", slog.String("op", op), slog.Int64("id", int64(id)))
	return nil
}
