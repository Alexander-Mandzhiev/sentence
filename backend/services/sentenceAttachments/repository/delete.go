package repository

import (
	"context"
	"fmt"
	"log/slog"
)

func (r *Repository) Delete(ctx context.Context, sentenceID int64, attachmentID int32) error {
	const op = "repository.Delete"
	log := r.logger.With(slog.String("op", op), slog.Int64("sentence_id", sentenceID), slog.Int64("attachment_id", int64(attachmentID)))
	query := `DELETE FROM sentences_attachments WHERE sentence_id = $1 AND attachment_id = $2`

	cmdTag, err := r.db.Exec(ctx, query, sentenceID, attachmentID)
	if err != nil {
		log.Error("failed to delete link", "error", err)
		return fmt.Errorf("failed to delete link: %w", err)
	}

	if cmdTag.RowsAffected() == 0 {
		log.Warn("link not found")
		return ErrLinkNotFound
	}

	log.Debug("link deleted successfully")
	return nil
}
