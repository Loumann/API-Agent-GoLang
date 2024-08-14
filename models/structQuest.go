package models

type Quest struct {
	AgentId string `json:"agentid" db:"agentid"`
	Quest   string `json:"quest" db:"quest"`
}
