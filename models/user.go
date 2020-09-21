package models

import (
	"dev-event/requests"
	"errors"
	"log"
	"time"

	"github.com/badoux/checkmail"
	"golang.org/x/crypto/bcrypt"
)

// gen:qs
type Users struct {
	ID        uint      `gorm:"primary_key;auto_increment" json:"id"`
	Nickname  string    `gorm:"size:255;not null;unique"   json:"nickname"`
	Email     string    `gorm:"size:100;not null;unique"   json:"email"`
	Password  string    `gorm:"size:100;not null;"         json:"password"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"  json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"  json:"updated_at"`
}

func EncPassword(password string) (encPassword string, err error) {
	hashedPassword, err := Hash(password)
	if err != nil {
		return
	}
	encPassword = string(hashedPassword)
	return
}

func RegisterUser(req requests.RegisterUser) (err error) {
	EmailValidation(req.Email)
	if err != nil {
		return
	}
	NicknameValication(req.Nickname)
	if err != nil {
		return
	}

	if err != nil {
		log.Fatal(err)
	}

	encPassword, err := EncPassword(req.Password)
	if err != nil {
		return
	}

	user := &Users{
		Nickname:  req.Nickname,
		Email:     req.Email,
		Password:  encPassword,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = user.Create(gGormDB)
	return
}

func EmailValidation(email string) (err error) {
	var user []Users

	qs := NewUsersQuerySet(gGormDB)
	qs = qs.EmailEq(email)
	err = qs.All(&user)
	if len(user) >= 1 {
		return errors.New("Exist Email")
	}

	if err := checkmail.ValidateFormat(email); err != nil {
		return errors.New("Invalid Email")
	}

	return
}

func NicknameValication(nickname string) (err error) {
	var user []Users

	qs := NewUsersQuerySet(gGormDB)
	qs = qs.NicknameEq(nickname)
	err = qs.All(&user)
	if len(user) >= 1 {
		return
	}

	return
}

func GetMyProfile(userEmail string) (user Users, err error) {
	qs := NewUsersQuerySet(gGormDB)
	qs = qs.EmailEq(userEmail)
	err = qs.One(&user)

	return
}

func ModifyProfile(req requests.RegisterUser, userID string) (err error) {
	_, err = NewUsersQuerySet(gGormDB).
		GetUpdater().
		SetEmail(req.Email).
		SetNickname(req.Nickname).
		SetPassword(req.Password).
		UpdateNum()

	return
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
