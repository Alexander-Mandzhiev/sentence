package repository

import (
	"backend/protos/gen/go/attachments"
	"context"
	"fmt"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log/slog"
	"time"
)

func (r *Repository) List(ctx context.Context, limit, offset int32) ([]*attachments.AttachmentResponse, error) {
	const op = "repository.List"
	logger := r.logger.With(slog.String("op", op))

	if limit <= 0 || limit > 100 {
		limit = 50
	}

	if offset < 0 {
		offset = 0
	}

	query := `SELECT id, attachment_type_id, file_name, file_path, mime_type, file_size, created_at FROM attachments ORDER BY created_at DESC LIMIT $1 OFFSET $2`

	rows, err := r.db.Query(ctx, query, limit, offset)
	if err != nil {
		logger.Error("failed to list attachments", slog.String("error", err.Error()))
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
			logger.Error("failed to scan attachment", slog.String("error", err.Error()))
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		att.CreatedAt = timestamppb.New(createdAt)
		attachmentsList = append(attachmentsList, &att)
	}

	if err = rows.Err(); err != nil {
		logger.Error("error iterating attachments", slog.String("error", err.Error()))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return attachmentsList, nil
}
