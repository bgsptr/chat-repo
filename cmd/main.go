package main

import (
	"groupservice/config"
	"groupservice/handler"
	"groupservice/middleware"
	"groupservice/repository"
	"groupservice/service"
	"log"

	"github.com/labstack/echo"
)

func main() {
	cv := config.NewValidator()

	msgRepo := repository.NewMessageRepository(db)
	msgSrv := service.NewGroupService(msgRepo)
	m := handler.NewMessageHandler(msgSrv)

	db, err := config.SqlConnection()
	if err != nil {
		log.Println(err)
	}

	groupRepo := repository.NewGroupRepository(db)
	groupSrv := service.NewGroupService(groupRepo)
	hub := handler.NewHub()

	
	h := handler.NewGroupChatHandler(groupSrv, hub)

	config.Migrate(db)
	j := middleware.GetJwtValidate()
	
	e := echo.New()

	e.Use(j.ValidateJWT())

	go h.Run()

	e.GET("/websocket", h.WebsocketGroupHandler)

	e.Start(":8002")
}