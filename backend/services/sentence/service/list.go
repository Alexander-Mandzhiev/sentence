package service

import (
	"backend/protos/gen/go/sentences"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

func (s *Service) List(ctx context.Context, limit, offset int32) (*sentences.SentencesListResponse, error) {
	const op = "service.List"
	log := s.logger.With(slog.String("op", op))

	// Валидация параметров пагинации
	if limit < 0 || offset < 0 {
		log.Error("некорректные параметры пагинации",
			slog.Int("limit", int(limit)),
			slog.Int("offset", int(offset)),
		)
		return nil, status.Error(codes.InvalidArgument, "некорректные параметры пагинации")
	}

	if limit == 0 {
		limit = 50
	}

	sentencesList, totalCount, err := s.provider.List(ctx, limit, offset)
	if err != nil {
		log.Error("ошибка при получении списка предложений", "error", err)
		return nil, status.Error(codes.Internal, "не удалось получить список предложений")
	}

	log.Debug("список предложений получен",
		"count", len(sentencesList),
		"total", totalCount,
	)
	return &sentences.SentencesListResponse{
		Data:       sentencesList,
		TotalCount: totalCount,
	}, nil
}
