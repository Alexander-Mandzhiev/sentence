package repository

import (
	"backend/protos/gen/go/history"
	"context"
	"fmt"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log/slog"
	"time"
)

func (r *Repository) Create(ctx context.Context, req *history.CreateHistoryRequest) (*history.HistoryResponse, error) {
	const op = "repository.HistoryRepository.Create"
	var id int32
	var createdAt time.Time
	query := `INSERT INTO history (sentence_id, status_id, reason, details) VALUES ($1, $2, $3, $4) RETURNING id, created_at`
	err := r.db.QueryRow(ctx, query, req.SentenceId, req.StatusId, req.Reason, req.Details).Scan(&id, &createdAt)

	if err != nil {
		r.logger.Error("failed to create history record", slog.String("op", op), slog.Int64("sentence_id", req.SentenceId),
			slog.Int64("status_id", int64(req.StatusId)), slog.String("error", err.Error()))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &history.HistoryResponse{
		Id:         id,
		SentenceId: req.SentenceId,
		StatusId:   req.StatusId,
		CreatedAt:  timestamppb.New(createdAt),
		Reason:     req.Reason,
		Details:    req.Details,
	}, nil
}
