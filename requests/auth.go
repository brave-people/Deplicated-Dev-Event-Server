package requests

type Login struct {
	Email    string `gorm:"size:100;not null;unique"   json:"email"`
	Password string `gorm:"size:100;not null;"         json:"password"`
}
