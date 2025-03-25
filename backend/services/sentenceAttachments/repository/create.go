package repository

import (
	"backend/protos/gen/go/sentences_attachments"
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgconn"
	"log/slog"
)

func (r *Repository) Create(ctx context.Context, link *sentences_attachments.SentenceAttachmentResponse) error {
	const op = "repository.Create"
	log := r.logger.With(slog.String("op", op), slog.Int64("sentence_id", link.SentenceId), slog.Int64("attachment_id", int64(link.AttachmentId)))

	query := `INSERT INTO sentences_attachments (sentence_id, attachment_id) VALUES ($1, $2)`

	_, err := r.db.Exec(ctx, query, link.SentenceId, link.AttachmentId)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				log.Warn("link already exists")
				return ErrLinkAlreadyExists
			}
		}
		log.Error("failed to create link", "error", err)
		return fmt.Errorf("failed to create link: %w", err)
	}

	log.Debug("link created successfully")
	return nil
}
