package service

import (
	"groupservice/config"
	"groupservice/model"
	"groupservice/repository"
	"log"
	"time"

	// "github.com/go-playground/validator/v10"
	"github.com/gorilla/websocket"
)

type MessageService struct {
	MessageRepository *repository.MessageRepository
	Validator *config.CustomValidator
}

func NewMessageService(msg *MessageRepository, v *config.CustomValidator) *MessageService {
	return &MessageService{
		MessageRepository: msg,
		Validator: v,
	}
}

type Client struct {
	*model.Client
}

func NewClient(client *model.Client) *Client {
	return &Client {
		Client: client,
	}
}

func (cl Client) ReadMessageService(hub *model.Hub) {
	defer func() {
		hub.Unregister <- cl.Client
		cl.Conn.Close()
	}()

	for {
		_, m, err := cl.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		msg := &model.Message{
			Content:  string(m),
			CreatedAt: time.Now().Local().String(),
			// RoomID:   cl.RoomID,
			// Username: cl.Username,
		}

		hub.Broadcast <- msg
	}
}

func (cl Client) WriteMessageService() {
	defer func() {
		cl.Conn.Close()
	}()

	for {
		message, ok := <- cl.Message

		if !ok {
			return
		}

		cl.Conn.WriteJSON(message)
	}
}

func (m *MessageService) SendMessage(sender interface{}, receiver interface{}, msg string) error {
	err := m.Validator.TryValidate(sender, receiver, msg)
	if err != nil {
		return err
	}

	err, resp := m.MessageRepository.Send(sender, receiver, msg)
	if err != nil {
		return nil
	}

	return nil
}
