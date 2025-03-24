package repository

import (
	"backend/protos/gen/go/history"
	"context"
	"fmt"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log/slog"
	"time"
)

func (r *Repository) ListBySentence(ctx context.Context, req *history.ListBySentenceRequest) (*history.HistoryListResponse, error) {
	const op = "repository.HistoryRepository.ListBySentence"
	query := `SELECT id, sentence_id, status_id, created_at, reason, details FROM histories WHERE sentence_id = $1 ORDER BY created_at DESC`
	rows, err := r.db.Query(ctx, query, req.SentenceId)

	if err != nil {
		r.logger.Error("failed to list histories by sentence", slog.String("op", op), slog.String("error", err.Error()))
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	var histories []*history.HistoryResponse
	for rows.Next() {
		var h history.HistoryResponse
		var createdAt time.Time

		err = rows.Scan(&h.Id, &h.SentenceId, &h.StatusId, &createdAt, &h.Reason, &h.Details)
		if err != nil {
			r.logger.Error("failed to scan history row", slog.String("op", op), slog.String("error", err.Error()))
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		h.CreatedAt = timestamppb.New(createdAt)
		histories = append(histories, &h)
	}

	if err = rows.Err(); err != nil {
		r.logger.Error("error after iterating history rows", slog.String("op", op), slog.String("error", err.Error()))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if len(histories) == 0 {
		return nil, ErrHistoriesNotFound
	}

	return &history.HistoryListResponse{Data: histories}, nil
}
