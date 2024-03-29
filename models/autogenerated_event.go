// Code generated by go-queryset. DO NOT EDIT.
package models

import (
	"errors"
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
)

// ===== BEGIN of all query sets

// ===== BEGIN of query set EventQuerySet

// EventQuerySet is an queryset type for Event
type EventQuerySet struct {
	db *gorm.DB
}

// NewEventQuerySet constructs new EventQuerySet
func NewEventQuerySet(db *gorm.DB) EventQuerySet {
	return EventQuerySet{
		db: db.Model(&Event{}),
	}
}

func (qs EventQuerySet) w(db *gorm.DB) EventQuerySet {
	return NewEventQuerySet(db)
}

func (qs EventQuerySet) Select(fields ...EventDBSchemaField) EventQuerySet {
	names := []string{}
	for _, f := range fields {
		names = append(names, f.String())
	}

	return qs.w(qs.db.Select(strings.Join(names, ",")))
}

// Create is an autogenerated method
// nolint: dupl
func (o *Event) Create(db *gorm.DB) error {
	return db.Create(o).Error
}

// Delete is an autogenerated method
// nolint: dupl
func (o *Event) Delete(db *gorm.DB) error {
	return db.Delete(o).Error
}

// All is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) All(ret *[]Event) error {
	return qs.db.Find(ret).Error
}

// ContactEq is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) ContactEq(contact string) EventQuerySet {
	return qs.w(qs.db.Where("contact = ?", contact))
}

// ContactGt is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) ContactGt(contact string) EventQuerySet {
	return qs.w(qs.db.Where("contact > ?", contact))
}

// ContactGte is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) ContactGte(contact string) EventQuerySet {
	return qs.w(qs.db.Where("contact >= ?", contact))
}

// ContactIn is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) ContactIn(contact ...string) EventQuerySet {
	if len(contact) == 0 {
		qs.db.AddError(errors.New("must at least pass one contact in ContactIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("contact IN (?)", contact))
}

// ContactLike is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) ContactLike(contact string) EventQuerySet {
	return qs.w(qs.db.Where("contact LIKE ?", contact))
}

// ContactLt is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) ContactLt(contact string) EventQuerySet {
	return qs.w(qs.db.Where("contact < ?", contact))
}

// ContactLte is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) ContactLte(contact string) EventQuerySet {
	return qs.w(qs.db.Where("contact <= ?", contact))
}

// ContactNe is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) ContactNe(contact string) EventQuerySet {
	return qs.w(qs.db.Where("contact != ?", contact))
}

// ContactNotIn is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) ContactNotIn(contact ...string) EventQuerySet {
	if len(contact) == 0 {
		qs.db.AddError(errors.New("must at least pass one contact in ContactNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("contact NOT IN (?)", contact))
}

// ContactNotlike is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) ContactNotlike(contact string) EventQuerySet {
	return qs.w(qs.db.Where("contact NOT LIKE ?", contact))
}

// Count is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) Count() (int, error) {
	var count int
	err := qs.db.Count(&count).Error
	return count, err
}

// DayEq is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) DayEq(day uint) EventQuerySet {
	return qs.w(qs.db.Where("day = ?", day))
}

// DayGt is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) DayGt(day uint) EventQuerySet {
	return qs.w(qs.db.Where("day > ?", day))
}

// DayGte is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) DayGte(day uint) EventQuerySet {
	return qs.w(qs.db.Where("day >= ?", day))
}

// DayIn is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) DayIn(day ...uint) EventQuerySet {
	if len(day) == 0 {
		qs.db.AddError(errors.New("must at least pass one day in DayIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("day IN (?)", day))
}

// DayLt is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) DayLt(day uint) EventQuerySet {
	return qs.w(qs.db.Where("day < ?", day))
}

// DayLte is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) DayLte(day uint) EventQuerySet {
	return qs.w(qs.db.Where("day <= ?", day))
}

// DayNe is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) DayNe(day uint) EventQuerySet {
	return qs.w(qs.db.Where("day != ?", day))
}

