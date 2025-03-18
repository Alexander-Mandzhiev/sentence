package repository

import (
	"backend/protos/gen/go/attachment_types"
	"context"
	"fmt"
	"log/slog"
)

func (r *Repository) Create(ctx context.Context, attachmentType *attachment_types.AttachmentTypeResponse) (int32, error) {
	op := "repository.Create"
	r.logger.Info("Creating direction", slog.String("op", op), slog.Any("attachment type", attachmentType))

	query := `INSERT INTO attachment_types (name, description) VALUES ($1, $2) RETURNING id`
	var id int32
	err := r.db.QueryRow(ctx, query, attachmentType.Name, attachmentType.Description).Scan(&id)
	if err != nil {
		r.logger.Error("Failed to create direction", slog.String("op", op), slog.String("error", err.Error()))
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	r.logger.Info("Direction created", slog.String("op", op), slog.Int64("id", int64(id)))
	return id, nil
}
