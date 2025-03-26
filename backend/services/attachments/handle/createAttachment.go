package handle

import (
	"backend/protos/gen/go/attachments"
	"bytes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
)

func (s *serverAPI) CreateAttachment(stream attachments.AttachmentsProvider_CreateAttachmentServer) error {
	var metadata *attachments.AttachmentMetadata
	var fileData bytes.Buffer

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return status.Errorf(codes.Internal, "failed to receive chunk: %v", err)
		}

		switch data := req.Data.(type) {
		case *attachments.CreateAttachmentRequest_Metadata:
			if metadata != nil {
				return status.Error(codes.InvalidArgument, "metadata already received")
			}
			metadata = data.Metadata
		case *attachments.CreateAttachmentRequest_Chunk:
			if metadata == nil {
				return status.Error(codes.InvalidArgument, "metadata must be sent first")
			}
			if _, err = fileData.Write(data.Chunk); err != nil {
				return status.Errorf(codes.Internal, "failed to write chunk: %v", err)
			}
		}
	}

	if metadata == nil {
		return status.Error(codes.InvalidArgument, "metadata is required")
	}
	if metadata.FileName == "" {
		return status.Error(codes.InvalidArgument, "file_name is required")
	}
	if metadata.MimeType == "" {
		return status.Error(codes.InvalidArgument, "mime_type is required")
	}
	if metadata.AttachmentTypeId <= 0 {
		return status.Error(codes.InvalidArgument, "attachment_type_id must be positive")
	}

	resp, err := s.service.CreateAttachment(stream.Context(), metadata, bytes.NewReader(fileData.Bytes()))
	if err != nil {
		return status.Errorf(codes.Internal, "failed to create attachment: %v", err)
	}

	return stream.SendAndClose(resp)
}
