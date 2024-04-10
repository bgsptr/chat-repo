// handler package
package handler

import (
	"context"
	"encoding/json"
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

	response := &model.SuccessResponse{
		StatusCode: http.StatusOK,
		Message: "success create account",
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "500", http.StatusInternalServerError)
		return
	}
}

func (u *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "POST" {
		http.Error(w, "400", http.StatusBadRequest)
		return
	}

	var user *model.User
	d := json.NewDecoder(r.Body)
	if err := d.Decode(&user); err != nil {
		http.Error(w, "400", http.StatusBadRequest)
		return
	}

	deadline := time.Now().Add(3 * time.Minute)

	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	token, err := u.UserService.Find(ctx, user)
	if err != nil  {
		switch err {
        case u.UserService.ErrUserNotFound():
            http.Error(w, "User Not Found", http.StatusNotFound)
        case u.UserService.ErrInvalidPassword():
            http.Error(w, "Invalid Password", http.StatusUnauthorized)
        default:
            http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        }
		return	
	}

	response := &model.SuccessResponse{
		StatusCode: http.StatusOK,
		Message: token,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "500", http.StatusInternalServerError)
		return
	}

	// w.Write([]byte(fmt.Sprintf("access_token: %s", token)))
}