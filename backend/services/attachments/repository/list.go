package repository

import (
	"backend/protos/gen/go/attachments"
	"context"
	"fmt"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log/slog"
	"time"
)

func (r *Repository) List(ctx context.Context) ([]*attachments.AttachmentResponse, error) {
	op := "repository.List"
	r.logger.Info("Listing attachments", slog.String("op", op))
	query := `SELECT id, attachment_type_id, file_name, file_path, mime_type, file_size, created_at	FROM attachments`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		r.logger.Error("Failed to list attachments", slog.String("op", op), slog.String("error", err.Error()))
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	var attachmentsList []*attachments.AttachmentResponse
	for rows.Next() {
		var (
			att       attachments.AttachmentResponse
			createdAt time.Time
		)
		if err = rows.Scan(&att.Id, &att.AttachmentTypeId, &att.FileName, &att.FilePath, &att.MimeType, &att.FileSize, &createdAt); err != nil {
			r.logger.Error("Failed to scan attachment", slog.String("op", op), slog.String("error", err.Error()))
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		att.CreatedAt = timestamppb.New(createdAt)
		attachmentsList = append(attachmentsList, &att)
	}

	if err = rows.Err(); err != nil {
		r.logger.Error("Error iterating over attachments", slog.String("op", op), slog.String("error", err.Error()))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	r.logger.Info("Attachments listed", slog.String("op", op), slog.Int("count", len(attachmentsList)))
	return attachmentsList, nil
}
