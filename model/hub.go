package model

type Hub struct {
	// Register requests from the clients.
	Register chan *Client

	// Unregister requests from clients.
	Unregister chan *Client

	Broadcast chan []byte
}