// DayNotIn is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) DayNotIn(day ...uint) EventQuerySet {
	if len(day) == 0 {
		qs.db.AddError(errors.New("must at least pass one day in DayNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("day NOT IN (?)", day))
}

// Delete is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) Delete() error {
	return qs.db.Delete(Event{}).Error
}

// DeleteNum is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) DeleteNum() (int64, error) {
	db := qs.db.Delete(Event{})
	return db.RowsAffected, db.Error
}

// DeleteNumUnscoped is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) DeleteNumUnscoped() (int64, error) {
	db := qs.db.Unscoped().Delete(Event{})
	return db.RowsAffected, db.Error
}

// DescriptionEq is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) DescriptionEq(description string) EventQuerySet {
	return qs.w(qs.db.Where("description = ?", description))
}

// DescriptionGt is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) DescriptionGt(description string) EventQuerySet {
	return qs.w(qs.db.Where("description > ?", description))
}

// DescriptionGte is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) DescriptionGte(description string) EventQuerySet {
	return qs.w(qs.db.Where("description >= ?", description))
}

// DescriptionIn is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) DescriptionIn(description ...string) EventQuerySet {
	if len(description) == 0 {
		qs.db.AddError(errors.New("must at least pass one description in DescriptionIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("description IN (?)", description))
}

// DescriptionLike is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) DescriptionLike(description string) EventQuerySet {
	return qs.w(qs.db.Where("description LIKE ?", description))
}

// DescriptionLt is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) DescriptionLt(description string) EventQuerySet {
	return qs.w(qs.db.Where("description < ?", description))
}

// DescriptionLte is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) DescriptionLte(description string) EventQuerySet {
	return qs.w(qs.db.Where("description <= ?", description))
}

// DescriptionNe is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) DescriptionNe(description string) EventQuerySet {
	return qs.w(qs.db.Where("description != ?", description))
}

// DescriptionNotIn is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) DescriptionNotIn(description ...string) EventQuerySet {
	if len(description) == 0 {
		qs.db.AddError(errors.New("must at least pass one description in DescriptionNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("description NOT IN (?)", description))
}

// DescriptionNotlike is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) DescriptionNotlike(description string) EventQuerySet {
	return qs.w(qs.db.Where("description NOT LIKE ?", description))
}

// FreeEq is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) FreeEq(free uint) EventQuerySet {
	return qs.w(qs.db.Where("free = ?", free))
}

// FreeGt is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) FreeGt(free uint) EventQuerySet {
	return qs.w(qs.db.Where("free > ?", free))
}

// FreeGte is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) FreeGte(free uint) EventQuerySet {
	return qs.w(qs.db.Where("free >= ?", free))
}

// FreeIn is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) FreeIn(free ...uint) EventQuerySet {
	if len(free) == 0 {
		qs.db.AddError(errors.New("must at least pass one free in FreeIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("free IN (?)", free))
}

// FreeLt is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) FreeLt(free uint) EventQuerySet {
	return qs.w(qs.db.Where("free < ?", free))
}

// FreeLte is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) FreeLte(free uint) EventQuerySet {
	return qs.w(qs.db.Where("free <= ?", free))
}

// FreeNe is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) FreeNe(free uint) EventQuerySet {
	return qs.w(qs.db.Where("free != ?", free))
}

// FreeNotIn is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) FreeNotIn(free ...uint) EventQuerySet {
	if len(free) == 0 {
		qs.db.AddError(errors.New("must at least pass one free in FreeNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("free NOT IN (?)", free))
}

// GetDB is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) GetDB() *gorm.DB {
	return qs.db
}

// GetUpdater is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) GetUpdater() EventUpdater {
	return NewEventUpdater(qs.db)
}

// IDEq is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) IDEq(ID uint) EventQuerySet {
	return qs.w(qs.db.Where("id = ?", ID))
}

// IDGt is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) IDGt(ID uint) EventQuerySet {
	return qs.w(qs.db.Where("id > ?", ID))
}

// IDGte is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) IDGte(ID uint) EventQuerySet {
	return qs.w(qs.db.Where("id >= ?", ID))
}

