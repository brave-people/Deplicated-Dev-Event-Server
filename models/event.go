package models

import (
	"dev-event/requests"
	"dev-event/utils"
)

// gen:qs
type Event struct {
	ID          uint   `gorm:"primary_key;auto_increment"  json:"id"`
	Title       string `gorm:"not null"                 json:"title"`
	Description string `gorm:"size:255;not null"    json:"nickname"`
	ImgURL      string `json:"imgURL"`
	Contact     string `json:"contact"`
	Year        uint   `json:"year"`
	Month       uint   `json:"month"`
	Day         uint   `json:"day"`
	Time        uint   `json:"time"`
	Free        uint   `json:"free"`
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

func GetYYMMEvents(year string, month string) (events []Event, err error) {
	uintYear, err := utils.ConvertStringToUint(year)
	if err != nil {
		return
	}

	uintMonth, err := utils.ConvertStringToUint(month)
	if err != nil {
		return
	}

	qs := NewEventQuerySet(gGormDB)
	qs = qs.YearEq(uintYear)
	qs = qs.MonthEq(uintMonth)
	qs = qs.OrderAscByDay()
	err = qs.All(&events)

	return
}

func CreateEvent(reqEvent requests.Event) (event requests.Event, err error) {
	if err = gGormDB.Create(&reqEvent).Error; err != nil {
		return
	}

	return reqEvent, nil
}
