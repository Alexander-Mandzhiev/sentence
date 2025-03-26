package repository

import (
	"backend/protos/gen/go/history"
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log/slog"
	"time"
)

func (r *Repository) Get(ctx context.Context, req *history.GetHistoryRequest) (*history.HistoryResponse, error) {
	const op = "repository.HistoryRepository.Get"
	var resp history.HistoryResponse
	var createdAt time.Time
	query := `SELECT id, sentence_id, status_id, created_at, reason, details FROM history WHERE id = $1`
	err := r.db.QueryRow(ctx, query, req.Id).Scan(&resp.Id, &resp.SentenceId, &resp.StatusId, &createdAt, &resp.Reason, &resp.Details)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrHistoriesNotFound
		}
		r.logger.Error("failed to get history record", slog.String("op", op), slog.Int64("id", int64(req.Id)), slog.String("error", err.Error()))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	resp.CreatedAt = timestamppb.New(createdAt)
	return &resp, nil
}
