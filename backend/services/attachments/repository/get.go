package repository

import (
	"backend/protos/gen/go/attachments"
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log/slog"
	"time"
)

func (r *Repository) Get(ctx context.Context, id int32) (*attachments.AttachmentResponse, error) {
	op := "repository.Get"
	r.logger.Info("Getting attachment", slog.String("op", op), slog.Int64("id", int64(id)))

	if id <= 0 {
		r.logger.Error("Invalid attachment ID", slog.String("op", op), slog.Int64("id", int64(id)))
		return nil, fmt.Errorf("%s: invalid attachment ID", op)
	}

	query := `SELECT id, attachment_type_id, file_name, file_path, mime_type, file_size, created_at FROM attachments WHERE id = $1`

	var (
		att       attachments.AttachmentResponse
		createdAt time.Time
	)

	err := r.db.QueryRow(ctx, query, id).Scan(&att.Id, &att.AttachmentTypeId, &att.FileName, &att.FilePath, &att.MimeType, &att.FileSize, &createdAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			r.logger.Error("Attachment not found", slog.String("op", op), slog.Int64("id", int64(id)))
			return nil, ErrAttachmentNotFound
		}
		r.logger.Error("Failed to get attachment", slog.String("op", op), slog.String("error", err.Error()))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	att.CreatedAt = timestamppb.New(createdAt)

	r.logger.Info("Attachment fetched", slog.String("op", op), slog.Any("attachment", att))
	return &att, nil
}
