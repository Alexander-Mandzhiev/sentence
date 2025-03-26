package service

import (
	"backend/protos/gen/go/attachments"
	"context"
	"log/slog"
	"os"
)

type AttachmentProvider interface {
	Create(ctx context.Context, attachment *attachments.AttachmentResponse) (*attachments.AttachmentResponse, error)
	Get(ctx context.Context, id int32) (*attachments.AttachmentResponse, error)
	Delete(ctx context.Context, id int32) (string, error)
	List(ctx context.Context, limit int32, offset int32) ([]*attachments.AttachmentResponse, error)
}

type Service struct {
	logger   *slog.Logger
	provider AttachmentProvider
	mediaDir string
}

func New(provider AttachmentProvider, logger *slog.Logger, mediaDir string) *Service {
	op := "service.New"
	if err := os.MkdirAll(mediaDir, 0755); err != nil {
		logger.Error("failed to create media directory", slog.String("error", err.Error()), slog.String("op", op))
		return nil
	}

	if provider == nil {
		logger.Error("attachment provider is nil", slog.String("op", op))
		return nil
	}

	logger.Info("Service initialized", slog.String("op", op))
	return &Service{provider: provider, logger: logger, mediaDir: mediaDir}
}
