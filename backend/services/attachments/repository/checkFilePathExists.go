package repository

import (
	"context"
	"fmt"
	"log/slog"
)

func (r *Repository) checkFilePathExists(ctx context.Context, filePath string) (bool, error) {
	const op = "repository.checkFilePathExists"
	query := "SELECT EXISTS(SELECT 1 FROM attachments WHERE file_path = $1)"
	var exists bool
	err := r.db.QueryRow(ctx, query, filePath).Scan(&exists)

	if err != nil {
		r.logger.Error("failed to check file path existence", slog.String("op", op), slog.String("error", err.Error()))
		return false, fmt.Errorf("%s: %w", op, err)
	}
	return exists, nil
}
