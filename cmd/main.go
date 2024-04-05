package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"golang.org/x/calender-service/config"
	"golang.org/x/calender-service/db"
	"golang.org/x/calender-service/handler"
	"golang.org/x/calender-service/internal/repository"
	"golang.org/x/calender-service/internal/service"
)
func main() {
	initService()
	r := gin.New()

	r.Run()
}

func initService() {
	db, err := db.NewGorm()
	if err != nil {
		log.Println("cant instatiate db")
	}

	tx := db.Begin()

	getDb, err := db.DB()
	if err != nil {
		return
	}

	config.RunPostgresMigrate(getDb)
	cv := config.NewValidator()
	repo := repository.NewEventRepository(tx)
	e := service.NewEventService(repo, cv.Validate)
	handler.NewEventHandler(e)
}


func Route() {

}