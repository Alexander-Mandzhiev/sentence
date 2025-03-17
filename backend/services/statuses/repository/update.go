package repository

import (
	"backend/protos/gen/go/statuses"
	"context"
	"fmt"
	"log/slog"
)

func (r *Repository) Update(ctx context.Context, status *statuses.StatusResponse) error {
	op := "repository.Update"
	r.logger.Info("Updating status", slog.String("op", op), slog.Any("status", status))

	query := `UPDATE statuses SET name = $1, description = $2 WHERE id = $3`
	result, err := r.db.Exec(ctx, query, status.Name, status.Description, status.Id)
	if err != nil {
		r.logger.Error("Failed to update status", slog.String("op", op), slog.String("error", err.Error()))
		return fmt.Errorf("%s: %w", op, err)
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		r.logger.Warn("Status not found for update", slog.String("op", op), slog.Int("id", int(status.Id)))
		return ErrStatusNotFound
	}

	r.logger.Info("Status updated", slog.String("op", op), slog.Any("status", status))
	return nil
}
