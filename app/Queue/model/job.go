package model

type IJob struct {
	ID      int    `json:"id"`
	Type    string `json:"type"`
	Payload any    `json:"payload"`
}
