package repository

import (
	"backend/protos/gen/go/attachment_types"
	"context"
	"fmt"
	"log/slog"
)

func (r *Repository) List(ctx context.Context) ([]*attachment_types.AttachmentTypeResponse, error) {
	op := "repository.List"
	r.logger.Info("Listing attachment types", slog.String("op", op))

	query := `SELECT id, name, description FROM attachment_types`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		r.logger.Error("Failed to list attachment types", slog.String("op", op), slog.String("error", err.Error()))
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	var attachmentTypes []*attachment_types.AttachmentTypeResponse
	for rows.Next() {
		var attachmentType attachment_types.AttachmentTypeResponse
		if err = rows.Scan(&attachmentType.Id, &attachmentType.Name, &attachmentType.Description); err != nil {
			r.logger.Error("Failed to scan attachment type", slog.String("op", op), slog.String("error", err.Error()))
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		attachmentTypes = append(attachmentTypes, &attachmentType)
	}

	if err = rows.Err(); err != nil {
		r.logger.Error("Error iterating over attachment types", slog.String("op", op), slog.String("error", err.Error()))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	r.logger.Info("Attachment types listed", slog.String("op", op), slog.Int("count", len(attachmentTypes)))
	return attachmentTypes, nil
}
