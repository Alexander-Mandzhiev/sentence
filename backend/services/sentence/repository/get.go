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

	query := `SELECT id, status_id, name, phone, department_id, problem, solution, 
               created_at, direction_id, implementor_id, priority_id, encouragement
        FROM sentences WHERE id = $1`

	row := r.db.QueryRow(ctx, query, id)

	var sentence sentences.SentenceResponse
	var createdAt time.Time

	err := row.Scan(&sentence.Id, &sentence.StatusId, &sentence.Name, &sentence.Phone, &sentence.DepartmentId, &sentence.Problem,
		&sentence.Solution, &createdAt, &sentence.DirectionId, &sentence.ImplementorId, &sentence.PriorityId, &sentence.Encouragement)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			r.logger.Warn("Sentence not found", slog.String("op", op), slog.Int64("id", id))
			return nil, ErrSentenceNotFound
		}
		r.logger.Error("Failed to get sentence", slog.String("op", op), slog.String("error", err.Error()))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	sentence.CreatedAt = timestamppb.New(createdAt)

	r.logger.Info("Sentence retrieved", slog.String("op", op), slog.Int64("id", id))
	return &sentence, nil
}
