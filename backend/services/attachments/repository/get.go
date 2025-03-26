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
	const op = "repository.Get"
	logger := r.logger.With(slog.String("op", op), slog.Int("id", int(id)))

	if id <= 0 {
		return nil, fmt.Errorf("%s: invalid id", op)
	}

	query := `SELECT id, attachment_type_id, file_name, file_path, mime_type, file_size, created_at FROM attachments WHERE id = $1`

	var (
		att       attachments.AttachmentResponse
		createdAt time.Time
	)

	err := r.db.QueryRow(ctx, query, id).Scan(&att.Id, &att.AttachmentTypeId, &att.FileName, &att.FilePath, &att.MimeType, &att.FileSize, &createdAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrAttachmentNotFound
		}
		logger.Error("failed to get attachment", slog.String("error", err.Error()))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	att.CreatedAt = timestamppb.New(createdAt)
	return &att, nil
}
