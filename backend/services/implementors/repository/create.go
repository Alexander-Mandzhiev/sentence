package repository

import (
	"backend/protos/gen/go/implementors"
	"context"
	"fmt"
	"log/slog"
)

func (r *Repository) Create(ctx context.Context, implementor *implementors.ImplementorResponse) (int, error) {
	op := "repository.Create"
	r.logger.Info("Creating implementor", slog.String("op", op), slog.Any("implementor", implementor))

	query := `INSERT INTO implementors (name, email, is_active) VALUES ($1, $2, $3) RETURNING id`
	var id int
	err := r.db.QueryRow(ctx, query, implementor.Name, implementor.Email, implementor.IsActive).Scan(&id)
	if err != nil {
		r.logger.Error("Failed to create implementor", slog.String("op", op), slog.String("error", err.Error()))
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	r.logger.Info("Implementor created", slog.String("op", op), slog.Int("id", id))
	return id, nil
}
