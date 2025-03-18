package repository

import (
	"backend/protos/gen/go/sentences"
	"context"
	"fmt"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log/slog"
	"time"
)

func (r *Repository) Create(ctx context.Context, sentence *sentences.SentenceResponse) error {
	op := "repository.Create"
	r.logger.Info("Creating sentence", slog.String("op", op), slog.Any("sentence", sentence))

	query := `INSERT INTO sentences (status_id, name, phone, department_id, problem, solution, direction_id, implementor_id, priority_id, encouragement)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id, created_at`

	var id int64
	var createdAt time.Time

	err := r.db.QueryRow(ctx, query, sentence.StatusId, sentence.Name, sentence.Phone, sentence.DepartmentId, sentence.Problem,
		sentence.Solution, sentence.DirectionId, sentence.ImplementorId, sentence.PriorityId, sentence.Encouragement).Scan(&id, &createdAt)
	if err != nil {
		r.logger.Error("Failed to create sentence", slog.String("op", op), slog.String("error", err.Error()))
		return fmt.Errorf("%s: %w", op, err)
	}

	sentence.Id = id
	sentence.CreatedAt = timestamppb.New(createdAt)

	r.logger.Info("Sentence created", slog.String("op", op), slog.Int64("id", id))
	return nil
}
