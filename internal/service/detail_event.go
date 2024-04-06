package service

import "golang.org/x/calender-service/internal/dto"

func (e *EventService) DetailEvent(id string) (*dto.EventDTO, error) {
	evt, err := e.EventRepository.FindEventByID(id)
	if err != nil {
		return nil, err
	}

	guests, err := e.EventRepository.FindGuestsInEvent(id)
	if err != nil {
		return nil, err
	}

	var guest []string
	for _, username := range guests {
		guest = append(guest, username)
	}

	eventObject := &dto.EventDTO{
		EventName: evt.EventName,
		FromDate: evt.FromDate,
		ToDate: evt.ToDate,
		Location: evt.EventLocation,
		Description: evt.Descriptions,
		PersonConfirmed: guest,
	}

	return eventObject, nil
}