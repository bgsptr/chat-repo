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

func (e *EventRepository) Update(id string, evtEntity *model.Event) (*model.Event, error) {
	defer func() {
		if r := recover(); r != nil {
		  e.TX.Rollback()
		}
	  }()


}

func (e *EventRepository) FindGuestsInEvent(idEvt string) ([]string, error) {
	var guests []string
	err := e.TX.Model(&model.EventPersonConfirmed{}).Select("username").Where("id = ?", idEvt).Find(&guests).Error
	if err != nil {
		return nil, errDatabase
	}

	return guests, nil
}

func (e *EventRepository) Delete(id string) error {
	defer func() {
		if r := recover(); r != nil {
		  e.TX.Rollback()
		}
	  }()

	err := e.TX.Where("id = ?", id).Delete(&model.Event{}).Error
	if err != nil {
		e.TX.Rollback()
		return err
	}

	err = e.TX.Where("id = ?", id).Delete(&model.EventPersonConfirmed{}).Error
	if err != nil {
		e.TX.Rollback()
		return err
	}
	return e.TX.Commit().Error
}

func (e *EventRepository) FindEventByID(id string) (*model.Event, error) {
	defer func() {
		if r := recover(); r != nil {
		  e.TX.Rollback()
		}
	  }()

	var model *model.Event
	
	if err := e.TX.Where("id = ?").Find(&model).Error; err != nil {
		e.TX.Rollback()
		return nil, err
	}

	return model, nil
}

func (e *EventRepository) FindEventsByHost(username string) ([]*model.Event, error) {
	var event []*model.Event
	err := e.TX.First(&event, "username = ?", username).Order("from_date DESC")
	if err != nil {
		return nil, errDatabase
	}

	return event, errDatabase
}

func (e *EventRepository) UpdateGuestByEventID(id string, guests []*model.EveryPerson) (*model.EventPersonConfirmed, error) {
	
}