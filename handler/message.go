package handler

import (
	"encoding/json"
	"groupservice/model"
	"groupservice/service"

	"github.com/labstack/echo"
)

type MessageHandler struct {
	MessageService *service.MessageService
}

func NewMessageHandler(srv *service.MessageService) *MessageHandler {
	return &MessageHandler{
        MessageService: srv,
    }
}

func (m *MessageHandler) SendMessage(c echo.Context) error {
	var mg *model.Chat
	d := json.NewDecoder(c.Request().Body)
	err := d.Decode(&mg)
	if err != nil {
		return err
	}

	err = m.MessageService.SendMessage(mg.From, mg.To, mg.Content)
	if err != nil {
		return err
	}
}