package model

import "github.com/gorilla/websocket"

type Client struct {
	Conn     *websocket.Conn
	Username string `json:"username"`
	Password string `json:"password"`
	Message  chan   *Message `json:"message"`
	RoomID   string `json:"room_id"`
}