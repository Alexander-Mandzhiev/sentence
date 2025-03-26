package service

import (
	"backend/protos/gen/go/attachments"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"io"
	"log/slog"
	"os"
	"path/filepath"
)

func (s *Service) CreateAttachment(ctx context.Context, metadata *attachments.AttachmentMetadata, file io.Reader) (*attachments.AttachmentResponse, error) {
	const op = "Service.CreateAttachment"
	logger := s.logger.With(slog.String("op", op))

	if metadata == nil {
		return nil, status.Error(codes.InvalidArgument, "metadata is required")
	}
	if metadata.FileName == "" {
		return nil, status.Error(codes.InvalidArgument, "file name is required")
	}
	if metadata.MimeType == "" {
		return nil, status.Error(codes.InvalidArgument, "mime type is required")
	}
	if metadata.AttachmentTypeId <= 0 {
		return nil, status.Error(codes.InvalidArgument, "attachment type ID must be positive")
	}

	filePath := filepath.Join(s.mediaDir, metadata.FileName)

	fileSize, err := s.saveFile(filePath, file)
	if err != nil {
		logger.Error("failed to save file", slog.String("error", err.Error()))
		return nil, status.Errorf(codes.Internal, "failed to save file")
	}

	attachment := &attachments.AttachmentResponse{
		AttachmentTypeId: metadata.AttachmentTypeId,
		FileName:         metadata.FileName,
		FilePath:         filePath,
		MimeType:         metadata.MimeType,
		FileSize:         fileSize,
		CreatedAt:        timestamppb.Now(),
	}

	created, err := s.provider.Create(ctx, attachment)
	if err != nil {
		if removeErr := os.Remove(filePath); removeErr != nil {
			logger.Error("failed to cleanup file after DB error",
				slog.String("error", removeErr.Error()))
		}
		logger.Error("failed to create attachment in DB", slog.String("error", err.Error()))
		return nil, status.Errorf(codes.Internal, "failed to create attachment")
	}

	logger.Info("attachment created successfully", slog.Int("id", int(created.Id)))
	return created, nil
}
