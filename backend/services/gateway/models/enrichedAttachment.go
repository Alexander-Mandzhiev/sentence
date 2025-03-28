package models

type EnrichedAttachment struct {
	ID       int32  `json:"id"`
	FileName string `json:"file_name"`
	FileType string `json:"file_type"`
	MimeType string `json:"mime_type"`
	FilePath string `json:"file_path"`
}
