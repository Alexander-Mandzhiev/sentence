package attachment_service

import (
	"backend/protos/gen/go/attachments"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"log/slog"
)

func (s *Service) DownloadFile(ctx context.Context, id int32) (io.ReadCloser, *attachments.FileMetadata, error) {
	const op = "attachment_service.DownloadFile"
	log := s.logger.With(slog.String("op", op), slog.Int("id", int(id)))

	if id <= 0 {
		return nil, nil, status.Error(codes.InvalidArgument, "invalid attachment ID")
	}

	stream, err := s.client.DownloadFile(ctx, &attachments.GetAttachmentRequest{Id: id})
	if err != nil {
		log.Error("failed to create download stream", slog.String("error", err.Error()))
		return nil, nil, status.Errorf(codes.Internal, "failed to download file")
	}

	resp, err := stream.Recv()
	if err != nil {
		log.Error("failed to receive metadata", slog.String("error", err.Error()))
		return nil, nil, status.Errorf(codes.Internal, "failed to get file metadata")
	}

	metadata := resp.GetMetadata()
	if metadata == nil {
		return nil, nil, status.Error(codes.Internal, "missing file metadata")
	}

	pr, pw := io.Pipe()

	go func() {
		defer pw.Close()

		for {
			resp, err = stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Error("failed to receive chunk", slog.String("error", err.Error()))
				pw.CloseWithError(status.Errorf(codes.Internal, "failed to receive file chunk"))
				return
			}

			if chunk := resp.GetChunk(); chunk != nil {
				if _, err = pw.Write(chunk); err != nil {
					log.Error("failed to write chunk", slog.String("error", err.Error()))
					pw.CloseWithError(status.Errorf(codes.Internal, "failed to write file chunk"))
					return
				}
			}
		}
	}()

	return pr, metadata, nil
}
