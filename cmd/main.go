package main

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/calender-service/config"
	"golang.org/x/calender-service/internal/service"
)

func main() {

	db, err := database.NewGorm()
	if err != nil {
		log.Println("cant instatiate db")
	}

	config.RunPostgresMigrate(db)
	cv := config.NewValidator()
	repo := repository.NewRepository()
	e := service.NewEventService(repo, cv.Validate)
	r := gin.New()

	r.Run()
}

func Route() {

}