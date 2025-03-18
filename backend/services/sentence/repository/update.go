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

	query := `UPDATE sentences 
        SET status_id = $1, name = $2, phone = $3, department_id = $4, 
            problem = $5, solution = $6, direction_id = $7, implementor_id = $8, 
            priority_id = $9, encouragement = $10
        WHERE id = $11`

	result, err := r.db.Exec(ctx, query, sentence.StatusId, sentence.Name, sentence.Phone, sentence.DepartmentId, sentence.Problem,
		sentence.Solution, sentence.DirectionId, sentence.ImplementorId, sentence.PriorityId, sentence.Encouragement, sentence.Id)
	if err != nil {
		r.logger.Error("Failed to update sentence", slog.String("op", op), slog.String("error", err.Error()))
		return fmt.Errorf("%s: %w", op, err)
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		r.logger.Warn("Sentence not found for update", slog.String("op", op), slog.Int64("id", sentence.Id))
		return ErrSentenceNotFound
	}

	r.logger.Info("Sentence updated", slog.String("op", op), slog.Int64("id", sentence.Id))
	return nil
}
