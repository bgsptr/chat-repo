package model

type Room struct {
	ID     string    `json:"id"`
	Name   string `json:"group_name"`
	CreatedAt string `json:"created_at"`
	OwnerUsername string `json:"owner_username"`
	Clients map[string]*Client `json:"clients"`
}