// IDIn is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) IDIn(ID ...uint) EventQuerySet {
	if len(ID) == 0 {
		qs.db.AddError(errors.New("must at least pass one ID in IDIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("id IN (?)", ID))
}

// IDLt is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) IDLt(ID uint) EventQuerySet {
	return qs.w(qs.db.Where("id < ?", ID))
}

// IDLte is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) IDLte(ID uint) EventQuerySet {
	return qs.w(qs.db.Where("id <= ?", ID))
}

// IDNe is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) IDNe(ID uint) EventQuerySet {
	return qs.w(qs.db.Where("id != ?", ID))
}

// IDNotIn is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) IDNotIn(ID ...uint) EventQuerySet {
	if len(ID) == 0 {
		qs.db.AddError(errors.New("must at least pass one ID in IDNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("id NOT IN (?)", ID))
}

// ImgURLEq is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) ImgURLEq(imgURL string) EventQuerySet {
	return qs.w(qs.db.Where("img_url = ?", imgURL))
}

// ImgURLGt is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) ImgURLGt(imgURL string) EventQuerySet {
	return qs.w(qs.db.Where("img_url > ?", imgURL))
}

// ImgURLGte is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) ImgURLGte(imgURL string) EventQuerySet {
	return qs.w(qs.db.Where("img_url >= ?", imgURL))
}

// ImgURLIn is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) ImgURLIn(imgURL ...string) EventQuerySet {
	if len(imgURL) == 0 {
		qs.db.AddError(errors.New("must at least pass one imgURL in ImgURLIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("img_url IN (?)", imgURL))
}

// ImgURLLike is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) ImgURLLike(imgURL string) EventQuerySet {
	return qs.w(qs.db.Where("img_url LIKE ?", imgURL))
}

// ImgURLLt is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) ImgURLLt(imgURL string) EventQuerySet {
	return qs.w(qs.db.Where("img_url < ?", imgURL))
}

// ImgURLLte is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) ImgURLLte(imgURL string) EventQuerySet {
	return qs.w(qs.db.Where("img_url <= ?", imgURL))
}

// ImgURLNe is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) ImgURLNe(imgURL string) EventQuerySet {
	return qs.w(qs.db.Where("img_url != ?", imgURL))
}

// ImgURLNotIn is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) ImgURLNotIn(imgURL ...string) EventQuerySet {
	if len(imgURL) == 0 {
		qs.db.AddError(errors.New("must at least pass one imgURL in ImgURLNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("img_url NOT IN (?)", imgURL))
}

// ImgURLNotlike is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) ImgURLNotlike(imgURL string) EventQuerySet {
	return qs.w(qs.db.Where("img_url NOT LIKE ?", imgURL))
}

// Limit is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) Limit(limit int) EventQuerySet {
	return qs.w(qs.db.Limit(limit))
}

// MonthEq is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) MonthEq(month uint) EventQuerySet {
	return qs.w(qs.db.Where("month = ?", month))
}

// MonthGt is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) MonthGt(month uint) EventQuerySet {
	return qs.w(qs.db.Where("month > ?", month))
}

// MonthGte is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) MonthGte(month uint) EventQuerySet {
	return qs.w(qs.db.Where("month >= ?", month))
}

// MonthIn is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) MonthIn(month ...uint) EventQuerySet {
	if len(month) == 0 {
		qs.db.AddError(errors.New("must at least pass one month in MonthIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("month IN (?)", month))
}

// MonthLt is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) MonthLt(month uint) EventQuerySet {
	return qs.w(qs.db.Where("month < ?", month))
}

// MonthLte is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) MonthLte(month uint) EventQuerySet {
	return qs.w(qs.db.Where("month <= ?", month))
}

// MonthNe is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) MonthNe(month uint) EventQuerySet {
	return qs.w(qs.db.Where("month != ?", month))
}

