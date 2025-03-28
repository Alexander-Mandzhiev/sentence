package sentences_service

import (
	"backend/protos/gen/go/history"
	"backend/protos/gen/go/sentences_attachments"
	"backend/services/gateway/models"
	"context"
	"fmt"
	"log/slog"
	"time"
)

func (s *Service) Create(ctx context.Context, req *models.CreateSentenceRequest) (*models.EnrichedSentence, error) {
	const op = "SentenceService.Create"
	log := s.logger.With(slog.String("op", op))
	deps, err := s.validateAndGetDependencies(ctx, req.CreateSentenceRequest)
	if err != nil {
		return nil, err
	}

	sentenceResp, err := s.sentenceClient.Create(ctx, req.CreateSentenceRequest)
	if err != nil {
		log.Error("failed to create sentence", "error", err)
		return nil, fmt.Errorf("%w: %v", ErrSentenceCreation, err)
	}

	if len(req.AttachmentIDs) > 0 {
		for _, attachID := range req.AttachmentIDs {
			if _, err = s.sentenceAttachService.Create(ctx, &sentences_attachments.CreateSentenceAttachmentRequest{
				SentenceId:   sentenceResp.GetId(),
				AttachmentId: attachID,
			}); err != nil {
				log.Warn("failed to link attachment",
					"attachment_id", attachID,
					"error", err)
			}
		}
	}

	if _, err = s.historyService.Create(ctx, &history.CreateHistoryRequest{
		SentenceId: sentenceResp.GetId(),
		StatusId:   req.GetStatusId(),
		Reason:     "Initial creation",
		Details:    "Sentence was created",
	}); err != nil {
		log.Warn("failed to create history record", "error", err)
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

	if len(req.AttachmentIDs) > 0 {
		attachs := make([]*models.EnrichedAttachment, 0, len(req.AttachmentIDs))
		for _, attachID := range req.AttachmentIDs {
			if attach, err := s.attachmentService.Get(ctx, attachID); err == nil {
				enrichedAttach := &models.EnrichedAttachment{
					ID:       attach.ID,
					FileName: attach.FileName,
					MimeType: attach.MimeType,
					FileType: attach.AttachmentType.Name,
					FilePath: attach.FilePath,
				}

				attachs = append(attachs, enrichedAttach)
			}
		}
		enriched.Attachments = attachs
	}

	return enriched, nil
}
