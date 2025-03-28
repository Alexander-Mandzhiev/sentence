package repository

import (
	"backend/protos/gen/go/sentences"
	"context"
	"fmt"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log/slog"
	"time"
)

func (r *Repository) List(ctx context.Context, limit, offset int32) ([]*sentences.SentenceResponse, int32, error) {
	const op = "repository.List"
	log := r.logger.With(slog.String("op", op))

	// Проверка параметров пагинации
	if limit < 0 || offset < 0 {
		log.Error("некорректные параметры пагинации",
			slog.Int("limit", int(limit)),
			slog.Int("offset", int(offset)),
		)
		return nil, 0, fmt.Errorf("%w: некорректные параметры пагинации", ErrInvalidInput)
	}

	if limit == 0 {
		limit = 50
	}

	var totalCount int32
	countQuery := `SELECT COUNT(*) FROM sentences`
	err := r.db.QueryRow(ctx, countQuery).Scan(&totalCount)
	if err != nil {
		log.Error("ошибка при получении общего количества предложений", "error", err)
		return nil, 0, fmt.Errorf("%s: %w", op, err)
	}

	query := `SELECT id, status_id, name, phone, department_id, problem, 
        solution, created_at, direction_id, implementor_id, priority_id, encouragement
        FROM sentences
        LIMIT $1 OFFSET $2`

	rows, err := r.db.Query(ctx, query, limit, offset)
	if err != nil {
		log.Error("ошибка при получении списка предложений", "error", err)
		return nil, 0, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	var sentencesList []*sentences.SentenceResponse
	for rows.Next() {
		var (
			sentence  sentences.SentenceResponse
			createdAt time.Time
		)
		if err = rows.Scan(
			&sentence.Id, &sentence.StatusId, &sentence.Name, &sentence.Phone,
			&sentence.DepartmentId, &sentence.Problem, &sentence.Solution, &createdAt,
			&sentence.DirectionId, &sentence.ImplementorId, &sentence.PriorityId,
			&sentence.Encouragement,
		); err != nil {
			log.Error("ошибка при сканировании предложения", "error", err)
			return nil, 0, fmt.Errorf("%s: %w", op, err)
		}
		sentence.CreatedAt = timestamppb.New(createdAt)
		sentencesList = append(sentencesList, &sentence)
	}

	if err = rows.Err(); err != nil {
		log.Error("ошибка при обработке результатов", "error", err)
		return nil, 0, fmt.Errorf("%s: %w", op, err)
	}

	log.Debug("список предложений получен",
		"count", len(sentencesList),
		"total", totalCount,
	)
	return sentencesList, totalCount, nil
}
