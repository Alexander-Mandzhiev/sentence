package service

import (
	"backend/protos/gen/go/attachments"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log/slog"
	"time"
)

func (s *Service) Create(ctx context.Context, req *attachments.CreateAttachmentRequest) (*attachments.AttachmentResponse, error) {
	op := "service.Create"
	s.logger.Info("Creating attachment", slog.String("op", op), slog.Any("request", req))

	if req == nil || req.FileName == "" || req.FilePath == "" || req.AttachmentTypeId <= 0 {
		s.logger.Error("Invalid input data", slog.String("op", op), slog.Any("request", req))
		return nil, status.Error(codes.InvalidArgument, "invalid input data: request is nil or required fields are empty/invalid")
	}

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
