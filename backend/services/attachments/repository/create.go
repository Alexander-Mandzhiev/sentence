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
	op := "repository.Create"
	r.logger.Info("Creating attachment", slog.String("op", op), slog.Any("attachment", attachment))

	if attachment == nil || attachment.FileName == "" || attachment.FilePath == "" || attachment.AttachmentTypeId <= 0 {
		r.logger.Error("Invalid input data", slog.String("op", op), slog.Any("attachment", attachment))
		return nil, fmt.Errorf("%s: invalid input data", op)
	}

	query := `INSERT INTO attachments (attachment_type_id, file_name, file_path, mime_type, file_size, created_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, created_at`

	var (
		id        int32
		createdAt time.Time
	)

	err := r.db.QueryRow(ctx, query, attachment.AttachmentTypeId, attachment.FileName, attachment.FilePath, attachment.MimeType, attachment.FileSize, attachment.CreatedAt.AsTime()).Scan(&id, &createdAt)
	if err != nil {
		r.logger.Error("Failed to create attachment", slog.String("op", op), slog.String("error", err.Error()))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	attachment.Id = id
	attachment.CreatedAt = timestamppb.New(createdAt)

	r.logger.Info("Attachment created", slog.String("op", op), slog.Any("attachment", attachment))
	return attachment, nil
}
