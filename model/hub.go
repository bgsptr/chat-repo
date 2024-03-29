package model

type Hub struct {

	// map[cl.chatId]*Chat
	Room map[string]*Chat
	// Register requests from the clients.
	Register chan *Client

	// Unregister requests from clients.
	Unregister chan *Client

	Broadcast chan *Message
}