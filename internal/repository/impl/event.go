package impl

import "golang.org/x/calender-service/model"

type DBTX interface {
	CreateEvent() error
	UpdateEvent(id string) (*model.Event, error)
	FindGuestsInEvent(id string) ([]string, error)
	UpdateGuestByEventID(id string, guests []*model.EveryPerson) (*model.EventPersonConfirmed, error)
	FindEventByID(id string) (*model.Event, error)
	FindEventsByHost(username string) ([]*model.Event, error) 
	DeleteEvent(id string) error
}