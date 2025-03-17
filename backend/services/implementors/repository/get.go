package repository

import (
	"backend/protos/gen/go/implementors"
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log/slog"
)

func (r *Repository) Get(ctx context.Context, id int) (*implementors.ImplementorResponse, error) {
	op := "repository.Get"
	r.logger.Info("Getting implementor", slog.String("op", op), slog.Int("id", id))

	query := `SELECT id, name, email, is_active FROM implementors WHERE id = $1`
	row := r.db.QueryRow(ctx, query, id)

	var implementor implementors.ImplementorResponse
	err := row.Scan(&implementor.Id, &implementor.Name, &implementor.Email, &implementor.IsActive)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			r.logger.Warn("Implementor not found", slog.String("op", op), slog.Int("id", id))
			return nil, ErrImplementorsNotFound
		}
		r.logger.Error("Failed to get implementor", slog.String("op", op), slog.String("error", err.Error()))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	r.logger.Info("Implementor retrieved", slog.String("op", op), slog.Any("implementor", implementor))
	return &implementor, nil
}
