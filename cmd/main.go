package main

import (
	"log"
	"net/http"

	// "userservice/dataservice"
	// "userservice/usecase"
	"userservice/config"
	"userservice/dataservice"
	"userservice/dataservice/usersqldb"
	"userservice/handler"
	"userservice/handler/helper"
	"userservice/usecase"

	"github.com/gorilla/mux"
)

func depedencyService() *handler.UserHandler {
	userRepository := dataservice.NewUserData()

	userService := usecase.NewUserService(userRepository)
	helperHandler := &helper.HelperHandler{}

	u := handler.NewUserHandler(userService, helperHandler)

	return u
}

func main() {

	log.Println("running")

	u := depedencyService()

	db, err := usersqldb.NewGorm()
	if err != nil {
		log.Println("cant instatiate db")
	}

	config.RunPostgresMigrate(db)

	// u := handler.NewUserHandler()
	router := mux.NewRouter()

	subRouter := router.PathPrefix("/api/v1").Subrouter()

	subRouter.HandleFunc("/user", u.CreateAccount).Methods("POST")

	http.Handle("/", router)
    http.ListenAndServe(":3000",nil)
}