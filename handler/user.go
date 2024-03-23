// handler package
package handler

import (
	"context"
	// "encoding/json"
	"log"
	"net/http"
	"time"
	"userservice/handler/helper"
	"userservice/model"
	"userservice/usecase"

	"github.com/sirupsen/logrus"
)

type UserHandler struct {
	Do         *helper.HelperHandler
	UserService *usecase.UserService
	Log        *logrus.Logger
}

func NewUserHandler(userService *usecase.UserService, helperHandler *helper.HelperHandler) *UserHandler {
	return &UserHandler{
		UserService: userService,
		Do:          helperHandler,
	}
}

// func NewUserHandler() *UserHandler {
// 	return &UserHandler{}
// }

func GetStructUser() *model.User {
	return &model.User{}
}

func (u *UserHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "POST" {
		http.Error(w, "400", http.StatusBadRequest)
		return
	}

	user := GetStructUser()
	err := u.Do.DecodeJson(r.Body, user)
	if err != nil {
		log.Println("Failed to decode JSON")
		http.Error(w, "500", http.StatusInternalServerError)
		return
	}

	deadline := time.Now().Add(10 * time.Second)

	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	err = u.UserService.CreateAccount(ctx, user)
	if err != nil {
		log.Println("Failed to create account")
		http.Error(w, "500", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Success create account"))
}