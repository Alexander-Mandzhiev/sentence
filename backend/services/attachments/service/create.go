package service

import (
	"backend/protos/gen/go/attachments"
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log/slog"
	"time"
)

func (s *Service) Create(ctx context.Context, req *attachments.CreateAttachmentRequest) (*attachments.AttachmentResponse, error) {
	op := "service.Create"
	s.logger.Info("Creating attachment", slog.String("op", op), slog.Any("request", req))

	attachment := &attachments.AttachmentResponse{
		AttachmentTypeId: req.AttachmentTypeId,
		FileName:         req.FileName,
		FilePath:         req.FilePath,
		MimeType:         req.MimeType,
		FileSize:         req.FileSize,
		CreatedAt:        timestamppb.New(time.Now()),
	}

	resp, err := s.provider.Create(ctx, attachment)
	if err != nil {
		s.logger.Error("Failed to create attachment", slog.String("op", op), slog.String("error", err.Error()))
		return nil, err
	}

	s.logger.Info("Attachment created", slog.String("op", op), slog.Any("attachment", resp))
	return resp, nil
}
