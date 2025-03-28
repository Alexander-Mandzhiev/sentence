package sentences_service

import (
	"backend/protos/gen/go/history"
	"backend/protos/gen/go/sentences"
	"backend/protos/gen/go/statuses"
	"backend/services/gateway/models"
	"context"
	"fmt"
	"log/slog"
	"time"
)

func (s *Service) Update(ctx context.Context, req *models.UpdateSentenceRequest) (*models.EnrichedSentence, error) {
	const op = "SentenceService.Update"
	log := s.logger.With(slog.String("op", op), slog.Int64("id", req.GetId()))

	currentSentence, err := s.sentenceClient.Get(ctx, &sentences.GetSentenceRequest{Id: req.GetId()})
	if err != nil {
		log.Error("failed to get current sentence", "error", err)
		return nil, fmt.Errorf("%w: %v", ErrSentenceNotFound, err)
	}

	updatedSentence, err := s.sentenceClient.Update(ctx, req.SentenceResponse)
	if err != nil {
		log.Error("failed to update sentence", "error", err)
		return nil, err
	}

	if currentSentence.GetStatusId() != updatedSentence.GetStatusId() {
		if _, err = s.historyService.Create(ctx, &history.CreateHistoryRequest{
			SentenceId: updatedSentence.GetId(),
			StatusId:   updatedSentence.GetStatusId(),
			Reason:     req.Reason,
			Details:    req.Details,
		}); err != nil {
			log.Warn("failed to create history record for status change", "error", err)
		}
	}

	deps, err := s.getDependencies(ctx, updatedSentence)
	if err != nil {
		log.Warn("failed to get dependencies", "error", err)
		return nil, err
	}

	enriched := &models.EnrichedSentence{
		ID:            updatedSentence.GetId(),
		Status:        deps.Status.Name,
		Name:          updatedSentence.GetName(),
		Phone:         updatedSentence.GetPhone(),
		Department:    deps.Department.Name,
		Problem:       updatedSentence.GetProblem(),
		Solution:      updatedSentence.GetSolution(),
		CreatedAt:     updatedSentence.GetCreatedAt().AsTime().Format(time.RFC3339),
		Direction:     deps.Direction.Name,
		Implementor:   deps.Implementor.Name,
		Priority:      deps.Priority.Name,
		Encouragement: updatedSentence.GetEncouragement(),
	}

	if historyResp, err := s.historyService.ListBySentence(ctx, &history.ListBySentenceRequest{SentenceId: updatedSentence.GetId()}); err == nil {
		historyItems := make([]*models.History, 0, len(historyResp.GetData()))
		for _, item := range historyResp.GetData() {
			status, err := s.statusService.Get(ctx, &statuses.GetStatusRequest{Id: item.GetStatusId()})
			if err != nil {
				log.Warn("failed to get status name for history item", "status_id", item.GetStatusId(), "error", err)
				continue
			}

			historyItems = append(historyItems, &models.History{
				Id:        int(item.GetId()),
				Status:    status.GetName(),
				CreatedAt: item.GetCreatedAt().AsTime().Format(time.RFC3339),
				Reason:    item.GetReason(),
				Details:   item.GetDetails(),
			})
		}
		enriched.History = historyItems
	}

	return enriched, nil
}
