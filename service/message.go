package service

import (
	"groupservice/model"

	"golang.org/x/text/message"
)

type MessageService struct {

}

func NewClient() *model.Client {
	return &model.Client {

	}
}

func (cl *model.Client) ReadMessageService(hub *model.Hub) {
	defer func() {
		hub.Unregister <- cl
		cl.Close()
	}()

	for {
		_, m, err := cl.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		msg := &Message{
			Content:  string(m),
			RoomID:   cl.RoomID,
			Username: cl.Username,
		}

		hub.Broadcast <- msg
	}
}

func (cl *model.Client) WriteMessageService() {
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
