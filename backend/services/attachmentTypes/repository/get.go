package repository

import (
	"backend/protos/gen/go/attachment_types"
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log/slog"
)

func (r *Repository) Get(ctx context.Context, id int32) (*attachment_types.AttachmentTypeResponse, error) {
	op := "repository.Get"
	r.logger.Info("Getting attachment type", slog.String("op", op), slog.Int64("id", int64(id)))

	query := `SELECT id, name, description FROM attachment_types WHERE id = $1`

	var attachmentType attachment_types.AttachmentTypeResponse
	err := r.db.QueryRow(ctx, query, id).Scan(&attachmentType.Id, &attachmentType.Name, &attachmentType.Description)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			r.logger.Error("Attachment type not found", slog.String("op", op), slog.Int64("id", int64(id)))
			return nil, ErrAttachmentTypesNotFound
		}
		r.logger.Error("Failed to get attachment type", slog.String("op", op), slog.String("error", err.Error()))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	r.logger.Info("Attachment type fetched", slog.String("op", op), slog.Any("attachmentType", attachmentType))
	return &attachmentType, nil
}
