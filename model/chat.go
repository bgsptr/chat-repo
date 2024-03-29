package model

type Chat struct {
	From *Client `json:"from"`
	To map[string]*Client `json:"to"`
	Content string `json:"content"`
	CreatedAt string `json:"created_at"`
}