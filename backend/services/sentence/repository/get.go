package repository

import (
	"backend/protos/gen/go/sentences"
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log/slog"
	"time"
)

func (r *Repository) Get(ctx context.Context, id int64) (*sentences.SentenceResponse, error) {
	op := "repository.Get"
	r.logger.Info("Getting sentence", slog.String("op", op), slog.Int64("id", id))

	if id <= 0 {
		r.logger.Error("Invalid input data", slog.String("op", op), slog.Int64("id", id))
		return nil, fmt.Errorf("%s: invalid input data", op)
	}

	query := `SELECT id, status_id, name, phone, department_id, problem, solution, created_at, direction_id, implementor_id, priority_id, encouragement 
		FROM sentences WHERE id = $1`

	var (
		sentence  sentences.SentenceResponse
		createdAt time.Time
	)

	err := r.db.QueryRow(ctx, query, id).Scan(
		&sentence.Id, &sentence.StatusId, &sentence.Name, &sentence.Phone, &sentence.DepartmentId,
		&sentence.Problem, &sentence.Solution, &createdAt, &sentence.DirectionId, &sentence.ImplementorId,
		&sentence.PriorityId, &sentence.Encouragement,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			r.logger.Error("Sentence not found", slog.String("op", op), slog.Int64("id", id))
			return nil, ErrSentenceNotFound
		}
		r.logger.Error("Failed to get sentence", slog.String("op", op), slog.String("error", err.Error()))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	sentence.CreatedAt = timestamppb.New(createdAt)

	r.logger.Info("Sentence fetched", slog.String("op", op), slog.Any("sentence", sentence))
	return &sentence, nil
}
