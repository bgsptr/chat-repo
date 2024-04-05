package model

type Event struct {
	EventName       string  `json:"event_name"`
	FromDate        string  `json:"from_date"`
	ToDate          string  `json:"to_date"`
	Location        string  `json:"location"`
	Description     string  `json:"description"`
	PersonAdded     []*User `json:"person_wish_list"`
	PersonConfirmed []*User `json:"person_confirmed"`
}