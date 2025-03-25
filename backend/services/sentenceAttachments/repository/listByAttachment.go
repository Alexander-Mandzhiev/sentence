package repository

import (
	"backend/protos/gen/go/sentences_attachments"
	"context"
	"fmt"
	"log/slog"
)

func (r *Repository) ListByAttachment(ctx context.Context, attachmentID int32) ([]*sentences_attachments.SentenceAttachmentResponse, error) {
	const op = "repository.ListByAttachment"
	log := r.logger.With(slog.String("op", op), slog.Int64("attachment_id", int64(attachmentID)))
	query := `SELECT sentence_id, attachment_id FROM sentences_attachments WHERE attachment_id = $1`

	rows, err := r.db.Query(ctx, query, attachmentID)
	if err != nil {
		log.Error("failed to list sentences for attachment", "error", err)
		return nil, fmt.Errorf("failed to list sentences: %w", err)
	}
	defer rows.Close()

	var links []*sentences_attachments.SentenceAttachmentResponse
	for rows.Next() {
		var link sentences_attachments.SentenceAttachmentResponse
		if err = rows.Scan(&link.SentenceId, &link.AttachmentId); err != nil {
			log.Error("failed to scan row", "error", err)
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		links = append(links, &link)
	}

	if err = rows.Err(); err != nil {
		log.Error("error during rows iteration", "error", err)
		return nil, fmt.Errorf("rows error: %w", err)
	}

	log.Debug("sentences listed successfully", "count", len(links))
	return links, nil
}
