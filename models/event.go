package models

import (
	"dev-event/requests"
	"dev-event/utils"
)

// gen:qs
type Event struct {
	ID          uint   `gorm:"primary_key;auto_increment"  json:"id"`
	Description string `gorm:"size:255;not null;unique"    json:"nickname"`
	Tages       []Tag  `gorm:"many2many:event_tags;" json:"tages"`
}

func GetOneEvent(eventID string) (event Event, err error) {
	if res := gGormDB.Where("id = ?", eventID).Find(&event); res.Error != nil {
		return
	}
	gGormDB.Model(&event).Association("Tages").Find(&event.Tages)

	return
}

func SaveEvent(event requests.Event, eventID string) (err error) {
	uintEventID, err := utils.ConvertStringToUint(eventID)
	if err != nil {
		return err
	}

	_, err = NewEventQuerySet(gGormDB).
		IDEq(uintEventID).
		GetUpdater().
		SetDescription(event.Description).
		UpdateNum()

	return
}

func DeleteEvent(eventID string) (err error) {
	uintEventID, err := utils.ConvertStringToUint(eventID)
	if err != nil {
		return
	}

	event := &Event{
		ID: uintEventID,
	}

	return event.Delete(gGormDB)
}
