package models

import (
	"dev-event/requests"
	"dev-event/utils"
	"errors"
)

func GetAllUsers() (users []Users, err error) {
	qs := NewUsersQuerySet(gGormDB)
	err = qs.All(&users)
	return
}

func GetUser(userID string) (user Users, err error) {
	uintUserID, err := utils.ConvertStringToUint(userID)
	if err != nil {
		err = errors.New("Invalid ID")
		return
	}

	qs := NewUsersQuerySet(gGormDB)
	qs = qs.IDEq(uintUserID)
	err = qs.One(&user)
	return
}

func AdminUpdateUser(user requests.Users, userID string) (err error) {
	uintUserID, err := utils.ConvertStringToUint(userID)
	if err != nil {
		return
	}

	EmailValidation(user.Email)
	if err != nil {
		return
	}
	NicknameValication(user.Nickname)
	if err != nil {
		return
	}

	encPassword, err := EncPassword(user.Password)
	if err != nil {
		return
	}

	_, err = NewUsersQuerySet(gGormDB).
		IDEq(uintUserID).
		GetUpdater().
		SetEmail(user.Email).
		SetNickname(user.Nickname).
		SetLevel(user.Level).
		SetPassword(encPassword).
		UpdateNum()

	return
}

func AdminDeleteUser(userID string) (err error) {
	uintUserID, err := utils.ConvertStringToUint(userID)
	if err != nil {
		return errors.New("Unvalid userID")
	}

	user := &Users{
		ID: uintUserID,
	}
	return user.Delete(gGormDB)
}

func AdminUserChecker(userID uint) (err error) {
	var user Users
	qs := NewUsersQuerySet(gGormDB)
	qs = qs.IDEq(userID)
	err = qs.One(&user)
	if user.Level != 0 {
		return errors.New("No admin auth")
	}

	return
}
