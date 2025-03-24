package repository

import (
	"backend/protos/gen/go/sentences"
	"context"
	"fmt"
	"log/slog"
)

func (r *Repository) Update(ctx context.Context, sentence *sentences.SentenceResponse) error {
	op := "repository.Update"
	r.logger.Info("Updating sentence", slog.String("op", op), slog.Any("sentence", sentence))

	// Проверка входных данных
	if sentence == nil || sentence.Id <= 0 || sentence.StatusId <= 0 || sentence.Name == "" || sentence.DirectionId <= 0 || sentence.PriorityId <= 0 {
		r.logger.Error("Invalid input data",
			slog.String("op", op),
			slog.Any("sentence", sentence),
			slog.Bool("is_nil", sentence == nil),
			slog.Bool("is_id_invalid", sentence != nil && sentence.Id <= 0),
			slog.Bool("is_status_id_invalid", sentence != nil && sentence.StatusId <= 0),
			slog.Bool("is_name_empty", sentence != nil && sentence.Name == ""),
			slog.Bool("is_direction_id_invalid", sentence != nil && sentence.DirectionId <= 0),
			slog.Bool("is_priority_id_invalid", sentence != nil && sentence.PriorityId <= 0),
		)
		return fmt.Errorf("%s: invalid input data", op)
	}

	query := `UPDATE sentences SET 
		status_id = $1, name = $2, phone = $3, department_id = $4, problem = $5, solution = $6, 
		direction_id = $7, implementor_id = $8, priority_id = $9, encouragement = $10 
		WHERE id = $11`

	result, err := r.db.Exec(ctx, query,
		sentence.StatusId, sentence.Name, sentence.Phone, sentence.DepartmentId, sentence.Problem,
		sentence.Solution, sentence.DirectionId, sentence.ImplementorId, sentence.PriorityId,
		sentence.Encouragement, sentence.Id,
	)
	if err != nil {
		r.logger.Error("Failed to update sentence", slog.String("op", op), slog.String("error", err.Error()))
		return fmt.Errorf("%s: %w", op, err)
	}

	if result.RowsAffected() == 0 {
		r.logger.Error("Sentence not found for update", slog.String("op", op), slog.Int64("id", sentence.Id))
		return ErrSentenceNotFound
	}

	r.logger.Info("Sentence updated", slog.String("op", op), slog.Any("sentence", sentence))
	return nil
}
