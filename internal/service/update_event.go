package service

import (
	"errors"

	"golang.org/x/calender-service/internal/dto"
	"golang.org/x/calender-service/model"
)

func (e *EventService) UpdateEvent(id string, evt *dto.EventDTO) (*dto.EventDTO, error) {
	evtModel := &model.Event{
		EventName: evt.EventName,
		FromDate: evt.FromDate,
		ToDate: evt.ToDate,
		EventLocation: evt.Location,
		Descriptions: evt.Description,
	}

	err := e.Validate.TryValidate(evtModel)
	if err != nil {
		return nil, err
	}

	event, err := e.EventRepository.Update(id, evtModel)
	if err != nil {
		return nil, err
	}

	dtoUpdatedEvent := &dto.EventDTO{
		Id: ,
	}

	personsEdited := evt.PersonAdded

	if len(AllGuest) == 0 {
		return dtoUpdatedEvent, nil
	}


	guests, err := e.EventRepository.FindGuestsInEvent(id)
	if err != nil {
		return nil, errors.New("failed")
	}

	for _, personEdited := range personsEdited {
		if guest.Username not in personEdited {

		}
	}
	return dtoUpdatedEvent, nil
}