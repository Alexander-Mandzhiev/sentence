package attachment_service

import (
	"backend/protos/gen/go/attachments"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"log/slog"
)

func (s *Service) CreateAttachment(ctx context.Context, metadata *attachments.AttachmentMetadata, file io.Reader) (*attachments.AttachmentResponse, error) {
	const op = "attachment_service.CreateAttachment"
	log := s.logger.With(slog.String("op", op))
	stream, err := s.client.CreateAttachment(ctx)
	if err != nil {
		log.Error("failed to create stream", slog.String("error", err.Error()))
		return nil, status.Errorf(codes.Internal, "failed to create stream")
	}

	err = stream.Send(&attachments.CreateAttachmentRequest{
		Data: &attachments.CreateAttachmentRequest_Metadata{
			Metadata: metadata,
		},
	})
	if err != nil {
		log.Error("failed to send metadata", slog.String("error", err.Error()))
		return nil, status.Errorf(codes.Internal, "failed to send metadata")
	}

	buf := make([]byte, 32*1024)
	for {
		var n int
		n, err = file.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Error("failed to read file chunk", slog.String("error", err.Error()))
			return nil, status.Errorf(codes.Internal, "failed to read file")
		}

		if err = stream.Send(&attachments.CreateAttachmentRequest{
			Data: &attachments.CreateAttachmentRequest_Chunk{
				Chunk: buf[:n],
			},
		}); err != nil {
			log.Error("failed to send chunk", slog.String("error", err.Error()))
			return nil, status.Errorf(codes.Internal, "failed to send file chunk")
		}
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Error("failed to receive response", slog.String("error", err.Error()))
		return nil, status.Errorf(codes.Internal, "failed to create attachment")
	}

	log.Info("attachment created successfully", slog.Int("id", int(resp.Id)))
	return resp, nil
}
