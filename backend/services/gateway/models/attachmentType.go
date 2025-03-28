package models

type AttachmentType struct {
	Id          int32  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}
