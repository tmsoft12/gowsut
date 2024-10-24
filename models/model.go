package models

type Ticket struct {
	ID     int    `json:"id"`
	Ticket int    `json:"ticket"`
	Type   string `json:"type"`
}
