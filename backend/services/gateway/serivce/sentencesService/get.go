package sentences_service

import (
	"backend/protos/gen/go/history"
	"backend/protos/gen/go/sentences"
	"backend/protos/gen/go/sentences_attachments"
	"backend/protos/gen/go/statuses"
	"backend/services/gateway/models"
	"context"
	"fmt"
	"log/slog"
	"time"
)

func (s *Service) Get(ctx context.Context, id int64) (*models.EnrichedSentence, error) {
	const op = "SentenceService.Get"
	log := s.logger.With(slog.String("op", op), slog.Int64("id", id))
	sentenceResp, err := s.sentenceClient.Get(ctx, &sentences.GetSentenceRequest{Id: id})
	if err != nil {
		log.Error("failed to get sentence", "error", err)
		return nil, fmt.Errorf("%w: %v", ErrSentenceNotFound, err)
	}

	deps, err := s.getDependencies(ctx, sentenceResp)
	if err != nil {
		log.Warn("failed to get dependencies", "error", err)
		return nil, err
	}

	enriched := &models.EnrichedSentence{
		ID:            sentenceResp.GetId(),
		Status:        deps.Status.Name,
		Name:          sentenceResp.GetName(),
		Phone:         sentenceResp.GetPhone(),
		Department:    deps.Department.Name,
		Problem:       sentenceResp.GetProblem(),
		Solution:      sentenceResp.GetSolution(),
		CreatedAt:     sentenceResp.GetCreatedAt().AsTime().Format(time.RFC3339),
		Direction:     deps.Direction.Name,
		Implementor:   deps.Implementor.Name,
		Priority:      deps.Priority.Name,
		Encouragement: sentenceResp.GetEncouragement(),
	}

	links, err := s.sentenceAttachService.ListBySentence(ctx, &sentences_attachments.ListBySentenceRequest{SentenceId: id})
	if err == nil && len(links.GetData()) > 0 {
		attachments := make([]*models.EnrichedAttachment, 0, len(links.GetData()))
		for _, link := range links.GetData() {
			if attach, err := s.attachmentService.Get(ctx, link.GetAttachmentId()); err == nil {
				enrichedAttach := &models.EnrichedAttachment{
					ID:       attach.ID,
					FileName: attach.FileName,
					MimeType: attach.MimeType,
					FileType: attach.AttachmentType.Name,
					FilePath: attach.FilePath,
				}

				attachments = append(attachments, enrichedAttach)
			}
		}
		enriched.Attachments = attachments
	}

	if historyResp, err := s.historyService.ListBySentence(ctx, &history.ListBySentenceRequest{SentenceId: id}); err == nil {
		historyItems := make([]*models.History, 0, len(historyResp.GetData()))
		for _, item := range historyResp.GetData() {
			// Получаем название статуса из сервиса статусов
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
