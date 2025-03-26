package service

import (
	"backend/protos/gen/go/attachments"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"log/slog"
	"os"
)

func (s *Service) DownloadFile(ctx context.Context, id int32) (io.ReadCloser, *attachments.FileMetadata, error) {
	const op = "Service.DownloadFile"
	logger := s.logger.With(slog.String("op", op), slog.Int("id", int(id)))

	attachment, err := s.provider.Get(ctx, id)
	if err != nil {
		logger.Error("failed to get attachment metadata", slog.String("error", err.Error()))
		return nil, nil, status.Errorf(codes.Internal, "failed to get attachment")
	}

	file, err := os.Open(attachment.FilePath)
	if err != nil {
		logger.Error("failed to open file", slog.String("filePath", attachment.FilePath), slog.String("error", err.Error()))
		return nil, nil, status.Errorf(codes.Internal, "failed to open file")
	}

	metadata := &attachments.FileMetadata{
		FileName: attachment.FileName,
		MimeType: attachment.MimeType,
		FileSize: attachment.FileSize,
	}

	return file, metadata, nil
}
