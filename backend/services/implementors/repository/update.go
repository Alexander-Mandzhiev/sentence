package repository

import (
	"backend/protos/gen/go/implementors"
	"context"
	"fmt"
	"log/slog"
)

func (r *Repository) Update(ctx context.Context, implementor *implementors.ImplementorResponse) error {
	op := "repository.Update"
	r.logger.Info("Updating implementor", slog.String("op", op), slog.Any("implementor", implementor))

	query := `UPDATE implementors SET name = $1, email = $2, is_active = $3 WHERE id = $4`
	result, err := r.db.Exec(ctx, query, implementor.Name, implementor.Email, implementor.IsActive, implementor.Id)
	if err != nil {
		r.logger.Error("Failed to update implementor", slog.String("op", op), slog.String("error", err.Error()))
		return fmt.Errorf("%s: %w", op, err)
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		r.logger.Warn("Implementor not found for update", slog.String("op", op), slog.Int("id", int(implementor.Id)))
		return ErrImplementorsNotFound
	}

	r.logger.Info("Implementor updated", slog.String("op", op), slog.Any("implementor", implementor))
	return nil
}
