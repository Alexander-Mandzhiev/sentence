package models

type EnrichedSentence struct {
	ID            int64                 `json:"id"`
	Status        string                `json:"status"`
	Name          string                `json:"name"`
	Phone         string                `json:"phone,omitempty"`
	Department    string                `json:"department,omitempty"`
	Problem       string                `json:"problem,omitempty"`
	Solution      string                `json:"solution,omitempty"`
	CreatedAt     string                `json:"created_at"`
	Direction     string                `json:"direction"`
	Implementor   string                `json:"implementor,omitempty"`
	Priority      string                `json:"priority"`
	Encouragement int32                 `json:"encouragement,omitempty"`
	Attachments   []*EnrichedAttachment `json:"attachments,omitempty"`
	History       []*History            `json:"history,omitempty"`
}
