package sentences_service

import (
	"backend/protos/gen/go/sentences"
	"backend/services/gateway/models"
	"context"
	"log/slog"
	"time"
)

func (s *Service) List(ctx context.Context, limit, offset int32) ([]*models.EnrichedSentence, int32, error) {
	const op = "SentenceService.List"
	log := s.logger.With(slog.String("op", op))

	listResp, err := s.sentenceClient.List(ctx, &sentences.ListSentencesRequest{Limit: limit, Offset: offset})
	if err != nil {
		log.Error("failed to list sentences", "error", err)
		return nil, 0, err
	}

	enrichedSentences := make([]*models.EnrichedSentence, 0, len(listResp.GetData()))
	for _, sentence := range listResp.GetData() {
		var deps *models.SentenceDependencies
		deps, err = s.getDependencies(ctx, sentence)
		if err != nil {
			log.Warn("failed to get dependencies for sentence", "sentence_id", sentence.GetId(), "error", err)
			continue
		}

		enriched := &models.EnrichedSentence{
			ID:            sentence.GetId(),
			Status:        deps.Status.Name,
			Name:          sentence.GetName(),
			Phone:         sentence.GetPhone(),
			Department:    deps.Department.Name,
			Problem:       sentence.GetProblem(),
			Solution:      sentence.GetSolution(),
			CreatedAt:     sentence.GetCreatedAt().AsTime().Format(time.RFC3339),
			Direction:     deps.Direction.Name,
			Implementor:   deps.Implementor.Name,
			Priority:      deps.Priority.Name,
			Encouragement: sentence.GetEncouragement(),
		}
		enrichedSentences = append(enrichedSentences, enriched)
	}

	return enrichedSentences, listResp.GetTotalCount(), nil
}
