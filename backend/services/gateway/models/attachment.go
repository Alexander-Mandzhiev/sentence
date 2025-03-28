package models

import (
	"backend/protos/gen/go/attachment_types"
	"backend/protos/gen/go/attachments"
)

type Attachment struct {
	ID             int32          `json:"id"`
	FileName       string         `json:"file_name"`
	FilePath       string         `json:"file_path,omitempty"`
	MimeType       string         `json:"mime_type"`
	FileSize       int64          `json:"file_size"`
	CreatedAt      string         `json:"created_at"` // RFC3339 format
	AttachmentType AttachmentType `json:"attachment_type,omitempty"`
}

func AttachmentFromProto(
	pb *attachments.AttachmentResponse,
	typeInfo *attachment_types.AttachmentTypeResponse,
) *Attachment {
	return &Attachment{
		ID:        pb.GetId(),
		FileName:  pb.GetFileName(),
		FilePath:  pb.GetFilePath(),
		MimeType:  pb.GetMimeType(),
		FileSize:  pb.GetFileSize(),
		CreatedAt: convertTimestampToRFC3339(pb.GetCreatedAt()),
		AttachmentType: AttachmentType{
			Id:          typeInfo.GetId(),
			Name:        typeInfo.GetName(),
			Description: typeInfo.GetDescription(),
		},
	}
}
