package model

type Message struct {
	Content string `json:"content"`
	CreatedAt string `json:"created_at"`
	RoomID string `json:"room_id"`
}