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

func (e *EventRepository) CreateEvent(user interface{}, event *model.Event) error {
	defer func() {
		if r := recover(); r != nil {
		  e.TX.Rollback()
		}
	  }()

	if err := e.TX.Create(&event).Error; err != nil {
		e.TX.Rollback()
		return err
	}

	return e.TX.Commit().Error
}