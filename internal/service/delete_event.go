package service

func (e *EventService) DeleteEvent(id string) error {
	err := e.EventRepository.Delete(id)
	if err != nil {
		return err
	}
	return err
}