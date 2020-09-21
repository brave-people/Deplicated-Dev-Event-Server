package requests

import "time"

type RegisterUser struct {
	Nickname  string    `gorm:"size:255;not null;unique"   json:"nickname"`
	Email     string    `gorm:"size:100;not null;unique"   json:"email"`
	Password  string    `gorm:"size:100;not null;"         json:"password"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"  json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"  json:"updated_at"`
}
