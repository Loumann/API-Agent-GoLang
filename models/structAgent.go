package models

type Agent struct {
	ID        int    `json:"id" db:"id"`
	AgentName string `json:"agentname" db:"agentname"`
	Status    string `json:"status" db:"status"`
}
