package service

import (
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"golang.org/x/calender-service/config"
	"golang.org/x/calender-service/internal/dto"
	"golang.org/x/calender-service/internal/repository"
	"golang.org/x/calender-service/model"
)

type EventService struct {
	EventRepository *repository.EventRepository
	Validate *config.CustomValidator
}

func NewEventService(e *repository.EventRepository, v *config.CustomValidator) *EventService {
	return &EventService{
		EventRepository: e,
		Validate: v,
	}
}

func (e *EventService) CreateEvent(user interface{}, event *dto.EventDTO) (interface{}, error) {

	user = &model.User{
		Username: user.Username,
	}

	err := e.Validate.TryValidate(user, event)
	if err != nil {
		return nil, err
	}

	var id string 
	// generateID()

	newEvent := &model.Event{
		// Id: id,
		EventName: event.EventName,
		FromDate: event.FromDate,
		ToDate: event.ToDate,
		EventLocation: event.Location,
		Descriptions: event.Description,
	}

	person := make(map[string][]string)

	var everyPerson []string
	for _, val := range event.PersonConfirmed {
		everyPerson = append(everyPerson, val.Username)
	}
	person[id] = everyPerson

	err = e.EventRepository.CreateEvent(user, newEvent, person)
	if err != nil {
		return nil, err
	}
	
	// event := &dto.EventDTO{
	// 	ID: 
	// }

	return event, nil
}