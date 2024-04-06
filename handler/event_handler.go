package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/calender-service/internal/dto"
	"golang.org/x/calender-service/internal/service"
	"golang.org/x/calender-service/model"
)

var (
	BASE_URL = "http://localhost:8000/user"
)

var (
	internalServerError = 500
)

type EventHandler struct {
	*http.Client
	EventService *service.EventService
}

func NewEventHandler(e *service.EventService) *EventHandler {
	return &EventHandler{
		Client: new(http.Client),
		EventService: e,
	}
}

func (e *EventHandler) AddEvent(c *gin.Context) {

	if c.Request.Method != http.MethodPost {
		c.JSON(http.StatusBadRequest, gin.H{"status": "bad request"})
		return
	}

	// nembak ke user service
	req, err := http.NewRequest("GET", BASE_URL, nil)
	if err != nil {
		// http.Error(c.Writer, "400", http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, gin.H{"status": "bad request"})
		return
	}

	resp, err := e.Client.Do(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "bad request"})
		return
	}
	defer resp.Body.Close()

	var userDto *dto.Response

	err = json.NewDecoder(resp.Body).Decode(&userDto.Data)
	if err != nil {
		// http.Error(c.Writer, "404", http.StatusNotFound)
		// c.JSON(http.StatusNotFound, gin.H{"status": "Data not found"})
		// return
		responseUser := &dto.Response{
			StatusCode: 404,
			Status: "404 Not Found",
			Data: "Data not found",
		}
		c.JSON(responseUser.StatusCode, gin.H{"status": responseUser.Status})
		return
	}

	// get user request from addevent contoller

	var event *dto.EventDTO

	err = c.ShouldBindJSON(&event)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"status": "error internal server"})
		return
	}

	// decode to struct in json
	// if err = c.ShouldBindJSON(&dataCreated); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"status": "bad request"})
	// }

	responseUser := &dto.Response{
		StatusCode: 200,
		Status: "200 OK",
		Data: userDto.Data.(*model.User),
	}

	data, err := e.EventService.CreateEvent(responseUser, event)
	if err != nil {
		c.AbortWithStatusJSON(internalServerError, gin.H{"status": "internal service error"})
		return
	}

	if data == nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "error 404"})
	}


	c.JSON(http.StatusOK, gin.H{"status": "success 200"})

}

func (e *EventHandler) UpdateEvent(c *gin.Context) {

	if c.Request.Method != http.MethodPatch {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "400 Bad Request"})
	}

	var evt *dto.EventDTO
	err := json.NewDecoder(c.Request.Body).Decode(&evt)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "500 Internal Server Error"})
		return
	}

	id := c.Param("id")

	eventObj, err := e.EventService.UpdateEvent(id, evt)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "500 Internal Server Error"})
		return
	}

	c.IndentedJSON(http.StatusOK, dto)
}

func (e *EventHandler) DeleteEvent(c *gin.Context) {
	if c.Request.Method != http.MethodDelete {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "400 Bad Request"})
		return
	}

	id := c.Param("id")

	err := e.EventService.DeleteEvent(id) 
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "500 Internal Server Error"})
		return 
	}

	c.IndentedJSON(http.StatusOK, gin.H{"status": "event deleted"})
}

func (e *EventHandler) DetailEvent(c *gin.Context) {
	if c.Request.Method != http.MethodGet {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "400 Bad Request"})
		return
	}

	id := c.Param("id")

	evt, err := e.EventService.DetailEvent(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "500 Internal Server Error"})
	}

	res := &dto.Response{
		Status: "200 OK",
		StatusCode: 200,
		Data: evt,
	}

	c.IndentedJSON(http.StatusOK, res)
}

func (e *EventHandler) ShowAllEvents(c *gin.Context) {
	if c.Request.Method != http.MethodGet {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "400 Bad Request"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"status": "event deleted"})
}