// MonthNotIn is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) MonthNotIn(month ...uint) EventQuerySet {
	if len(month) == 0 {
		qs.db.AddError(errors.New("must at least pass one month in MonthNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("month NOT IN (?)", month))
}

// Offset is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) Offset(offset int) EventQuerySet {
	return qs.w(qs.db.Offset(offset))
}

// One is used to retrieve one result. It returns gorm.ErrRecordNotFound
// if nothing was fetched
func (qs EventQuerySet) One(ret *Event) error {
	return qs.db.First(ret).Error
}

// OrderAscByContact is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) OrderAscByContact() EventQuerySet {
	return qs.w(qs.db.Order("contact ASC"))
}

// OrderAscByDay is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) OrderAscByDay() EventQuerySet {
	return qs.w(qs.db.Order("day ASC"))
}

// OrderAscByDescription is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) OrderAscByDescription() EventQuerySet {
	return qs.w(qs.db.Order("description ASC"))
}

// OrderAscByFree is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) OrderAscByFree() EventQuerySet {
	return qs.w(qs.db.Order("free ASC"))
}

// OrderAscByID is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) OrderAscByID() EventQuerySet {
	return qs.w(qs.db.Order("id ASC"))
}

// OrderAscByImgURL is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) OrderAscByImgURL() EventQuerySet {
	return qs.w(qs.db.Order("img_url ASC"))
}

// OrderAscByMonth is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) OrderAscByMonth() EventQuerySet {
	return qs.w(qs.db.Order("month ASC"))
}

// OrderAscByTime is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) OrderAscByTime() EventQuerySet {
	return qs.w(qs.db.Order("time ASC"))
}

// OrderAscByTitle is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) OrderAscByTitle() EventQuerySet {
	return qs.w(qs.db.Order("title ASC"))
}

// OrderAscByYear is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) OrderAscByYear() EventQuerySet {
	return qs.w(qs.db.Order("year ASC"))
}

// OrderDescByContact is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) OrderDescByContact() EventQuerySet {
	return qs.w(qs.db.Order("contact DESC"))
}

// OrderDescByDay is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) OrderDescByDay() EventQuerySet {
	return qs.w(qs.db.Order("day DESC"))
}

// OrderDescByDescription is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) OrderDescByDescription() EventQuerySet {
	return qs.w(qs.db.Order("description DESC"))
}

// OrderDescByFree is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) OrderDescByFree() EventQuerySet {
	return qs.w(qs.db.Order("free DESC"))
}

// OrderDescByID is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) OrderDescByID() EventQuerySet {
	return qs.w(qs.db.Order("id DESC"))
}

// OrderDescByImgURL is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) OrderDescByImgURL() EventQuerySet {
	return qs.w(qs.db.Order("img_url DESC"))
}

// OrderDescByMonth is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) OrderDescByMonth() EventQuerySet {
	return qs.w(qs.db.Order("month DESC"))
}

// OrderDescByTime is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) OrderDescByTime() EventQuerySet {
	return qs.w(qs.db.Order("time DESC"))
}

// OrderDescByTitle is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) OrderDescByTitle() EventQuerySet {
	return qs.w(qs.db.Order("title DESC"))
}

// OrderDescByYear is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) OrderDescByYear() EventQuerySet {
	return qs.w(qs.db.Order("year DESC"))
}

// TimeEq is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) TimeEq(time uint) EventQuerySet {
	return qs.w(qs.db.Where("time = ?", time))
}

// TimeGt is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) TimeGt(time uint) EventQuerySet {
	return qs.w(qs.db.Where("time > ?", time))
}

// TimeGte is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) TimeGte(time uint) EventQuerySet {
	return qs.w(qs.db.Where("time >= ?", time))
}

// TimeIn is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) TimeIn(time ...uint) EventQuerySet {
	if len(time) == 0 {
		qs.db.AddError(errors.New("must at least pass one time in TimeIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("time IN (?)", time))
}

// TimeLt is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) TimeLt(time uint) EventQuerySet {
	return qs.w(qs.db.Where("time < ?", time))
}

// TimeLte is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) TimeLte(time uint) EventQuerySet {
	return qs.w(qs.db.Where("time <= ?", time))
}

