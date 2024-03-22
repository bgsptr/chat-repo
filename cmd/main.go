package main

import (
	"github.com/labstack/echo"
	"groupservice/handler"
)

func main() {
	e := echo.New()

	h := handler.NewGroupChatHandler()

	go h.Run()

	e.GET("/websocket", handler.WebsocketGroupHandler)

	e.Start(":8002")
}