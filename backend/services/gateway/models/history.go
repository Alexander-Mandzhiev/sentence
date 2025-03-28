package models

type History struct {
	Id        int    `json:"id"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
	Reason    string `json:"reason"`
	Details   string `json:"details"`
}
