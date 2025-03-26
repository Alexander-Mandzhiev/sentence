package handle

import (
	"backend/protos/gen/go/attachments"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
)

func (s *serverAPI) DownloadFile(req *attachments.GetAttachmentRequest, stream attachments.AttachmentsProvider_DownloadFileServer) error {
	if req == nil || req.Id <= 0 {
		return status.Error(codes.InvalidArgument, "id must be positive")
	}

	file, metadata, err := s.service.DownloadFile(stream.Context(), req.Id)
	if err != nil {
		return status.Errorf(codes.Internal, "failed to download file: %v", err)
	}
	defer file.Close()

	if err = stream.Send(&attachments.DownloadFileResponse{
		Data: &attachments.DownloadFileResponse_Metadata{Metadata: metadata},
	}); err != nil {
		return status.Errorf(codes.Internal, "failed to send metadata: %v", err)
	}

	buf := make([]byte, 32*1024)
	for {
		var n int
		n, err = file.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return status.Errorf(codes.Internal, "failed to read file chunk: %v", err)
		}

		if err = stream.Send(&attachments.DownloadFileResponse{
			Data: &attachments.DownloadFileResponse_Chunk{Chunk: buf[:n]},
		}); err != nil {
			return status.Errorf(codes.Internal, "failed to send file chunk: %v", err)
		}
	}

	return nil
}
