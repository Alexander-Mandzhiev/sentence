package repository

import (
	"backend/protos/gen/go/priorities"
	"context"
	"fmt"
	"log/slog"
)

func (r *Repository) Update(ctx context.Context, priority *priorities.PrioritiesResponse) error {
	op := "repository.Update"
	r.logger.Info("Updating priority", slog.String("op", op), slog.Any("priority", priority))

	query := `UPDATE priorities SET name = $1, description = $2 WHERE id = $3`
	result, err := r.db.Exec(ctx, query, priority.Name, priority.Description, priority.Id)
	if err != nil {
		r.logger.Error("Failed to update priority", slog.String("op", op), slog.String("error", err.Error()))
		return fmt.Errorf("%s: %w", op, err)
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		r.logger.Warn("Priority not found for update", slog.String("op", op), slog.Int("id", int(priority.Id)))
		return ErrPrioritiesNotFound
	}

	r.logger.Info("Priority updated", slog.String("op", op), slog.Any("priority", priority))
	return nil
}
