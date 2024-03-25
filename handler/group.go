package handler

import (
	"encoding/json"
	"fmt"
	"groupservice/model"
	"log"
	"net/http"
	"strconv"
    "errors"
	"github.com/aws/aws-sdk-go/service"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
	"golang.org/x/text/cases"
)

type GroupChatHandler struct {
    Hub *model.Hub
    Room map[string]*model.Room
    GroupService *service.GroupService
}

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

func NewGroupChatHandler() *GroupChatHandler {
    return &GroupChatHandler{}
}

func NewHub() *model.Hub {
	return &model.Hub{
        Register:   make(chan *model.Client),
		Unregister: make(chan *model.Client),
        Broadcast:  make(chan []byte),
	}
}

func WebsocketGroupHandler(c echo.Context) error {
    ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
    if err != nil {
        log.Println(err)
        return err
    }
    defer ws.Close()


    // read user information from 
    user := c.Get("token")

    cl := &model.Client{
        name: user.Username,
    }

    go ReadMessageService()
    WriteMessageService()

    for {
        // Write
        err := ws.WriteMessage(websocket.TextMessage, )
        if err != nil {
            c.Logger().Error(err)
            return err
        }

        // Read
        _, msg, err := ws.ReadMessage()        
        if err != nil {
            c.Logger().Error(err)
            return err
        }
        fmt.Printf("%s\n", msg)
    }
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
    err := json.NewDecoder(c.Request().Body).Decode(room)
    if err != nil {
        return errors.New("error create group")
    }

    username := c.QueryParam("username")

    client := &model.Client{
        Conn: nil,
        Username: username,
        Password: "",
        Message: nil,
        RoomID: room.ID,
    }

    msg, err := h.GroupService.CreateGroup(client, room.ID)
    if err != nil {
        return err
    }

    h.Room = NewRoom(room.ID, room.Name)

    // message

	return c.JSON(http.StatusOK, u)
}

func JoinGroup(c echo.Context) error {

}

func (h *GroupChatHandler) Run() {
    for {
        select {
        case cl := <- h.Hub.Register:
            h.Room[cl.RoomID].Clients = append(h.Room[cl.RoomID].Clients, cl)
        case cl := <- h.Hub.Unregister:
            if _, ok := h.Room[cl.RoomID]; ok {
                if _, ok := h.Room[cl.RoomID].Clients[cl.Username]; ok {
                    
                    delete(h.Room[cl.RoomID].Clients, cl.Username)
                    close(cl.Message)
                }
            }
        case message := <- h.Hub.Broadcast:
            if _, ok := h.Room[cl.RoomID]; ok {
                for _, cl := range h.Room[cl.RoomID].Clients {
                    cl.Message <- message
                }
            }
        }
    }
}