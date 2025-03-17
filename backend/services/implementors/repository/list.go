package repository

import (
	"backend/protos/gen/go/implementors"
	"context"
	"fmt"
	"log/slog"
)

func (r *Repository) List(ctx context.Context) ([]*implementors.ImplementorResponse, error) {
	op := "repository.List"
	r.logger.Info("Listing implementors", slog.String("op", op))

	query := `SELECT id, name, email, is_active FROM implementors`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		r.logger.Error("Failed to list implementors", slog.String("op", op), slog.String("error", err.Error()))
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	var implementorsList []*implementors.ImplementorResponse
	for rows.Next() {
		var implementor implementors.ImplementorResponse
		if err = rows.Scan(&implementor.Id, &implementor.Name, &implementor.Email, &implementor.IsActive); err != nil {
			r.logger.Error("Failed to scan implementor", slog.String("op", op), slog.String("error", err.Error()))
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		implementorsList = append(implementorsList, &implementor)
	}

	if err = rows.Err(); err != nil {
		r.logger.Error("Error iterating over rows", slog.String("op", op), slog.String("error", err.Error()))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	r.logger.Info("Implementors listed", slog.String("op", op), slog.Int("count", len(implementorsList)))
	return implementorsList, nil
}
