package repository

import (
	"backend/protos/gen/go/attachment_types"
	"context"
	"fmt"
	"log/slog"
)

func (r *Repository) Update(ctx context.Context, attachmentType *attachment_types.AttachmentTypeResponse) error {
	op := "repository.Update"
	r.logger.Info("Updating attachment type", slog.String("op", op), slog.Any("attachmentType", attachmentType))

	query := `UPDATE attachment_types SET name = $1, description = $2 WHERE id = $3`

	result, err := r.db.Exec(ctx, query, attachmentType.Name, attachmentType.Description, attachmentType.Id)
	if err != nil {
		r.logger.Error("Failed to update attachment type", slog.String("op", op), slog.String("error", err.Error()))
		return fmt.Errorf("%s: %w", op, err)
	}

	if result.RowsAffected() == 0 {
		r.logger.Error("Attachment type not found for update", slog.String("op", op), slog.Int64("id", int64(attachmentType.Id)))
		return ErrAttachmentTypesNotFound
	}

	r.logger.Info("Attachment type updated", slog.String("op", op), slog.Any("attachmentType", attachmentType))
	return nil
}
