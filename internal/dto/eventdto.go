package dto

import "golang.org/x/calender-service/model"

type EventDTO struct {
	// ID              string        `json:"id_event"`
	EventName		string		  `json:"event_name"`
	FromDate        string        `json:"from_date"`
	ToDate          string        `json:"to_date"`
	Location        string        `json:"location"`
	Description     string        `json:"description"`
	PersonAdded     []*model.User `json:"person_wish_list"`
	PersonConfirmed []*model.User `json:"person_confirmed"`
	// Username        string        `json:"username"`
}