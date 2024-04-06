package repository

import (
	"errors"

	"golang.org/x/calender-service/model"
	"gorm.io/gorm"
)

var (
	errDatabase = errors.New("Something wrong")
)

type EventRepository struct {
	TX *gorm.DB
}

func NewEventRepository(tx *gorm.DB) *EventRepository {
	return &EventRepository{
		TX: tx,
	}
}

func (e *EventRepository) CreateEvent(user interface{}, event *model.Event, person map[string][]string) error {
	defer func() {
		if r := recover(); r != nil {
		  e.TX.Rollback()
		}
	  }()

	if err := e.TX.Create(&event).Error; err != nil {
		e.TX.Rollback()
		return errDatabase
	}

	var allPerson []*model.EventPersonConfirmed

	for _, username := range person[event.Id] {
		personConfirmed := &model.EventPersonConfirmed{
			Id: event.Id,
			Username: username,
			IsConfirmed: false,
		}
		allPerson = append(allPerson, personConfirmed)
	}
	
	err := e.TX.CreateInBatches(allPerson, len(person[event.Id]))
	if err != nil {
		e.TX.Rollback()
		return errDatabase
	}

	return e.TX.Commit().Error
}