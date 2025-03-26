package repository

import (
	"backend/protos/gen/go/attachments"
	"context"
	"fmt"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log/slog"
	"time"
)

func (r *Repository) Create(ctx context.Context, attachment *attachments.AttachmentResponse) (*attachments.AttachmentResponse, error) {
	const op = "repository.Create"
	logger := r.logger.With(slog.String("op", op))
	if attachment == nil || attachment.FileName == "" || attachment.FilePath == "" || attachment.AttachmentTypeId <= 0 {
		return nil, fmt.Errorf("%s: invalid input data", op)
	}

	exists, err := r.checkFilePathExists(ctx, attachment.FilePath)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	if exists {
		return nil, fmt.Errorf("%s: file path already exists", op)
	}

	query := `INSERT INTO attachments (attachment_type_id, file_name, file_path, mime_type, file_size, created_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, created_at`

	var (
		id        int32
		createdAt time.Time
	)

	err = r.db.QueryRow(ctx, query, attachment.AttachmentTypeId, attachment.FileName, attachment.FilePath, attachment.MimeType, attachment.FileSize, time.Now()).Scan(&id, &createdAt)

	if err != nil {
		logger.Error("failed to create attachment", slog.String("error", err.Error()))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	attachment.Id = id
	attachment.CreatedAt = timestamppb.New(createdAt)

	return attachment, nil
}
