package model

type Message struct {
	From Client
	RoomID string
	Content string
}