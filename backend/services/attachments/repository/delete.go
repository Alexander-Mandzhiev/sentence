package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log/slog"
)

func (r *Repository) Delete(ctx context.Context, id int32) (string, error) {
	const op = "repository.Delete"
	logger := r.logger.With(slog.String("op", op), slog.Int("id", int(id)))

	if id <= 0 {
		return "", fmt.Errorf("%s: invalid id", op)
	}

	tx, err := r.db.Begin(ctx)
	if err != nil {
		logger.Error("failed to begin transaction", slog.String("error", err.Error()))
		return "", fmt.Errorf("%s: %w", op, err)
	}
	defer tx.Rollback(ctx)

	var filePath string
	err = tx.QueryRow(ctx, "SELECT file_path FROM attachments WHERE id = $1", id).Scan(&filePath)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return "", ErrAttachmentNotFound
		}
		logger.Error("failed to get file path", slog.String("error", err.Error()))
		return "", fmt.Errorf("%s: %w", op, err)
	}

	result, err := tx.Exec(ctx, "DELETE FROM attachments WHERE id = $1", id)
	if err != nil {
		logger.Error("failed to delete attachment", slog.String("error", err.Error()))
		return "", fmt.Errorf("%s: %w", op, err)
	}

	if result.RowsAffected() == 0 {
		return "", ErrAttachmentNotFound
	}

	if err = tx.Commit(ctx); err != nil {
		logger.Error("failed to commit transaction", slog.String("error", err.Error()))
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return filePath, nil
}
