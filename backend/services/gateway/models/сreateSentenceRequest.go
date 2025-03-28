package models

import "backend/protos/gen/go/sentences"

type CreateSentenceRequest struct {
	*sentences.CreateSentenceRequest
	AttachmentIDs []int32 `json:"attachment_ids"`
}