// TimeNe is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) TimeNe(time uint) EventQuerySet {
	return qs.w(qs.db.Where("time != ?", time))
}

// TimeNotIn is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) TimeNotIn(time ...uint) EventQuerySet {
	if len(time) == 0 {
		qs.db.AddError(errors.New("must at least pass one time in TimeNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("time NOT IN (?)", time))
}

// TitleEq is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) TitleEq(title string) EventQuerySet {
	return qs.w(qs.db.Where("title = ?", title))
}

// TitleGt is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) TitleGt(title string) EventQuerySet {
	return qs.w(qs.db.Where("title > ?", title))
}

// TitleGte is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) TitleGte(title string) EventQuerySet {
	return qs.w(qs.db.Where("title >= ?", title))
}

// TitleIn is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) TitleIn(title ...string) EventQuerySet {
	if len(title) == 0 {
		qs.db.AddError(errors.New("must at least pass one title in TitleIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("title IN (?)", title))
}

// TitleLike is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) TitleLike(title string) EventQuerySet {
	return qs.w(qs.db.Where("title LIKE ?", title))
}

// TitleLt is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) TitleLt(title string) EventQuerySet {
	return qs.w(qs.db.Where("title < ?", title))
}

// TitleLte is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) TitleLte(title string) EventQuerySet {
	return qs.w(qs.db.Where("title <= ?", title))
}

// TitleNe is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) TitleNe(title string) EventQuerySet {
	return qs.w(qs.db.Where("title != ?", title))
}

// TitleNotIn is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) TitleNotIn(title ...string) EventQuerySet {
	if len(title) == 0 {
		qs.db.AddError(errors.New("must at least pass one title in TitleNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("title NOT IN (?)", title))
}

// TitleNotlike is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) TitleNotlike(title string) EventQuerySet {
	return qs.w(qs.db.Where("title NOT LIKE ?", title))
}

// YearEq is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) YearEq(year uint) EventQuerySet {
	return qs.w(qs.db.Where("year = ?", year))
}

// YearGt is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) YearGt(year uint) EventQuerySet {
	return qs.w(qs.db.Where("year > ?", year))
}

// YearGte is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) YearGte(year uint) EventQuerySet {
	return qs.w(qs.db.Where("year >= ?", year))
}

// YearIn is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) YearIn(year ...uint) EventQuerySet {
	if len(year) == 0 {
		qs.db.AddError(errors.New("must at least pass one year in YearIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("year IN (?)", year))
}

// YearLt is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) YearLt(year uint) EventQuerySet {
	return qs.w(qs.db.Where("year < ?", year))
}

// YearLte is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) YearLte(year uint) EventQuerySet {
	return qs.w(qs.db.Where("year <= ?", year))
}

// YearNe is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) YearNe(year uint) EventQuerySet {
	return qs.w(qs.db.Where("year != ?", year))
}

