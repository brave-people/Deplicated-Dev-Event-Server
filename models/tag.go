package models

import (
	"dev-event/requests"
	"dev-event/utils"
)

// gen:qs
type Tag struct {
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `json:"name"`
}

func GetAllTags() (tags []Tag, err error) {
	qs := NewTagQuerySet(gGormDB)
	err = qs.All(&tags)
	return
}

func DeleteTag(tagID string) (err error) {
	uintTagID, err := utils.ConvertStringToUint(tagID)
	if err != nil {
		return
	}

	tag := &Tag{
		ID: uintTagID,
	}

	return tag.Delete(gGormDB)
}

func UpdateTag(req requests.Tag, tagID string) (err error) {
	uintTagID, err := utils.ConvertStringToUint(tagID)
	if err != nil {
		return
	}

	_, err = NewTagQuerySet(gGormDB).
		IDEq(uintTagID).
		GetUpdater().
		SetName(req.Name).
		UpdateNum()

	return
}

func CreateTag(req requests.Tag) (id uint, err error) {
	tag := &Tag{
		Name: req.Name,
	}

	err = tag.Create(gGormDB)
	if err != nil {
		return
	}

	return tag.ID, err
}
