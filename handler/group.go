package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"groupservice/model"
	"groupservice/service"
	"log"
	"net/http"
	"strconv"
	"time"

	// "github.com/aws/aws-sdk-go/service"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
	"golang.org/x/text/cases"
)

type GroupChatHandler struct {
    Hub *model.Hub
    // Room map[string]*model.Room
    GroupService *service.GroupService
}

type Manager struct {
    *model.Hub
}

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

func NewGroupChatHandler(srv *service.GroupService, manager *model.Hub) *GroupChatHandler {
    return &GroupChatHandler{
        Hub: manager,
        GroupService: srv,
    }
}

func NewHub() *model.Hub {
	return &model.Hub{
        Room: make(map[string]*model.Chat),
        Register:   make(chan *model.Client),
		Unregister: make(chan *model.Client),
        Broadcast:  make(chan *model.Message),
	}
}

// join chat id

func (h *GroupChatHandler) WebsocketGroupHandler(c echo.Context) error {
    ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
    if err != nil {
        log.Println(err)
        return err
    }
    defer ws.Close()


    // read user information from token
    user := c.Get("token")

    values := c.Request().URL.Query()
    roomID := values.Get("roomID")

    newCl := &model.Client{
        Conn: ws,
        Username: user.Username,
        Password: "",
        Message: make(chan *model.Message),
        RoomID: roomID,
    }

    var messageSend *model.Message

    errDecode := json.NewDecoder(c.Request().Body).Decode(&messageSend)
    if errDecode != nil {
        return errDecode
    }

    msg := &model.Message{
        Content: "New User Has Joined",
        CreatedAt: time.Now().Local().String(),
        RoomID: roomID,
    }
    
    cl := service.NewClient(newCl)

    find, err := h.GroupService.JoinGroup(cl.Client.Username, cl.RoomID)
    if !find {
        h.Hub.Broadcast <- msg
    }

    if err != nil {
        c.Logger().Error(err)
    }


    h.Hub.Register <- cl.Client
    h.Hub.Broadcast <- messageSend
    

    go cl.WriteMessageService()
    cl.ReadMessageService(h.Hub)
}


func NewRoom(id string, name string) *model.Room {
    return &model.Room{
        ID: id,
        Name: name,
        Clients: make(map[string]*model.Client),
    }
}

func (h *GroupChatHandler) CreateGroup(c echo.Context) error {
    var room *model.Room
    err := json.NewDecoder(c.Request().Body).Decode(&room)
    if err != nil {
        return errors.New("error create group")
    }

    user := c.Get("token")

    client := &model.Client{
        Conn: nil,
        Username: user.Username,
        Password: "",
        Message: nil,
        RoomID: room.ID,
    }

    msg, err := h.GroupService.CreateGroup(client, room.ID)
    if err != nil {
        return err
    }

    // for _, l := range h.Room {
    //     h.Room[l] = NewRoom(room.ID, room.Name)
    // }

    // message

	return c.JSON(http.StatusOK, u)
}

func (h *GroupChatHandler) Run() {
    for {
        select {
        case cl := <- h.Hub.Register:
            if room, ok := h.Hub.Room[cl.RoomID]; ok {
                room.To[cl.Username] = cl
            }
        case cl := <- h.Hub.Unregister:
            if room, ok := h.Hub.Room[cl.RoomID]; ok {
                if _, ok := room.To[cl.Username]; ok {
                    
                    delete(h.Hub.Room[cl.RoomID].To, cl.Username)
                    close(cl.Message)
                }
            }
        case message := <- h.Hub.Broadcast:
            if _, ok := h.Hub.Room[message.RoomID]; ok {
                for _, cl := range h.Hub.Room[message.RoomID].To {
                    cl.Message <- message
                }
            }
        }
    }
}