// YearNotIn is an autogenerated method
// nolint: dupl
func (qs EventQuerySet) YearNotIn(year ...uint) EventQuerySet {
	if len(year) == 0 {
		qs.db.AddError(errors.New("must at least pass one year in YearNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("year NOT IN (?)", year))
}

// SetContact is an autogenerated method
// nolint: dupl
func (u EventUpdater) SetContact(contact string) EventUpdater {
	u.fields[string(EventDBSchema.Contact)] = contact
	return u
}

// SetDay is an autogenerated method
// nolint: dupl
func (u EventUpdater) SetDay(day uint) EventUpdater {
	u.fields[string(EventDBSchema.Day)] = day
	return u
}

// SetDescription is an autogenerated method
// nolint: dupl
func (u EventUpdater) SetDescription(description string) EventUpdater {
	u.fields[string(EventDBSchema.Description)] = description
	return u
}

// SetFree is an autogenerated method
// nolint: dupl
func (u EventUpdater) SetFree(free uint) EventUpdater {
	u.fields[string(EventDBSchema.Free)] = free
	return u
}

// SetID is an autogenerated method
// nolint: dupl
func (u EventUpdater) SetID(ID uint) EventUpdater {
	u.fields[string(EventDBSchema.ID)] = ID
	return u
}

// SetImgURL is an autogenerated method
// nolint: dupl
func (u EventUpdater) SetImgURL(imgURL string) EventUpdater {
	u.fields[string(EventDBSchema.ImgURL)] = imgURL
	return u
}

// SetMonth is an autogenerated method
// nolint: dupl
func (u EventUpdater) SetMonth(month uint) EventUpdater {
	u.fields[string(EventDBSchema.Month)] = month
	return u
}

// SetTime is an autogenerated method
// nolint: dupl
func (u EventUpdater) SetTime(time uint) EventUpdater {
	u.fields[string(EventDBSchema.Time)] = time
	return u
}

// SetTitle is an autogenerated method
// nolint: dupl
func (u EventUpdater) SetTitle(title string) EventUpdater {
	u.fields[string(EventDBSchema.Title)] = title
	return u
}

// SetYear is an autogenerated method
// nolint: dupl
func (u EventUpdater) SetYear(year uint) EventUpdater {
	u.fields[string(EventDBSchema.Year)] = year
	return u
}

// Update is an autogenerated method
// nolint: dupl
func (u EventUpdater) Update() error {
	return u.db.Updates(u.fields).Error
}

// UpdateNum is an autogenerated method
// nolint: dupl
func (u EventUpdater) UpdateNum() (int64, error) {
	db := u.db.Updates(u.fields)
	return db.RowsAffected, db.Error
}

// ===== END of query set EventQuerySet

// ===== BEGIN of Event modifiers

// EventDBSchemaField describes database schema field. It requires for method 'Update'
type EventDBSchemaField string

// String method returns string representation of field.
// nolint: dupl
func (f EventDBSchemaField) String() string {
	return string(f)
}

// EventDBSchema stores db field names of Event
var EventDBSchema = struct {
	ID          EventDBSchemaField
	Title       EventDBSchemaField
	Description EventDBSchemaField
	ImgURL      EventDBSchemaField
	Contact     EventDBSchemaField
	Year        EventDBSchemaField
	Month       EventDBSchemaField
	Day         EventDBSchemaField
	Time        EventDBSchemaField
	Free        EventDBSchemaField
}{

	ID:          EventDBSchemaField("id"),
	Title:       EventDBSchemaField("title"),
	Description: EventDBSchemaField("description"),
	ImgURL:      EventDBSchemaField("img_url"),
	Contact:     EventDBSchemaField("contact"),
	Year:        EventDBSchemaField("year"),
	Month:       EventDBSchemaField("month"),
	Day:         EventDBSchemaField("day"),
	Time:        EventDBSchemaField("time"),
	Free:        EventDBSchemaField("free"),
}

// Update updates Event fields by primary key
// nolint: dupl
func (o *Event) Update(db *gorm.DB, fields ...EventDBSchemaField) error {
	dbNameToFieldName := map[string]interface{}{
		"id":          o.ID,
		"title":       o.Title,
		"description": o.Description,
		"img_url":     o.ImgURL,
		"contact":     o.Contact,
		"year":        o.Year,
		"month":       o.Month,
		"day":         o.Day,
		"time":        o.Time,
		"free":        o.Free,
	}
	u := map[string]interface{}{}
	for _, f := range fields {
		fs := f.String()
		u[fs] = dbNameToFieldName[fs]
	}
	if err := db.Model(o).Updates(u).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return err
		}

		return fmt.Errorf("can't update Event %v fields %v: %s",
			o, fields, err)
	}

	return nil
}

// EventUpdater is an Event updates manager
type EventUpdater struct {
	fields map[string]interface{}
	db     *gorm.DB
}

// NewEventUpdater creates new Event updater
// nolint: dupl
func NewEventUpdater(db *gorm.DB) EventUpdater {
	return EventUpdater{
		fields: map[string]interface{}{},
		db:     db.Model(&Event{}),
	}
}

// ===== END of Event modifiers

// ===== END of all query sets
