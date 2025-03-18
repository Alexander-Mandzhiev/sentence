package repository

import (
	"backend/protos/gen/go/attachments"
	"context"
	"fmt"
	"log/slog"
)

func (r *Repository) Update(ctx context.Context, attachment *attachments.AttachmentResponse) error {
	op := "repository.Update"
	r.logger.Info("Updating attachment", slog.String("op", op), slog.Any("attachment", attachment))

	query := `UPDATE attachments SET attachment_type_id = $1, file_name = $2, file_path = $3, mime_type = $4, file_size = $5	WHERE id = $6`

	result, err := r.db.Exec(ctx, query, attachment.AttachmentTypeId, attachment.FileName, attachment.FilePath, attachment.MimeType, attachment.FileSize, attachment.Id)
	if err != nil {
		r.logger.Error("Failed to update attachment", slog.String("op", op), slog.String("error", err.Error()))
		return fmt.Errorf("%s: %w", op, err)
	}

	if result.RowsAffected() == 0 {
		r.logger.Error("Attachment not found for update", slog.String("op", op), slog.Int64("id", int64(attachment.Id)))
		return ErrAttachmentNotFound
	}

	r.logger.Info("Attachment updated", slog.String("op", op), slog.Any("attachment", attachment))
	return nil
}
