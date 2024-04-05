package service

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"golang.org/x/calender-service/internal/dto"
	"golang.org/x/calender-service/model"
)

type EventService struct {
	EventRepository *repository.EventRepository
	Validate *validator.Validate
}

func NewEventService(e *repository.EventRepository, v *validator.Validate) *EventService {
	return &EventService{
		EventRepository: e,
		Validate: v,
	}
}

func (e *EventService) CreateEventService(user interface{}, event *dto.EventDTO) (error) {

	user = &model.User{
		Username: user.Username,
		Pass: user.Pass,
	}

	e.Validate.TryValidate(user, event)

	event = &model.Event{
		ID: uuid.String(),
		FromDate: 
		ToDate:
		Location:
		Description:
		PersonAdded:
		PersonConfirmed: 
		Username: user.Username
	}



	e.EventRepository.CreateEvent(user)
	
	// event := &dto.EventDTO{
	// 	ID: 
	// }

	return event, nil
}