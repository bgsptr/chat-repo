package service

import (
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

	newEvent := &model.Event{
		// ID: uuid.String(),
		EventName: event.EventName,
		FromDate: event.FromDate,
		ToDate: event.ToDate,
		Location: event.Location,
		Description: event.Description,
		PersonAdded: event.PersonAdded,
		PersonConfirmed: nil,
	}

	e.EventRepository.CreateEvent(user, newEvent)
	
	// event := &dto.EventDTO{
	// 	ID: 
	// }

	return event, nil
}