package repository

import (
	"backend/protos/gen/go/sentences"
	"context"
	"fmt"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log/slog"
	"time"
)

func (r *Repository) List(ctx context.Context) ([]*sentences.SentenceResponse, error) {
	op := "repository.List"
	r.logger.Info("Listing sentences", slog.String("op", op))

	query := `SELECT id, status_id, name, phone, department_id, problem, solution, 
               created_at, direction_id, implementor_id, priority_id, encouragement FROM sentences`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		r.logger.Error("Failed to list sentences", slog.String("op", op), slog.String("error", err.Error()))
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	var sentencesList []*sentences.SentenceResponse
	for rows.Next() {
		var sentence sentences.SentenceResponse
		var createdAt time.Time

		err = rows.Scan(&sentence.Id, &sentence.StatusId, &sentence.Name, &sentence.Phone, &sentence.DepartmentId, &sentence.Problem,
			&sentence.Solution, &createdAt, &sentence.DirectionId, &sentence.ImplementorId, &sentence.PriorityId, &sentence.Encouragement)
		if err != nil {
			r.logger.Error("Failed to scan sentence", slog.String("op", op), slog.String("error", err.Error()))
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		sentence.CreatedAt = timestamppb.New(createdAt)
		sentencesList = append(sentencesList, &sentence)
	}

	if err = rows.Err(); err != nil {
		r.logger.Error("Error iterating over rows", slog.String("op", op), slog.String("error", err.Error()))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	r.logger.Info("Sentences listed", slog.String("op", op), slog.Int("count", len(sentencesList)))
	return sentencesList, nil
}
