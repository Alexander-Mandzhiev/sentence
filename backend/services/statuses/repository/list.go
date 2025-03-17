package repository

import (
	"backend/protos/gen/go/statuses"
	"context"
	"fmt"
	"log/slog"
)

func (r *Repository) List(ctx context.Context) ([]*statuses.StatusResponse, error) {
	op := "repository.Statuses"
	r.logger.Info("Listing statuses", slog.String("op", op))

	query := `SELECT id, name, description FROM statuses`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		r.logger.Error("Failed to list statuses", slog.String("op", op), slog.String("error", err.Error()))
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	var statusesList []*statuses.StatusResponse
	for rows.Next() {
		var status statuses.StatusResponse
		if err = rows.Scan(&status.Id, &status.Name, &status.Description); err != nil {
			r.logger.Error("Failed to scan status", slog.String("op", op), slog.String("error", err.Error()))
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		statusesList = append(statusesList, &status)
	}

	if err = rows.Err(); err != nil {
		r.logger.Error("Error iterating over rows", slog.String("op", op), slog.String("error", err.Error()))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	r.logger.Info("Statuses listed", slog.String("op", op), slog.Int("count", len(statusesList)))
	return statusesList, nil
}
