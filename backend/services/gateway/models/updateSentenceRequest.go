package models

import "backend/protos/gen/go/sentences"

type UpdateSentenceRequest struct {
	*sentences.SentenceResponse
	Reason  string `json:"reason"`
	Details string `json:"details"`
}
