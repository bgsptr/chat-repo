package repository

import (
	"golang.org/x/calender-service/model"
	"gorm.io/gorm"
)

type EventRepository struct {
	TX *gorm.DB
}

func NewEventRepository(tx *gorm.DB) *EventRepository {
	return &EventRepository{
		TX: tx,
	}
}

func (e *EventRepository) CreateEvent(user interface{}, event *model.Event) (interface{}, error) {
	
	e.TX.G
}